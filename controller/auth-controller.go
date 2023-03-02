package controller

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/yusufbagussh/pet_hotel_backend/config"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"github.com/yusufbagussh/pet_hotel_backend/helper"
	"github.com/yusufbagussh/pet_hotel_backend/service"
	"github.com/yusufbagussh/pet_hotel_backend/utils"
	"log"
	"net/http"
	"time"
)

// AuthController interface is a contract what this controller can do
type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	Logout(ctx *gin.Context)
	Refresh(ctx *gin.Context)
	ForgotPass(ctx *gin.Context)
	ResetPass(ctx *gin.Context)
	TestFCM(ctx *gin.Context)
	RegisterUser(ctx *gin.Context)
	ActivationEmail(ctx *gin.Context)
}

type authController struct {
	authService  service.AuthService
	jwtService   service.JWTService
	redisService service.RedisService
}

// NewAuthController creates a new instance of AuthController
func NewAuthController(authService service.AuthService, jwtService service.JWTService, redisService service.RedisService) AuthController {
	return &authController{
		authService:  authService,
		jwtService:   jwtService,
		redisService: redisService,
	}
}

func (a *authController) RegisterUser(ctx *gin.Context) {
	var registerDTO dto.RegisterDTO
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if !a.authService.IsDuplicateEmail(registerDTO.Email) {
		response := helper.BuildErrorResponse("Failed to process request", "Duplicate email", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
		return
	}
	createdUser, createError, emailData := a.authService.CreateVerificationUser(registerDTO)
	if createError != nil {
		response := helper.BuildErrorResponse("Failed to insert data", createError.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
		return
	}
	errSend := utils.SendEmail(&createdUser, &emailData, "emailReset.html")
	if errSend != nil {
		response := helper.BuildErrorResponse("Failed to send verification email", errSend.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	message := "We sent an email with a verification code to " + createdUser.Email
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "message": message})
}

// [...] Verify Email
func (a *authController) ActivationEmail(ctx *gin.Context) {
	verificationCode := ctx.Params.ByName("verificationCode")

	userActive, err := a.authService.ActivationEmail(verificationCode)
	if err == false {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "No token found"})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "The reset token is invalid or has expired"})
		return
	}

	response := helper.BuildResponse(true, "Success reset password!", userActive)
	ctx.JSON(http.StatusOK, response)
}

func (a *authController) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterDTO
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if !a.authService.IsDuplicateEmail(registerDTO.Email) {
		response := helper.BuildErrorResponse("Failed to process request", "Duplicate email", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
		return
	}
	createdUser, createError := a.authService.CreateUser(registerDTO)
	if createError != nil {
		response := helper.BuildErrorResponse("Failed to insert data", createError.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
		return
	}
	token := a.jwtService.GenerateToken(createdUser.ID, createdUser.Email, createdUser.Role, createdUser.HotelID)
	a.redisService.SaveToken(createdUser.ID, token, 24*time.Hour)
	createdUser.Token = token
	response := helper.BuildResponse(true, "Register success!", createdUser)
	ctx.JSON(http.StatusCreated, response)
}

func (a *authController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := a.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if user, ok := authResult.(entity.User); ok {
		generatedToken := a.jwtService.GenerateToken(user.ID, user.Email, user.Role, user.HotelID)
		a.redisService.SaveToken(user.ID, generatedToken, 24*time.Hour)
		user.Token = generatedToken
		response := helper.BuildResponse(true, "Login success!", user)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helper.BuildErrorResponse("Please check again your credential", "Invalid Credential", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (a authController) Logout(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")

	token, _ := a.jwtService.ValidateToken(authHeader)

	claims := token.Claims.(jwt.MapClaims)
	a.redisService.ClearToken(claims["user_id"].(string))
	response := helper.BuildResponse(true, "Logout Berhasil!", nil)
	ctx.JSON(http.StatusOK, response)
}

func (a *authController) Refresh(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, _ := a.jwtService.ValidateToken(authHeader)

	if !token.Valid {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	oldClaims := token.Claims.(jwt.MapClaims)
	tokenString := a.jwtService.GenerateToken(oldClaims["user_id"].(string), oldClaims["email"].(string), oldClaims["role"].(string), oldClaims["hotel_id"].(string))

	//Delete old token by user_id
	a.redisService.ClearToken(oldClaims["user_id"].(string))
	//Save new token with key user_id
	a.redisService.SaveToken(oldClaims["user_id"].(string), tokenString, 24*time.Hour)

	response := helper.BuildResponse(true, "Refresh Token Berhasil!", tokenString)
	ctx.JSON(http.StatusOK, response)
}

func (a *authController) ForgotPass(ctx *gin.Context) {
	var forgotPassDTO dto.ForgotPasswordDTO
	errDTO := ctx.ShouldBind(&forgotPassDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	res, emailData := a.authService.VerificationEmail(forgotPassDTO)
	errSend := utils.SendEmail(&res, &emailData, "emailReset.html")
	if errSend != nil {
		response := helper.BuildErrorResponse("Failed to send verification email", errSend.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	message := "You will receive a reset email if user with that email exist"
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
}

func (a *authController) ResetPass(ctx *gin.Context) {
	var resetPass dto.ResetPasswordInput

	if err := ctx.ShouldBindJSON(&resetPass); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	if resetPass.Password != resetPass.PasswordConfirm {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Passwords do not match"})
		return
	}

	newResetPass, err := a.authService.ResetPass(resetPass)
	if err == false {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "No token found"})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "The reset token is invalid or has expired"})
		return
	}

	response := helper.BuildResponse(true, "Success reset password!", newResetPass)
	ctx.JSON(http.StatusOK, response)
}

func (c *authController) TestFCM(ctx *gin.Context) {
	app, _, _ := config.SetupFirebase()
	sendNotification(app)
}
func sendNotification(app *firebase.App) {
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	fmt.Println("Tes 1")

	registrationToken := "dVib4muX6vS4vk3SUDKpRG:APA91bGfNG2lQSubB6o5v9D0Aj5ClnhW3DEgPbZ0fifj-Yh0hJ_U40OsUGtp6Zd6jM_YDYwOu1Zx7ZTGKAKIvA9aoZU_YzZMq1gw_EoI-UyiPitpHHC4Z6Fl8hSFxNXAzqSm9QG1KQHz"

	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: "Notification Test",
			Body:  "Hello React!!",
		},
		Token: registrationToken,
	}

	fmt.Println("Tes 2")

	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Tes 3")

	fmt.Println("Successfully sent message:", response)
}
