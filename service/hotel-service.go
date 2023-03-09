package service

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"github.com/yusufbagussh/pet_hotel_backend/helper"
	"github.com/yusufbagussh/pet_hotel_backend/repository"
	"os"
	"path/filepath"
)

// HotelService is a contract of what hotelService can do
type HotelService interface {
	GetAllHotel(filterPagination dto.FilterPagination) ([]entity.Hotel, dto.Pagination, error)
	CreateHotel(hotel dto.CreateHotel, ctx *gin.Context) (interface{}, error)
	UpdateHotel(hotel dto.UpdateHotel, ctx *gin.Context) (entity.Hotel, error)
	DeleteHotel(hotelID string) error
	ShowHotel(hotelID string) (entity.Hotel, error)
	CreateHotelAdmin(hotel dto.CreateHotel, ctx *gin.Context) (interface{}, error)
	//IsAllowedHotel(hotelID string, userID string) bool
}

type hotelService struct {
	hotelRepository repository.HotelRepository
	checkHelper     helper.CheckHelper
}

// NewHotelService creates a new instance of HotelService
func NewHotelService(hotelRepo repository.HotelRepository, checkHelp helper.CheckHelper) HotelService {
	return &hotelService{
		hotelRepository: hotelRepo,
		checkHelper:     checkHelp,
	}
}

func (u *hotelService) CreateHotelAdmin(hotel dto.CreateHotel, ctx *gin.Context) (interface{}, error) {
	hotelEntity := entity.Hotel{}
	userEntity := entity.User{}
	if hotel.Image != nil {
		fileName := hotel.Image.Filename
		extension := filepath.Ext(hotel.Image.Filename) // Generate random file name for the new uploaded file, so it doesn't override the old file with same name
		newImage := fileName + extension
		hotelEntity.Image = newImage

		err := ctx.SaveUploadedFile(hotel.Image, "uploads/user-profiles/"+newImage)
		if err != nil {
			return hotelEntity, err
		}
	}
	if hotel.Document != nil {
		extension := filepath.Ext(hotel.Document.Filename) // Generate random file name for the new uploaded file, so it doesn't override the old file with same name
		newDocument := "Document" + "_" + hotel.Name + extension
		hotelEntity.Document = newDocument

		err := ctx.SaveUploadedFile(hotel.Document, "uploads/documents/"+newDocument)
		if err != nil {
			return hotelEntity, err
		}
	}
	if hotel.NPWP != nil {
		extension := filepath.Ext(hotel.NPWP.Filename) // Generate random file name for the new uploaded file, so it doesn't override the old file with same name
		newNPWP := "NPWP" + "_" + hotel.Name + extension
		hotelEntity.NPWP = newNPWP

		err := ctx.SaveUploadedFile(hotel.NPWP, "uploads/documents/"+newNPWP)
		if err != nil {
			return hotelEntity, err
		}
	}
	hotelEntity.Name = hotel.Name
	hotelEntity.Email = hotel.Email
	hotelEntity.Address = hotel.Address
	hotelEntity.Phone = hotel.Phone
	hotelEntity.ProvinceID = hotel.ProvinceID
	hotelEntity.CityID = hotel.CityID
	hotelEntity.CityID = hotel.CityID
	hotelEntity.DistrictID = hotel.DistrictID
	hotelEntity.Latitude = hotel.Latitude
	hotelEntity.Longitude = hotel.Longitude
	hotelEntity.Requirement = hotel.Requirement
	hotelEntity.Regulation = hotel.Regulation

	if hotel.UserHotelCreateDTO.Image != nil {
		extension := filepath.Ext(hotel.UserHotelCreateDTO.Image.Filename) // Generate random file name for the new uploaded file, so it doesn't override the old file with same name
		newImage := uuid.New().String() + extension
		userEntity.Image = newImage

		err := ctx.SaveUploadedFile(hotel.UserHotelCreateDTO.Image, "uploads/user-profiles/"+newImage)
		if err != nil {
			return userEntity, err
		}
	}

	if hotel.UserHotelCreateDTO.KTP != nil {
		extension := filepath.Ext(hotel.UserHotelCreateDTO.KTP.Filename) // Generate random file name for the new uploaded file, so it doesn't override the old file with same name
		newNIKFile := "NIK" + "_" + hotel.UserHotelCreateDTO.Name + extension
		userEntity.KTP = newNIKFile

		err := ctx.SaveUploadedFile(hotel.UserHotelCreateDTO.KTP, "uploads/user-profiles/"+newNIKFile)
		if err != nil {
			return userEntity, err
		}
	}

	userEntity.Name = hotel.UserHotelCreateDTO.Name
	userEntity.Email = hotel.UserHotelCreateDTO.Email
	userEntity.Password = hotel.UserHotelCreateDTO.Password
	userEntity.Address = hotel.UserHotelCreateDTO.Address
	userEntity.Phone = hotel.UserHotelCreateDTO.Phone
	userEntity.Role = hotel.UserHotelCreateDTO.Role
	userEntity.HotelID = hotel.UserHotelCreateDTO.HotelID
	userEntity.NIK = hotel.UserHotelCreateDTO.NIK

	updatedHotel, err := u.hotelRepository.InsertHotelAdmin(hotelEntity, userEntity)
	return updatedHotel, err
}

