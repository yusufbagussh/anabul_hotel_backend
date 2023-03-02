package service

import (
	"github.com/mashingan/smapping"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"github.com/yusufbagussh/pet_hotel_backend/repository"
)

// ReservationDetailService is a contract of what reservationDetailService can do
type ReservationDetailService interface {
	GetAllReservationDetail(filterPage dto.FilterPagination) ([]entity.ReservationDetail, dto.Pagination, error)
	CreateReservationDetail(reservationDetail dto.CreateReservationDetail) (entity.ReservationDetail, error)
	UpdateReservationDetail(reservationDetail dto.UpdateReservationDetail) (entity.ReservationDetail, error)
	DeleteReservationDetail(reservationDetailID string) error
	ShowReservationDetail(reservationDetailID string) (entity.ReservationDetail, error)
}

type reservationDetailService struct {
	reservationDetailRepository repository.ReservationDetailRepository
}

// NewReservationDetailService creates a new instance of ReservationDetailService
func NewReservationDetailService(reservationDetailRepo repository.ReservationDetailRepository) ReservationDetailService {
	return &reservationDetailService{
		reservationDetailRepository: reservationDetailRepo,
	}
}

func (u *reservationDetailService) CreateReservationDetail(reservationDetail dto.CreateReservationDetail) (entity.ReservationDetail, error) {
	reservationDetailToCreate := entity.ReservationDetail{}
	errMap := smapping.FillStruct(&reservationDetailToCreate, smapping.MapFields(&reservationDetail))
	if errMap != nil {
		return reservationDetailToCreate, errMap
	}
	updatedReservationDetail, err := u.reservationDetailRepository.InsertReservationDetail(reservationDetailToCreate)
	return updatedReservationDetail, err
}
func (u *reservationDetailService) UpdateReservationDetail(reservationDetail dto.UpdateReservationDetail) (entity.ReservationDetail, error) {
	reservationDetailToUpdate := entity.ReservationDetail{}
	errMap := smapping.FillStruct(&reservationDetailToUpdate, smapping.MapFields(&reservationDetail))
	if errMap != nil {
		return reservationDetailToUpdate, errMap
	}
	updatedReservationDetail, err := u.reservationDetailRepository.UpdateReservationDetail(reservationDetailToUpdate)
	return updatedReservationDetail, err
}
func (u *reservationDetailService) DeleteReservationDetail(reservationDetailID string) error {
	reservationDetail, err := u.reservationDetailRepository.FindReservationDetailByID(reservationDetailID)
	if err != nil {
		return err
	}
	errDel := u.reservationDetailRepository.DeleteReservationDetail(reservationDetail)
	return errDel
}
func (u *reservationDetailService) ShowReservationDetail(reservationDetailID string) (entity.ReservationDetail, error) {
	result, err := u.reservationDetailRepository.FindReservationDetailByID(reservationDetailID)
	return result, err
}
func (u *reservationDetailService) GetAllReservationDetail(filterPage dto.FilterPagination) ([]entity.ReservationDetail, dto.Pagination, error) {
	admins, pagination, err := u.reservationDetailRepository.GetAllReservationDetail(filterPage)
	return admins, pagination, err
}
