package service

import (
	"github.com/mashingan/smapping"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"github.com/yusufbagussh/pet_hotel_backend/repository"
)

// RateService is a contract of what rateService can do
type RateService interface {
	GetAllRate(hotelID string, filterPagination dto.FilterPagination) ([]entity.Rate, dto.Pagination, error)
	CreateRate(rate dto.CreateRate) (entity.Rate, error)
	UpdateRate(rate dto.UpdateRate) (entity.Rate, error)
	DeleteRate(rateID string) error
	ShowRate(rateID string) (entity.Rate, error)
}

type rateService struct {
	rateRepository repository.RateRepository
}

// NewRateService creates a new instance of RateService
func NewRateService(rateRepo repository.RateRepository) RateService {
	return &rateService{
		rateRepository: rateRepo,
	}
}

func (u *rateService) CreateRate(rate dto.CreateRate) (entity.Rate, error) {
	rateToCreate := entity.Rate{}
	errMap := smapping.FillStruct(&rateToCreate, smapping.MapFields(&rate))
	if errMap != nil {
		return rateToCreate, errMap
	}
	updatedRate, err := u.rateRepository.InsertRate(rateToCreate)
	return updatedRate, err
}
func (u *rateService) UpdateRate(rate dto.UpdateRate) (entity.Rate, error) {
	rateToUpdate := entity.Rate{}
	errMap := smapping.FillStruct(&rateToUpdate, smapping.MapFields(&rate))
	if errMap != nil {
		return rateToUpdate, errMap
	}
	updatedRate, err := u.rateRepository.UpdateRate(rateToUpdate)
	return updatedRate, err
}
func (u *rateService) DeleteRate(rateID string) error {
	rate, err := u.rateRepository.FindRateByID(rateID)
	if err != nil {
		return err
	}
	errDel := u.rateRepository.DeleteRate(rate)
	return errDel
}
func (u *rateService) ShowRate(rateID string) (entity.Rate, error) {
	result, err := u.rateRepository.FindRateByID(rateID)
	return result, err
}
func (u *rateService) GetAllRate(hotelID string, filterPaginate dto.FilterPagination) ([]entity.Rate, dto.Pagination, error) {
	admins, pagination, err := u.rateRepository.GetAllRate(hotelID, filterPaginate)
	return admins, pagination, err
}
