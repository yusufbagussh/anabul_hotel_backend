package service

import (
	"github.com/mashingan/smapping"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"github.com/yusufbagussh/pet_hotel_backend/repository"
)

// DistrictService is a contract of what districtService can do
type DistrictService interface {
	GetAllDistrict(filterPagination dto.FilterPagination) ([]entity.District, dto.Pagination, error)
	CreateDistrict(district dto.CreateDistrict) (entity.District, error)
	UpdateDistrict(district dto.UpdateDistrict) (entity.District, error)
	DeleteDistrict(districtID string) error
	ShowDistrict(districtID string) (entity.District, error)
}

type districtService struct {
	districtRepository repository.DistrictRepository
}

// NewDistrictService creates a new instance of DistrictService
func NewDistrictService(districtRepo repository.DistrictRepository) DistrictService {
	return &districtService{
		districtRepository: districtRepo,
	}
}

func (u *districtService) CreateDistrict(district dto.CreateDistrict) (entity.District, error) {
	districtToCreate := entity.District{}
	errMap := smapping.FillStruct(&districtToCreate, smapping.MapFields(&district))
	if errMap != nil {
		return districtToCreate, errMap
	}
	updatedDistrict, err := u.districtRepository.InsertDistrict(districtToCreate)
	return updatedDistrict, err
}
func (u *districtService) UpdateDistrict(district dto.UpdateDistrict) (entity.District, error) {
	districtToUpdate := entity.District{}
	errMap := smapping.FillStruct(&districtToUpdate, smapping.MapFields(&district))
	if errMap != nil {
		return districtToUpdate, errMap
	}
	updatedDistrict, err := u.districtRepository.UpdateDistrict(districtToUpdate)
	return updatedDistrict, err
}
func (u *districtService) DeleteDistrict(districtID string) error {
	district, err := u.districtRepository.FindDistrictByID(districtID)
	if err != nil {
		return err
	}
	errDel := u.districtRepository.DeleteDistrict(district)
	return errDel
}
func (u *districtService) ShowDistrict(districtID string) (entity.District, error) {
	result, err := u.districtRepository.FindDistrictByID(districtID)
	return result, err
}
func (u *districtService) GetAllDistrict(filterPaginate dto.FilterPagination) ([]entity.District, dto.Pagination, error) {
	admins, pagination, err := u.districtRepository.GetAllDistrict(filterPaginate)
	return admins, pagination, err
}
