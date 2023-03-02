package service

import (
	"github.com/mashingan/smapping"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"github.com/yusufbagussh/pet_hotel_backend/repository"
)

// ReservationServiceService is a contract of what reservationServiceService can do
type ReservationServiceService interface {
	GetAllReservationService(filterPage dto.FilterPagination) ([]entity.ReservationService, dto.Pagination, error)
	CreateReservationService(reservationService dto.CreateReservationService) (entity.ReservationService, error)
	UpdateReservationService(reservationService dto.UpdateReservationService) (entity.ReservationService, error)
	DeleteReservationService(reservationServiceID string) error
	ShowReservationService(reservationServiceID string) (entity.ReservationService, error)
}

type reservationServiceService struct {
	reservationServiceRepository repository.ReservationServiceRepository
}

// NewReservationServiceService creates a new instance of ReservationServiceService
func NewReservationServiceService(reservationServiceRepo repository.ReservationServiceRepository) ReservationServiceService {
	return &reservationServiceService{
		reservationServiceRepository: reservationServiceRepo,
	}
}

func (u *reservationServiceService) CreateReservationService(reservationService dto.CreateReservationService) (entity.ReservationService, error) {
	reservationServiceToCreate := entity.ReservationService{}
	errMap := smapping.FillStruct(&reservationServiceToCreate, smapping.MapFields(&reservationService))
	if errMap != nil {
		return reservationServiceToCreate, errMap
	}
	updatedReservationService, err := u.reservationServiceRepository.InsertReservationService(reservationServiceToCreate)
	return updatedReservationService, err
}
func (u *reservationServiceService) UpdateReservationService(reservationService dto.UpdateReservationService) (entity.ReservationService, error) {
	reservationServiceToUpdate := entity.ReservationService{}
	errMap := smapping.FillStruct(&reservationServiceToUpdate, smapping.MapFields(&reservationService))
	if errMap != nil {
		return reservationServiceToUpdate, errMap
	}
	updatedReservationService, err := u.reservationServiceRepository.UpdateReservationService(reservationServiceToUpdate)
	return updatedReservationService, err
}
func (u *reservationServiceService) DeleteReservationService(reservationServiceID string) error {
	reservationService, err := u.reservationServiceRepository.FindReservationServiceByID(reservationServiceID)
	if err != nil {
		return err
	}
	errDel := u.reservationServiceRepository.DeleteReservationService(reservationService)
	return errDel
}
func (u *reservationServiceService) ShowReservationService(reservationServiceID string) (entity.ReservationService, error) {
	result, err := u.reservationServiceRepository.FindReservationServiceByID(reservationServiceID)
	return result, err
}
func (u *reservationServiceService) GetAllReservationService(filterPage dto.FilterPagination) ([]entity.ReservationService, dto.Pagination, error) {
	admins, pagination, err := u.reservationServiceRepository.GetAllReservationService(filterPage)
	return admins, pagination, err
}
