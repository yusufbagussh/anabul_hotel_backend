package service

import (
	"github.com/mashingan/smapping"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"github.com/yusufbagussh/pet_hotel_backend/repository"
)

// ReservationInventoryService is a contract of what reservationInventoryService can do
type ReservationInventoryService interface {
	GetAllReservationInventory(filterPage dto.FilterPagination) ([]entity.ReservationInventory, dto.Pagination, error)
	CreateReservationInventory(reservationInventory dto.CreateReservationInventory) (entity.ReservationInventory, error)
	UpdateReservationInventory(reservationInventory dto.UpdateReservationInventory) (entity.ReservationInventory, error)
	DeleteReservationInventory(reservationInventoryID string) error
	ShowReservationInventory(reservationInventoryID string) (entity.ReservationInventory, error)
}

type reservationInventoryService struct {
	reservationInventoryRepository repository.ReservationInventoryRepository
}

// NewReservationInventoryService creates a new instance of ReservationInventoryService
func NewReservationInventoryService(reservationInventoryRepo repository.ReservationInventoryRepository) ReservationInventoryService {
	return &reservationInventoryService{
		reservationInventoryRepository: reservationInventoryRepo,
	}
}

func (u *reservationInventoryService) CreateReservationInventory(reservationInventory dto.CreateReservationInventory) (entity.ReservationInventory, error) {
	reservationInventoryToCreate := entity.ReservationInventory{}
	errMap := smapping.FillStruct(&reservationInventoryToCreate, smapping.MapFields(&reservationInventory))
	if errMap != nil {
		return reservationInventoryToCreate, errMap
	}
	updatedReservationInventory, err := u.reservationInventoryRepository.InsertReservationInventory(reservationInventoryToCreate)
	return updatedReservationInventory, err
}
func (u *reservationInventoryService) UpdateReservationInventory(reservationInventory dto.UpdateReservationInventory) (entity.ReservationInventory, error) {
	reservationInventoryToUpdate := entity.ReservationInventory{}
	errMap := smapping.FillStruct(&reservationInventoryToUpdate, smapping.MapFields(&reservationInventory))
	if errMap != nil {
		return reservationInventoryToUpdate, errMap
	}
	updatedReservationInventory, err := u.reservationInventoryRepository.UpdateReservationInventory(reservationInventoryToUpdate)
	return updatedReservationInventory, err
}
func (u *reservationInventoryService) DeleteReservationInventory(reservationInventoryID string) error {
	reservationInventory, err := u.reservationInventoryRepository.FindReservationInventoryByID(reservationInventoryID)
	if err != nil {
		return err
	}
	errDel := u.reservationInventoryRepository.DeleteReservationInventory(reservationInventory)
	return errDel
}
func (u *reservationInventoryService) ShowReservationInventory(reservationInventoryID string) (entity.ReservationInventory, error) {
	result, err := u.reservationInventoryRepository.FindReservationInventoryByID(reservationInventoryID)
	return result, err
}
func (u *reservationInventoryService) GetAllReservationInventory(filterPage dto.FilterPagination) ([]entity.ReservationInventory, dto.Pagination, error) {
	admins, pagination, err := u.reservationInventoryRepository.GetAllReservationInventory(filterPage)
	return admins, pagination, err
}
