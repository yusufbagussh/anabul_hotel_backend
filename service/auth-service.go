package service

import (
	"fmt"
	//"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
	"github.com/thanhpk/randstr"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"github.com/yusufbagussh/pet_hotel_backend/repository"
	"github.com/yusufbagussh/pet_hotel_backend/utils"
	"golang.org/x/crypto/bcrypt"
	"log"
	//"net/http"
	"strings"
	"time"
)

// AuthService is a contract about something that this service can do
type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(user dto.RegisterDTO) (entity.User, error)
	CreateVerificationUser(user dto.RegisterDTO) (entity.User, error, utils.EmailData)
	IsDuplicateEmail(email string) bool
	VerificationEmail(user dto.ForgotPasswordDTO) (entity.User, utils.EmailData)
	ResetPass(newPass dto.ResetPasswordInput) (entity.User, interface{})
	ActivationEmail(code string) (entity.User, interface{})
	//FindByEmail(email string) (entity.User, error)
}

type authService struct {
	userRepository repository.UserRepository
	redisService   RedisService
}

// NewAuthService creates a new instance of AuthService
func NewAuthService(userRep repository.UserRepository, redisServ RedisService) AuthService {
	return &authService{
		userRepository: userRep,
		redisService:   redisServ,
	}
}

// VerifyCredential is to checking data user in database
func (s *authService) VerifyCredential(email string, password string) interface{} {
	res := s.userRepository.VerifyCredential(email)
	if v, ok := res.(entity.User); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if (v.Email == email && comparedPassword) && v.Verified == true {
			return res
		}
		return false
	}
	return false
}

func (s *authService) ResetPass(newPass dto.ResetPasswordInput) (entity.User, interface{}) {
	passwordResetToken := utils.Encode(newPass.Token)

	res, err := s.userRepository.FindByEmail(newPass.Email)
	fmt.Println(res)
	if err != nil {
		return res, err
	}

	//check token in redis
	fmt.Println(passwordResetToken)
	checkRedis := s.redisService.CheckValueKey(newPass.Email, passwordResetToken)
	if checkRedis != true {
		return entity.User{}, false
	}

	res.Password = newPass.Password
	res.PasswordResetToken = ""

	newDataPass, errChange := s.userRepository.ChangePass(res)
	if errChange == nil {
		s.redisService.ClearToken(newPass.Email)
	}
	return newDataPass, errChange
}

func (s *authService) ActivationEmail(verificationCode string) (entity.User, interface{}) {

	//res, err := s.userRepository.FindByEmail(newPass.Email)
	//fmt.Println(res)
	//if err != nil {
	//	return res, err
	//}
	fmt.Println(verificationCode)
	//check token in redis
	//verificationCode = utils.Encode(verificationCode)
	checkRedis := s.redisService.ValidateToken(verificationCode)
	fmt.Println(checkRedis)
	if checkRedis != true {
		return entity.User{}, false
	}

	value := s.redisService.GetValueByKey(verificationCode)

	if value != "" {
		fmt.Println(value)
	} else {
		fmt.Println("Value tidak ada")
	}

	var UserID string
	if strings.Contains(value, "_") {
		UserID = strings.Split(value, "_")[1]
	}

	if UserID != "" {
		fmt.Println(UserID)
	} else {
		fmt.Println("Id tidak ada")
	}

	res, err := s.userRepository.FindUserByID(UserID)
	fmt.Println(res.ID)
	if err != nil {
		return entity.User{}, err
	}

	res.Verified = true
	user, errUpdate := s.userRepository.UpdateVerified(res)
	if errUpdate == nil {
		s.redisService.ClearToken(verificationCode)
	}
	return user, errUpdate
}

func (s *authService) VerificationEmail(email dto.ForgotPasswordDTO) (entity.User, utils.EmailData) {
	res, err := s.userRepository.FindByEmail(email.Email)

	//seharusnya tanggi errornya terlebih dahulu
	//if err != nil{}
	if err == nil {
		// Generate Verification Code
		resetToken := randstr.String(20)
		fmt.Println(resetToken)

		passwordResetToken := utils.Encode(resetToken)
		fmt.Println(passwordResetToken)
		//res.PasswordResetToken = passwordResetToken
		//res.PasswordResetAt = time.Now().Add(time.Minute * 15)

		s.redisService.SaveToken(email.Email, passwordResetToken, 15*time.Minute)

		//s.userRepository.UpdateUser(res)

		var firstName = res.Name

		if strings.Contains(firstName, " ") {
			firstName = strings.Split(firstName, " ")[1]
		}

		// ? Send Email
		emailData := utils.EmailData{
			URL:       "http://localhost:8080/api/user/resetpassword/" + resetToken,
			FirstName: firstName,
			Subject:   "Your password reset token (valid for 10min)",
		}
		return res, emailData
	}
	return res, utils.EmailData{}
}

// CreateUser is func to add user in database
func (s *authService) CreateUser(user dto.RegisterDTO) (entity.User, error) {
	userToCreate := entity.User{}
	//dengan memasukkan struct mana yang akan diisi, data dari mana yang akan diisikan
	errMap := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if errMap != nil {
		return userToCreate, errMap
	}
	res, err := s.userRepository.InsertUser(userToCreate)
	return res, err
}

func (s *authService) CreateVerificationUser(user dto.RegisterDTO) (entity.User, error, utils.EmailData) {
	userToCreate := entity.User{}
	//dengan memasukkan struct mana yang akan diisi, data dari mana yang akan diisikan
	errMap := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if errMap != nil {
		return userToCreate, errMap, utils.EmailData{}
	}

	code := randstr.String(20)

	verificationCode := utils.Encode(code)

	res, err := s.userRepository.InsertUser(userToCreate)

	codeID := verificationCode + "_" + res.ID
	s.redisService.SaveToken(verificationCode, codeID, 15*time.Minute)

	var firstName = res.Name

	if strings.Contains(firstName, " ") {
		firstName = strings.Split(firstName, " ")[1]
	}
	// ? Send Email
	emailData := utils.EmailData{
		URL:       "http://localhost:8080/api/auth/verify/" + verificationCode,
		FirstName: firstName,
		Subject:   "Your account verification code"}
	return res, err, emailData
}

// FindByEmail is func to show user by email
//func (s *authService) FindByEmail(email string) (entity.User, error) {
//	return s.userRepository.FindByEmail(email)
//}

// IsDuplicateEmail is func to check duplicate in database
func (s *authService) IsDuplicateEmail(email string) bool {
	_, err := s.userRepository.FindByEmail(email)
	return !(err == nil)
}

// comparePassword is to compare password in database that decoded with password was inputted by user
func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
