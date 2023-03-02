package service

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mashingan/smapping"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"github.com/yusufbagussh/pet_hotel_backend/repository"
	"path/filepath"
)

// ReservationConditionService is a contract of what reservationConditionService can do
type ReservationConditionService interface {
	GetAllReservationCondition(hotelID string, filterPagination dto.FilterPagination) ([]entity.ReservationCondition, dto.Pagination, error)
	CreateReservationCondition(reservationCondition dto.CreateReservationCondition, ctx *gin.Context) (entity.ReservationCondition, error)
	UpdateReservationCondition(reservationCondition dto.UpdateReservationCondition) (entity.ReservationCondition, error)
	DeleteReservationCondition(reservationConditionID string) error
	ShowReservationCondition(reservationConditionID string) (entity.ReservationCondition, error)
}

type reservationConditionService struct {
	reservationConditionRepository repository.ReservationConditionRepository
}

// NewReservationConditionService creates a new instance of ReservationConditionService
func NewReservationConditionService(reservationConditionRepo repository.ReservationConditionRepository) ReservationConditionService {
	return &reservationConditionService{
		reservationConditionRepository: reservationConditionRepo,
	}
}

func (u *reservationConditionService) CreateReservationCondition(reservationCondition dto.CreateReservationCondition, ctx *gin.Context) (entity.ReservationCondition, error) {
	reservationConditionToCreate := entity.ReservationCondition{}
	if reservationCondition.Category == "Makan" {
		reservationConditionToCreate.Title = "Pemberian Makan"
	}
	if reservationCondition.Category == "Bermain" {
		reservationConditionToCreate.Title = "Kegiatan Bermain"
	}

	reservationConditionToCreate.Description = reservationCondition.Description

	if reservationCondition.Image != nil {
		extension := filepath.Ext(reservationCondition.Image.Filename) // Generate random file name for the new uploaded file, so it doesn't override the old file with same name
		newImage := uuid.New().String() + extension
		reservationConditionToCreate.Image = newImage

		err := ctx.SaveUploadedFile(reservationCondition.Image, "uploads/user-profiles/"+newImage)
		if err != nil {
			return reservationConditionToCreate, err
		}
	}

	//errMap := smapping.FillStruct(&reservationConditionToCreate, smapping.MapFields(&reservationCondition))
	//if errMap != nil {
	//	return reservationConditionToCreate, errMap
	//}
	
	updatedReservationCondition, err := u.reservationConditionRepository.InsertReservationCondition(reservationConditionToCreate)
	return updatedReservationCondition, err
}
func (u *reservationConditionService) UpdateReservationCondition(reservationCondition dto.UpdateReservationCondition) (entity.ReservationCondition, error) {
	reservationConditionToUpdate := entity.ReservationCondition{}
	errMap := smapping.FillStruct(&reservationConditionToUpdate, smapping.MapFields(&reservationCondition))
	if errMap != nil {
		return reservationConditionToUpdate, errMap
	}
	updatedReservationCondition, err := u.reservationConditionRepository.UpdateReservationCondition(reservationConditionToUpdate)
	return updatedReservationCondition, err
}
func (u *reservationConditionService) DeleteReservationCondition(reservationConditionID string) error {
	reservationCondition, err := u.reservationConditionRepository.FindReservationConditionByID(reservationConditionID)
	if err != nil {
		return err
	}
	errDel := u.reservationConditionRepository.DeleteReservationCondition(reservationCondition)
	return errDel
}
func (u *reservationConditionService) ShowReservationCondition(reservationConditionID string) (entity.ReservationCondition, error) {
	result, err := u.reservationConditionRepository.FindReservationConditionByID(reservationConditionID)
	return result, err
}
func (u *reservationConditionService) GetAllReservationCondition(hotelID string, filterPaginate dto.FilterPagination) ([]entity.ReservationCondition, dto.Pagination, error) {
	admins, pagination, err := u.reservationConditionRepository.GetAllReservationCondition(hotelID, filterPaginate)
	return admins, pagination, err
}
