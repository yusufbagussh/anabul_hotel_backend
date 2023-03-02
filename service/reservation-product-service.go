package service

import (
	"github.com/mashingan/smapping"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"github.com/yusufbagussh/pet_hotel_backend/repository"
)

// ReservationProductService is a contract of what reservationProductService can do
type ReservationProductService interface {
	GetAllReservationProduct(filterPage dto.FilterPagination) ([]entity.ReservationProduct, dto.Pagination, error)
	CreateReservationProduct(reservationProduct dto.CreateReservationProduct) (entity.ReservationProduct, error)
	UpdateReservationProduct(reservationProduct dto.UpdateReservationProduct) (entity.ReservationProduct, error)
	DeleteReservationProduct(reservationProductID string) error
	ShowReservationProduct(reservationProductID string) (entity.ReservationProduct, error)
}

type reservationProductService struct {
	reservationProductRepository repository.ReservationProductRepository
}

// NewReservationProductService creates a new instance of ReservationProductService
func NewReservationProductService(reservationProductRepo repository.ReservationProductRepository) ReservationProductService {
	return &reservationProductService{
		reservationProductRepository: reservationProductRepo,
	}
}

func (u *reservationProductService) CreateReservationProduct(reservationProduct dto.CreateReservationProduct) (entity.ReservationProduct, error) {
	reservationProductToCreate := entity.ReservationProduct{}
	errMap := smapping.FillStruct(&reservationProductToCreate, smapping.MapFields(&reservationProduct))
	if errMap != nil {
		return reservationProductToCreate, errMap
	}
	updatedReservationProduct, err := u.reservationProductRepository.InsertReservationProduct(reservationProductToCreate)
	return updatedReservationProduct, err
}
func (u *reservationProductService) UpdateReservationProduct(reservationProduct dto.UpdateReservationProduct) (entity.ReservationProduct, error) {
	reservationProductToUpdate := entity.ReservationProduct{}
	errMap := smapping.FillStruct(&reservationProductToUpdate, smapping.MapFields(&reservationProduct))
	if errMap != nil {
		return reservationProductToUpdate, errMap
	}
	updatedReservationProduct, err := u.reservationProductRepository.UpdateReservationProduct(reservationProductToUpdate)
	return updatedReservationProduct, err
}
func (u *reservationProductService) DeleteReservationProduct(reservationProductID string) error {
	reservationProduct, err := u.reservationProductRepository.FindReservationProductByID(reservationProductID)
	if err != nil {
		return err
	}
	errDel := u.reservationProductRepository.DeleteReservationProduct(reservationProduct)
	return errDel
}
func (u *reservationProductService) ShowReservationProduct(reservationProductID string) (entity.ReservationProduct, error) {
	result, err := u.reservationProductRepository.FindReservationProductByID(reservationProductID)
	return result, err
}
func (u *reservationProductService) GetAllReservationProduct(filterPage dto.FilterPagination) ([]entity.ReservationProduct, dto.Pagination, error) {
	admins, pagination, err := u.reservationProductRepository.GetAllReservationProduct(filterPage)
	return admins, pagination, err
}
