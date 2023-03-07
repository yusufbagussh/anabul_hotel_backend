package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"github.com/yusufbagussh/pet_hotel_backend/helper"
	"github.com/yusufbagussh/pet_hotel_backend/repository"
	"github.com/yusufbagussh/pet_hotel_backend/utils"
	"log"
	"os"
	"path/filepath"
)

// RequestService is a contract of what requestService can do
type RequestService interface {
	GetAllRequest(filterPage dto.FilterPagination) ([]entity.Request, dto.Pagination, error)
	CreateRequest(request dto.CreateRequest, ctx *gin.Context) (entity.Request, error)
	UpdateRequest(request dto.UpdateRequest, ctx *gin.Context) (entity.Request, error, error)
	DeleteRequest(requestID string) error
	ShowRequest(requestID string) (entity.Request, error)
	UpdateRequestStatus(requestStatus dto.UpdateRequestStatus) (entity.Request, error, error)
}

type requestService struct {
	requestRepository repository.RequestRepository
	userRepository    repository.UserRepository
	hotelRepository   repository.HotelRepository
	checkHelper       helper.CheckHelper
}

// NewRequestService creates a new instance of RequestService
func NewRequestService(requestRepo repository.RequestRepository, userRepository repository.UserRepository, hotelRepository repository.HotelRepository, checkHelper helper.CheckHelper) RequestService {
	return &requestService{
		requestRepository: requestRepo,
		userRepository:    userRepository,
		hotelRepository:   hotelRepository,
		checkHelper:       checkHelper,
	}
}

func (u *requestService) UpdateRequestStatus(requestStatus dto.UpdateRequestStatus) (entity.Request, error, error) {
	var hotelEntity entity.Hotel
	var userEntity entity.User
	updatedRequest, errUpdate := u.requestRepository.UpdateRequestStatus(requestStatus)
	if errUpdate != nil {
		//return entity.Request{}, errUpdate, utils.EmailData{}
		return entity.Request{}, errUpdate, nil
	}
	fmt.Println(updatedRequest)
	if updatedRequest.Status == "Terima" {
		hotelEntity.Name = updatedRequest.HotelName
		hotelEntity.Phone = updatedRequest.HotelPhone
		hotelEntity.Email = updatedRequest.HotelEmail
		hotelEntity.NPWP = updatedRequest.NPWP
		hotelEntity.Document = updatedRequest.Document

		hotel, errHotel := u.hotelRepository.InsertHotel(hotelEntity)
		if errHotel != nil {
			//return updatedRequest, errHotel, utils.EmailData{}
			return updatedRequest, errHotel, nil
		}

		password := utils.RandomPassword(8, 2, 2, 2)

		userEntity.HotelID = hotel.IDHotel
		userEntity.Name = updatedRequest.AdminName
		userEntity.Password = password
		userEntity.Email = updatedRequest.HotelEmail
		userEntity.NIK = updatedRequest.NIK
		userEntity.KTP = updatedRequest.KTP
		userEntity.Selfie = updatedRequest.Selfie
		userEntity.Role = "Admin"
		userEntity.Phone = updatedRequest.AdminPhone
		userEntity.Verified = true

		_, errUser := u.userRepository.InsertUser(userEntity)
		if errUser != nil {
			//return updatedRequest, errUser, utils.EmailData{}
			return updatedRequest, errUser, nil
		}

		//Send Email v1
		data := utils.Account{
			Email:    userEntity.Email,
			Password: userEntity.Password,
		}
		errSend := utils.SendEmailAccept(userEntity.Email, data)
		if errSend != nil {
			return entity.Request{}, nil, errSend
		}

	} else if updatedRequest.Status == "Tolak" {
		errSend := utils.SendEmailReject(updatedRequest.HotelEmail)
		if errSend != nil {
			return entity.Request{}, nil, errSend
		}
	}

	return updatedRequest, errUpdate, nil
}

