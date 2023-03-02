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
	"github.com/yusufbagussh/pet_hotel_backend/helper"
	"github.com/yusufbagussh/pet_hotel_backend/service"
	"log"
	"net/http"
)

// UserController is a contract of what userController can do
type UserController interface {
	UpdateProfile(ctx *gin.Context)
	GetProfile(ctx *gin.Context)
	ChangePassword(ctx *gin.Context)
	GetAdmin(ctx *gin.Context)
	GetStaff(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	DeleteAdmin(ctx *gin.Context)
	ShowAdmin(ctx *gin.Context)
	UpdateAdmin(ctx *gin.Context)
	DeleteStaff(ctx *gin.Context)
	ShowStaff(ctx *gin.Context)
	UpdateStaff(ctx *gin.Context)
	SendNotif(ctx *gin.Context)
	SaveDeviceToken(ctx *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

// NewUserController is creating anew instance of UserControlller
func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (u *userController) SendNotif(ctx *gin.Context) {
	app, _, _ := config.SetupFirebase()
	sendToToken(app)
}

func sendToToken(app *firebase.App) {
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

func (u *userController) CreateUser(ctx *gin.Context) {
	var CreateUser dto.UserHotelCreateDTO
	errDTO := ctx.ShouldBind(&CreateUser)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, errCreate := u.userService.CreateUser(CreateUser, ctx)
	if errCreate != nil {
		res := helper.BuildErrorResponse("Failed to create user", errCreate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Create success", result)
	ctx.JSON(http.StatusOK, res)
}
func (u *userController) UpdateAdmin(ctx *gin.Context) {
	var UpdateUser dto.UserHotelUpdateDTO
	errDTO := ctx.ShouldBind(&UpdateUser)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, errCreate := u.userService.UpdateUser(UpdateUser, ctx)
	if errCreate != nil {
		res := helper.BuildErrorResponse("Failed to update user", errCreate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Update success", result)
	ctx.JSON(http.StatusOK, res)
}
func (u *userController) DeleteAdmin(ctx *gin.Context) {
	userID := ctx.Param("id")
	errDel := u.userService.DeleteUser(userID)
	if errDel != nil {
		res := helper.BuildErrorResponse("Failed to delete", errDel.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Delete success", nil)
	ctx.JSON(http.StatusOK, res)
}
func (u *userController) ShowAdmin(ctx *gin.Context) {
	userID := ctx.Param("id")
	result, errShow := u.userService.ShowUser(userID)
	if errShow != nil {
		res := helper.BuildErrorResponse("Failed to show data", errShow.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Show data success", result)
	ctx.JSON(http.StatusOK, res)
}

func (u *userController) GetAdmin(ctx *gin.Context) {
	var filterPagination dto.FilterPagination
	errDTO := ctx.ShouldBind(&filterPagination)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	var admins, pagination, err = u.userService.AllAdmin(filterPagination)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get all data", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponsePage(true, "Success Get All Admin", admins, pagination)
	ctx.JSON(http.StatusOK, res)
}

func (u *userController) GetStaff(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, _ := u.jwtService.ValidateToken(authHeader)
	claims := token.Claims.(jwt.MapClaims)
	hotelID := fmt.Sprintf("%v", claims["hotel_id"])

	var filterPagination dto.FilterPagination
	errDTO := ctx.ShouldBind(&filterPagination)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var staffs, pagination, errGet = u.userService.AllStaff(hotelID, filterPagination)
	if errGet != nil {
		res := helper.BuildErrorResponse("Failed to get all data", errGet.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponsePage(true, "Success Get All Staff", staffs, pagination)
	ctx.JSON(http.StatusOK, res)
}

func (u *userController) UpdateStaff(ctx *gin.Context) {
	var UpdateUser dto.UserHotelUpdateDTO
	errDTO := ctx.ShouldBind(&UpdateUser)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	authHeader := ctx.GetHeader("Authorization")
	token, _ := u.jwtService.ValidateToken(authHeader)
	claims := token.Claims.(jwt.MapClaims)
	hotelID := fmt.Sprintf("%v", claims["hotel_id"])
	if UpdateUser.HotelID != hotelID {
		res := helper.BuildErrorResponse(
			"You don't have permission",
			"You are not the owner",
			helper.EmptyObj{},
		)
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}
	result, errCreate := u.userService.UpdateUser(UpdateUser, ctx)
	if errCreate != nil {
		res := helper.BuildErrorResponse("Failed to update user", errCreate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Update success", result)
	ctx.JSON(http.StatusOK, res)
}
func (u *userController) DeleteStaff(ctx *gin.Context) {
	userID := ctx.Param("id")
	authHeader := ctx.GetHeader("Authorization")
	token, _ := u.jwtService.ValidateToken(authHeader)
	claims := token.Claims.(jwt.MapClaims)
	hotelID := fmt.Sprintf("%v", claims["hotel_id"])
	errDel, ok := u.userService.DeleteStaff(userID, hotelID)
	if ok == false {
		res := helper.BuildErrorResponse(
			"You don't have permission",
			"You are not the owner",
			helper.EmptyObj{},
		)
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}
	if errDel != nil {
		res := helper.BuildErrorResponse("Failed to delete", errDel.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Delete success", nil)
	ctx.JSON(http.StatusOK, res)
}
func (u *userController) ShowStaff(ctx *gin.Context) {
	userID := ctx.Param("id")
	authHeader := ctx.GetHeader("Authorization")
	token, _ := u.jwtService.ValidateToken(authHeader)
	claims := token.Claims.(jwt.MapClaims)
	hotelID := fmt.Sprintf("%v", claims["hotel_id"])
	result, errShow, ok := u.userService.ShowStaff(userID, hotelID)
	if ok == false {
		res := helper.BuildErrorResponse(
			"You don't have permission",
			"You are not the owner",
			helper.EmptyObj{},
		)
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}
	if errShow != nil {
		res := helper.BuildErrorResponse("Failed to show data", errShow.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Show data success", result)
	ctx.JSON(http.StatusOK, res)
}

func (u *userController) SaveDeviceToken(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, _ := u.jwtService.ValidateToken(authHeader)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	var device dto.Device
	errDTO := ctx.ShouldBind(&device)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	newUserToken, err := u.userService.SaveDeviceToken(device, userID)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	res := helper.BuildResponse(true, "Success Update Token", newUserToken)
	ctx.JSON(http.StatusOK, res)

}

func (u *userController) ChangePassword(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, _ := u.jwtService.ValidateToken(authHeader)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	var changePassDTO dto.ChangePasswordDTO
	errDTO := ctx.ShouldBind(&changePassDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if changePassDTO.NewPassword != changePassDTO.ConfPassword {
		response := helper.BuildErrorResponse("Failed to process request", "Password not match", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	newUserPass, err := u.userService.ChangePassword(changePassDTO, userID)
	if err != nil {
		response := helper.BuildErrorResponse("Please check your credential", "Invalid credential", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	} else {
		response := helper.BuildResponse(true, "Success change password!", newUserPass)
		ctx.JSON(http.StatusOK, response)
		return
	}
}

func (u *userController) UpdateProfile(ctx *gin.Context) {
	var userUpdateDTO dto.UserHotelUpdateDTO
	errDTO := ctx.ShouldBind(&userUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token, _ := u.jwtService.ValidateToken(authHeader)
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	role := fmt.Sprintf("%v", claims["role"])
	userUpdateDTO.ID = id
	userUpdateDTO.Role = role
	user, errUpdate := u.userService.Update(userUpdateDTO, ctx)
	if errUpdate != nil {
		res := helper.BuildErrorResponse("Failed to update profile", errUpdate.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Success Update Data Profile", user)
	ctx.JSON(http.StatusOK, res)
}

func (u *userController) GetProfile(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, _ := u.jwtService.ValidateToken(authHeader)
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	user, errProfile := u.userService.GetProfile(id)
	if errProfile != nil {
		res := helper.BuildErrorResponse("No data found", errProfile.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Success Get Data Profile", user)
	ctx.JSON(http.StatusOK, res)
}