func (u *hotelService) CreateHotel(hotel dto.CreateHotel, ctx *gin.Context) (interface{}, error) {
	hotelEntity := entity.Hotel{}
	if hotel.Image != nil {
		extension := filepath.Ext(hotel.Image.Filename) // Generate random file name for the new uploaded file, so it doesn't override the old file with same name
		newImage := uuid.NewString() + extension
		hotelEntity.Image = newImage

		err := ctx.SaveUploadedFile(hotel.Image, "uploads/user-profiles/"+newImage)
		if err != nil {
			return hotelEntity, err
		}
	}
	if hotel.Document != nil {
		extension := filepath.Ext(hotel.Document.Filename) // Generate random file name for the new uploaded file, so it doesn't override the old file with same name
		newDocument := "Document" + hotel.Name + extension
		hotelEntity.Document = newDocument

		err := ctx.SaveUploadedFile(hotel.Document, "uploads/documents/"+newDocument)
		if err != nil {
			return hotelEntity, err
		}
	}
	if hotel.NPWP != nil {
		extension := filepath.Ext(hotel.NPWP.Filename) // Generate random file name for the new uploaded file, so it doesn't override the old file with same name
		newNPWP := "NPWP" + hotel.Name + extension
		hotelEntity.NPWP = newNPWP

		err := ctx.SaveUploadedFile(hotel.NPWP, "uploads/documents/"+newNPWP)
		if err != nil {
			return hotelEntity, err
		}
	}
	hotelEntity.Name = hotel.Name
	hotelEntity.Email = hotel.Email
	hotelEntity.Address = hotel.Address
	hotelEntity.Phone = hotel.Phone
	hotelEntity.ProvinceID = hotel.ProvinceID
	hotelEntity.CityID = hotel.CityID
	hotelEntity.CityID = hotel.CityID
	hotelEntity.DistrictID = hotel.DistrictID
	hotelEntity.Latitude = hotel.Latitude
	hotelEntity.Longitude = hotel.Longitude
	hotelEntity.Requirement = hotel.Requirement
	hotelEntity.Regulation = hotel.Regulation

	updatedHotel, err := u.hotelRepository.InsertHotel(hotelEntity)
	return updatedHotel, err
}
func (u *hotelService) UpdateHotel(hotel dto.UpdateHotel, ctx *gin.Context) (entity.Hotel, error) {
	hotelToUpdate := entity.Hotel{}
	if hotel.Image != nil {
		extension := filepath.Ext(hotel.Image.Filename) // Generate random file name for the new uploaded file, so it doesn't override the old file with same name
		newImage := uuid.New().String() + extension
		hotelToUpdate.Image = newImage

		err := ctx.SaveUploadedFile(hotel.Image, "uploads/user-profiles/"+newImage)
		if err != nil {
			return hotelToUpdate, err
		} else {
			hotelByID, _ := u.hotelRepository.FindHotelByID(hotel.IDHotel)
			if hotelByID.Image != "" {
				errRemove := os.Remove("./uploads/user-profiles/" + hotelByID.Image)
				if errRemove != nil {
					return hotelToUpdate, errRemove
				}
			}
		}
	}
	if hotel.Document != nil {
		extension := filepath.Ext(hotel.Document.Filename) // Generate random file name for the new uploaded file, so it doesn't override the old file with same name
		newImage := uuid.New().String() + extension
		hotelToUpdate.Document = newImage

		err := ctx.SaveUploadedFile(hotel.Document, "uploads/user-profiles/"+newImage)
		if err != nil {
			return hotelToUpdate, err
		} else {
			hotelByID, _ := u.hotelRepository.FindHotelByID(hotel.IDHotel)
			if hotelByID.Document != "" {
				errRemove := os.Remove("./uploads/user-profiles/" + hotelByID.Document)
				if errRemove != nil {
					return hotelToUpdate, errRemove
				}
			}
		}
	}
	if hotel.NPWP != nil {
		extension := filepath.Ext(hotel.NPWP.Filename) // Generate random file name for the new uploaded file, so it doesn't override the old file with same name
		newImage := uuid.New().String() + extension
		hotelToUpdate.NPWP = newImage

		err := ctx.SaveUploadedFile(hotel.NPWP, "uploads/user-profiles/"+newImage)
		if err != nil {
			return hotelToUpdate, err
		} else {
			hotelByID, _ := u.hotelRepository.FindHotelByID(hotel.IDHotel)
			if hotelByID.NPWP != "" {
				errRemove := os.Remove("./uploads/user-profiles/" + hotelByID.NPWP)
				if errRemove != nil {
					return hotelToUpdate, errRemove
				}
			}
		}
	}
	hotelToUpdate.IDHotel = hotel.IDHotel
	hotelToUpdate.Name = hotel.Name
	hotelToUpdate.Email = hotel.Email
	hotelToUpdate.Address = hotel.Address
	hotelToUpdate.Phone = hotel.Phone
	hotelToUpdate.ProvinceID = hotel.ProvinceID
	hotelToUpdate.CityID = hotel.CityID
	hotelToUpdate.DistrictID = hotel.DistrictID
	hotelToUpdate.Latitude = hotel.Latitude
	hotelToUpdate.Longitude = hotel.Longitude
	hotelToUpdate.MapLink = hotel.MapLink
	hotelToUpdate.Requirement = hotel.Requirement
	hotelToUpdate.Description = hotel.Description
	hotelToUpdate.Regulation = hotel.Regulation
	hotelToUpdate.OpenTime = hotel.OpenTime
	hotelToUpdate.CloseTime = hotel.CloseTime

	updatedHotel, err := u.hotelRepository.UpdateHotel(hotelToUpdate)
	return updatedHotel, err
}
func (u *hotelService) DeleteHotel(hotelID string) error {
	hotel, err := u.hotelRepository.FindHotelByID(hotelID)
	if err != nil {
		return err
	}
	if hotel.Image != "" {
		errRemove := os.Remove("./uploads/user-profiles/" + hotel.Image)
		if errRemove != nil {
			return errRemove
		}
	}
	if hotel.Document != "" {
		errRemove := os.Remove("./uploads/user-profiles/" + hotel.Document)
		if errRemove != nil {
			return errRemove
		}
	}
	if hotel.NPWP != "" {
		errRemove := os.Remove("./uploads/user-profiles/" + hotel.NPWP)
		if errRemove != nil {
			return errRemove
		}
	}
	errDel := u.hotelRepository.DeleteHotel(hotel)
	return errDel
}
func (u *hotelService) ShowHotel(hotelID string) (entity.Hotel, error) {
	result, err := u.hotelRepository.FindHotelByID(hotelID)
	return result, err
}
func (u *hotelService) GetAllHotel(filterPagination dto.FilterPagination) ([]entity.Hotel, dto.Pagination, error) {
	admins, pagination, err := u.hotelRepository.GetAllHotel(filterPagination)
	return admins, pagination, err
}
