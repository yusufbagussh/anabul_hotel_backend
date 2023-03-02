package service

import (
	"github.com/mashingan/smapping"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"github.com/yusufbagussh/pet_hotel_backend/repository"
)

// CageService is a contract of what cageService can do
type CageService interface {
	GetAllCage(hotelID string, filterPage dto.FilterPagination) ([]entity.Cage, dto.Pagination, error)
	CreateCage(cage dto.CreateCage) (entity.Cage, error)
	UpdateCage(cage dto.UpdateCage) (entity.Cage, error)
	DeleteCage(cageID string, userHotelID string) (error, interface{})
	ShowCage(cageID string, userHotelID string) (entity.Cage, error, interface{})
}

type cageService struct {
	cageRepository repository.CageRepository
}

// NewCageService creates a new instance of CageService
func NewCageService(cageRepo repository.CageRepository) CageService {
	return &cageService{
		cageRepository: cageRepo,
	}
}

func (u *cageService) CreateCage(cage dto.CreateCage) (entity.Cage, error) {
	cageToCreate := entity.Cage{}
	errMap := smapping.FillStruct(&cageToCreate, smapping.MapFields(&cage))
	if errMap != nil {
		return cageToCreate, errMap
	}
	updatedCage, err := u.cageRepository.InsertCage(cageToCreate)
	return updatedCage, err
}
func (u *cageService) UpdateCage(cage dto.UpdateCage) (entity.Cage, error) {
	cageToUpdate := entity.Cage{}
	errMap := smapping.FillStruct(&cageToUpdate, smapping.MapFields(&cage))
	if errMap != nil {
		return cageToUpdate, errMap
	}
	updatedCage, err := u.cageRepository.UpdateCage(cageToUpdate)
	return updatedCage, err
}
func (u *cageService) DeleteCage(cageID string, userHotelID string) (error, interface{}) {
	cage, err := u.cageRepository.FindCageByID(cageID)
	if cage.HotelID != userHotelID {
		return nil, false
	}
	if err != nil {
		return err, nil
	}
	errDel := u.cageRepository.DeleteCage(cage)
	return errDel, nil
}
func (u *cageService) ShowCage(cageID string, userHotelID string) (entity.Cage, error, interface{}) {
	result, err := u.cageRepository.FindCageByID(cageID)
	if result.HotelID != userHotelID {
		return entity.Cage{}, nil, false
	}
	return result, err, nil
}
func (u *cageService) GetAllCage(hotelID string, filterPage dto.FilterPagination) ([]entity.Cage, dto.Pagination, error) {
	admins, pagination, err := u.cageRepository.GetAllCage(hotelID, filterPage)
	return admins, pagination, err
}