func (u *requestService) CreateRequest(request dto.CreateRequest, ctx *gin.Context) (entity.Request, error) {
	requestEntity := entity.Request{}
	fmt.Println(request.HotelName)
	errDir := os.Mkdir("uploads/documents/"+request.HotelName, 0750)
	if errDir != nil && !os.IsExist(errDir) {
		log.Fatal(errDir)
	}
	//errDir = os.WriteFile("uploads/documents/"+request.HotelName+"/testfile.txt", []byte("Hello, Gophers!"), 0660)
	//if errDir != nil {
	//	log.Fatal(errDir)
	//}

	var err error

	if request.Document != nil {
		extension := filepath.Ext(request.Document.Filename) // Generate random file name for the new uploaded file, so it doesn't override the old file with same name
		newDocument := "Document" + "_" + request.HotelName + extension
		requestEntity.Document = newDocument

		err = ctx.SaveUploadedFile(request.Document, "uploads/documents/"+request.HotelName+"/"+newDocument)
		if err != nil {
			return requestEntity, err
		}
	}
	if request.NPWP != nil {
		extension := filepath.Ext(request.NPWP.Filename) // Generate random file name for the new uploaded file, so it doesn't override the old file with same name
		newNPWP := "NPWP" + "_" + request.HotelName + extension
		requestEntity.NPWP = newNPWP

		err = ctx.SaveUploadedFile(request.NPWP, "uploads/documents/"+request.HotelName+"/"+newNPWP)
		if err != nil {
			return requestEntity, err
		}
	}
	if request.KTP != nil {
		extension := filepath.Ext(request.KTP.Filename) // Generate random file name for the new uploaded file, so it doesn't override the old file with same name
		newKTP := "KTP" + "_" + request.HotelName + extension
		requestEntity.KTP = newKTP

		err = ctx.SaveUploadedFile(request.KTP, "uploads/documents/"+request.HotelName+"/"+newKTP)
		if err != nil {
			return requestEntity, err
		}
	}
	if request.Selfie != nil {
		extension := filepath.Ext(request.Selfie.Filename) // Generate random file name for the new uploaded file, so it doesn't override the old file with same name
		newSelfie := "Selfie" + "_" + request.HotelName + extension
		requestEntity.Selfie = newSelfie

		err = ctx.SaveUploadedFile(request.Selfie, "uploads/documents/"+request.HotelName+"/"+newSelfie)
		if err != nil {
			return requestEntity, err
		}
	}

	requestEntity.HotelName = request.HotelName
	requestEntity.HotelPhone = request.HotelPhone
	requestEntity.HotelEmail = request.HotelEmail
	requestEntity.AdminName = request.AdminName
	requestEntity.AdminPhone = request.AdminPhone
	requestEntity.NIK = request.NIK

	updatedRequest, errCreate := u.requestRepository.InsertRequest(requestEntity)
	return updatedRequest, errCreate
}
func (u *requestService) UpdateRequest(request dto.UpdateRequest, ctx *gin.Context) (entity.Request, error, error) {
	requestEntity := entity.Request{}
	userEntity := entity.User{}
	hotelEntity := entity.Hotel{}

	requestByID, _ := u.requestRepository.FindRequestByID(request.IDRequest)

	var err error
	//if requestByID.HotelName != request.HotelName {
	//	err = os.Rename("uploads/documents/"+request.HotelName, "newTemp")
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}

	if request.Document != nil {
		extension := filepath.Ext(request.Document.Filename) // Generate random file name for the new uploaded file, so it doesn't override the old file with same name
		newDocument := "Document_" + request.HotelName + extension
		requestEntity.Document = newDocument

		err = ctx.SaveUploadedFile(request.Document, "uploads/documents/"+requestByID.HotelName+"/"+newDocument)
		if err != nil {
			//return requestEntity, err, utils.EmailData{}
			return requestEntity, err, nil
		} else {
			//requestByID, _ := u.requestRepository.FindRequestByID(request.IDRequest)
			if requestByID.Document != "" {
				errRemove := os.Remove("./uploads/documents/" + requestByID.HotelName + "/" + requestByID.Document)
				if errRemove != nil {
					return requestEntity, errRemove, nil
				}
			}
		}
	}

	if request.NPWP != nil {
		extension := filepath.Ext(request.NPWP.Filename) // Generate random file name for the new uploaded file, so it doesn't override the old file with same name
		newNPWP := "NPWP_" + request.HotelName + extension
		requestEntity.NPWP = newNPWP

		err = ctx.SaveUploadedFile(request.NPWP, "uploads/documents/"+requestByID.HotelName+"/"+newNPWP)
		if err != nil {
			//return requestEntity, err, utils.EmailData{}
			return requestEntity, err, nil
		} else {
			//requestByID, _ := u.requestRepository.FindRequestByID(request.IDRequest)
			if requestByID.NPWP != "" {
				errRemove := os.Remove("./uploads/documents/" + requestByID.HotelName + "/" + requestByID.NPWP)
				if errRemove != nil {
					//return requestEntity, errRemove, utils.EmailData{}
					return requestEntity, errRemove, nil
				}
			}
		}
	}

	if request.KTP != nil {
		extension := filepath.Ext(request.KTP.Filename) // Generate random file name for the new uploaded file, so it doesn't override the old file with same name
		newKTP := "KTP_" + request.HotelName + extension
		requestEntity.KTP = newKTP

		err = ctx.SaveUploadedFile(request.KTP, "uploads/documents/"+requestByID.HotelName+"/"+newKTP)
		if err != nil {
			//return requestEntity, err, utils.EmailData{}
			return requestEntity, err, nil
		} else {
			//requestByID, _ := u.requestRepository.FindRequestByID(request.IDRequest)
			if requestByID.KTP != "" {
				errRemove := os.Remove("./uploads/documents/" + requestByID.HotelName + "/" + requestByID.KTP)
				if errRemove != nil {
					//return requestEntity, errRemove, utils.EmailData{}
					return requestEntity, errRemove, nil
				}
			}
		}
	}

	if request.Selfie != nil {
		extension := filepath.Ext(request.Selfie.Filename) // Generate random file name for the new uploaded file, so it doesn't override the old file with same name
		newSelfie := "Selfie_" + request.HotelName + extension
		requestEntity.Selfie = newSelfie

		err = ctx.SaveUploadedFile(request.Selfie, "uploads/documents/"+requestByID.HotelName+"/"+newSelfie)
		if err != nil {
			return requestEntity, err, nil
			//return requestEntity, err, utils.EmailData{}
		} else {
			//requestByID, _ := u.requestRepository.FindRequestByID(request.IDRequest)
			if requestByID.Selfie != "" {
				errRemove := os.Remove("./uploads/documents/" + requestByID.HotelName + "/" + requestByID.Selfie)
				if errRemove != nil {
					return requestEntity, errRemove, nil
					//return requestEntity, errRemove, utils.EmailData{}
				}
			}
		}
	}

	requestEntity.IDRequest = request.IDRequest
	requestEntity.HotelName = request.HotelName
	requestEntity.HotelPhone = request.HotelPhone
	requestEntity.HotelEmail = request.HotelEmail
	requestEntity.AdminName = request.AdminName
	requestEntity.AdminPhone = request.AdminPhone
	requestEntity.NIK = request.NIK
	requestEntity.Status = request.Status

	updatedRequest, errUpdate := u.requestRepository.UpdateRequest(requestEntity)
	if errUpdate != nil {
		//return entity.Request{}, errUpdate, utils.EmailData{}
		return entity.Request{}, errUpdate, nil
	}
	fmt.Println(updatedRequest)
	if updatedRequest.Status == "Terima" {
		hotelEntity.Name = updatedRequest.HotelName
		hotelEntity.Phone = updatedRequest.HotelPhone
		hotelEntity.Email = updatedRequest.HotelEmail
		hotelEntity.NPWP = updatedRequest.NPWP
		hotelEntity.Document = updatedRequest.Document

		hotel, errHotel := u.hotelRepository.InsertHotel(hotelEntity)
		if errHotel != nil {
			//return updatedRequest, errHotel, utils.EmailData{}
			return updatedRequest, errHotel, nil
		}

		password := utils.RandomPassword(8, 2, 2, 2)

		userEntity.HotelID = hotel.IDHotel
		userEntity.Name = updatedRequest.AdminName
		userEntity.Password = password
		userEntity.Email = updatedRequest.HotelEmail
		userEntity.NIK = updatedRequest.NIK
		userEntity.KTP = updatedRequest.KTP
		userEntity.Selfie = updatedRequest.Selfie
		userEntity.Role = "Admin"
		userEntity.Phone = updatedRequest.AdminPhone
		userEntity.Verified = true

		_, errUser := u.userRepository.InsertUser(userEntity)
		if errUser != nil {
			//return updatedRequest, errUser, utils.EmailData{}
			return updatedRequest, errUser, nil
		}

		//Send Email v1
		data := utils.Account{
			Email:    userEntity.Email,
			Password: userEntity.Password,
		}
		errSend := utils.SendEmailAccept(userEntity.Email, data)
		if errSend != nil {
			return entity.Request{}, nil, errSend
		}

		//Send Email v2
		//var firstName = updatedRequest.AdminName
		//
		//if strings.Contains(firstName, " ") {
		//	firstName = strings.Split(firstName, " ")[1]
		//}

		//emailData := utils.EmailData{
		//	URL:       "http://localhost:8080/api/user/resetpassword/" + password,
		//	FirstName: firstName,
		//	Subject:   "Selamat anda kami terima",
		//}
		//return updatedRequest, err, emailData

	} else if updatedRequest.Status == "Tolak" {
		//Send Email V1
		errSend := utils.SendEmailReject(updatedRequest.HotelEmail)
		if errSend != nil {
			return entity.Request{}, nil, errSend
		}
		//Send Email V2
		//var firstName = updatedRequest.AdminName
		//
		//if strings.Contains(firstName, " ") {
		//	firstName = strings.Split(firstName, " ")[1]
		//}

		//emailData := utils.EmailData{
		//	URL:       "http://localhost:8080/api/user/resetpassword/",
		//	FirstName: firstName,
		//	Subject:   "Maaf anda kami tolak",
		//}
		//return updatedRequest, err, emailData
	}

	return updatedRequest, err, nil
}
func (u *requestService) DeleteRequest(requestID string) error {
	request, err := u.requestRepository.FindRequestByID(requestID)
	if err != nil {
		return err
	}
	errDel := u.requestRepository.DeleteRequest(request)
	return errDel
}
func (u *requestService) ShowRequest(requestID string) (entity.Request, error) {
	result, err := u.requestRepository.FindRequestByID(requestID)
	return result, err
}
func (u *requestService) GetAllRequest(filterPage dto.FilterPagination) ([]entity.Request, dto.Pagination, error) {
	admins, pagination, err := u.requestRepository.GetAllRequest(filterPage)
	return admins, pagination, err
}
