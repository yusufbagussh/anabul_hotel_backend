package service

import (
	"github.com/mashingan/smapping"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"github.com/yusufbagussh/pet_hotel_backend/repository"
)

// ProvinceService is a contract of what provinceService can do
type ProvinceService interface {
	GetAllProvince(filterPage dto.FilterPagination) ([]entity.Province, dto.Pagination, error)
	CreateProvince(province dto.CreateProvince) (entity.Province, error)
	UpdateProvince(province dto.UpdateProvince) (entity.Province, error)
	DeleteProvince(provinceID string) error
	ShowProvince(provinceID string) (entity.Province, error)
}

type provinceService struct {
	provinceRepository repository.ProvinceRepository
}

// NewProvinceService creates a new instance of ProvinceService
func NewProvinceService(provinceRepo repository.ProvinceRepository) ProvinceService {
	return &provinceService{
		provinceRepository: provinceRepo,
	}
}

func (u *provinceService) CreateProvince(province dto.CreateProvince) (entity.Province, error) {
	provinceToCreate := entity.Province{}
	errMap := smapping.FillStruct(&provinceToCreate, smapping.MapFields(&province))
	if errMap != nil {
		return provinceToCreate, errMap
	}
	updatedProvince, err := u.provinceRepository.InsertProvince(provinceToCreate)
	return updatedProvince, err
}
func (u *provinceService) UpdateProvince(province dto.UpdateProvince) (entity.Province, error) {
	provinceToUpdate := entity.Province{}
	errMap := smapping.FillStruct(&provinceToUpdate, smapping.MapFields(&province))
	if errMap != nil {
		return provinceToUpdate, errMap
	}
	updatedProvince, err := u.provinceRepository.UpdateProvince(provinceToUpdate)
	return updatedProvince, err
}
func (u *provinceService) DeleteProvince(provinceID string) error {
	province, err := u.provinceRepository.FindProvinceByID(provinceID)
	if err != nil {
		return err
	}
	errDel := u.provinceRepository.DeleteProvince(province)
	return errDel
}
func (u *provinceService) ShowProvince(provinceID string) (entity.Province, error) {
	result, err := u.provinceRepository.FindProvinceByID(provinceID)
	return result, err
}
func (u *provinceService) GetAllProvince(filterPage dto.FilterPagination) ([]entity.Province, dto.Pagination, error) {
	admins, pagination, err := u.provinceRepository.GetAllProvince(filterPage)
	return admins, pagination, err
}
