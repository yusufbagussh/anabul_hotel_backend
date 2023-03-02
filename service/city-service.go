package service

import (
	"github.com/mashingan/smapping"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"github.com/yusufbagussh/pet_hotel_backend/repository"
)

// CityService is a contract of what cityService can do
type CityService interface {
	GetAllCity(filterPagination dto.FilterPagination) ([]entity.City, dto.Pagination, error)
	CreateCity(city dto.CreateCity) (entity.City, error)
	UpdateCity(city dto.UpdateCity) (entity.City, error)
	DeleteCity(cityID string) error
	ShowCity(cityID string) (entity.City, error)
}

type cityService struct {
	cityRepository repository.CityRepository
}

// NewCityService creates a new instance of CityService
func NewCityService(cityRepo repository.CityRepository) CityService {
	return &cityService{
		cityRepository: cityRepo,
	}
}

func (u *cityService) CreateCity(city dto.CreateCity) (entity.City, error) {
	cityToCreate := entity.City{}
	errMap := smapping.FillStruct(&cityToCreate, smapping.MapFields(&city))
	if errMap != nil {
		return cityToCreate, errMap
	}
	updatedCity, err := u.cityRepository.InsertCity(cityToCreate)
	return updatedCity, err
}
func (u *cityService) UpdateCity(city dto.UpdateCity) (entity.City, error) {
	cityToUpdate := entity.City{}
	errMap := smapping.FillStruct(&cityToUpdate, smapping.MapFields(&city))
	if errMap != nil {
		return cityToUpdate, errMap
	}
	updatedCity, err := u.cityRepository.UpdateCity(cityToUpdate)
	return updatedCity, err
}
func (u *cityService) DeleteCity(cityID string) error {
	city, err := u.cityRepository.FindCityByID(cityID)
	if err != nil {
		return err
	}
	errDel := u.cityRepository.DeleteCity(city)
	return errDel
}
func (u *cityService) ShowCity(cityID string) (entity.City, error) {
	result, err := u.cityRepository.FindCityByID(cityID)
	return result, err
}
func (u *cityService) GetAllCity(filterPaginate dto.FilterPagination) ([]entity.City, dto.Pagination, error) {
	admins, pagination, err := u.cityRepository.GetAllCity(filterPaginate)
	return admins, pagination, err
}
