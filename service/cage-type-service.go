package service

import (
	"github.com/mashingan/smapping"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"github.com/yusufbagussh/pet_hotel_backend/repository"
)

// CageTypeService is a contract of what cageTypeService can do
type CageTypeService interface {
	GetAllCageType(hotelID string, filterPage dto.FilterPagination) ([]entity.CageType, dto.Pagination, error)
	CreateCageType(cageType dto.CreateCageType) (entity.CageType, error)
	UpdateCageType(cageType dto.UpdateCageType) (entity.CageType, error)
	DeleteCageType(cageTypeID string, userHotelID string) (error, interface{})
	ShowCageType(cageTypeID string, userHotelID string) (entity.CageType, error, interface{})
}

type cageTypeService struct {
	cageTypeRepository repository.CageTypeRepository
}

// NewCageTypeService creates a new instance of CageTypeService
func NewCageTypeService(cageTypeRepo repository.CageTypeRepository) CageTypeService {
	return &cageTypeService{
		cageTypeRepository: cageTypeRepo,
	}
}

func (u *cageTypeService) CreateCageType(cageType dto.CreateCageType) (entity.CageType, error) {
	cageTypeToCreate := entity.CageType{}
	errMap := smapping.FillStruct(&cageTypeToCreate, smapping.MapFields(&cageType))
	if errMap != nil {
		return cageTypeToCreate, errMap
	}
	updatedCageType, err := u.cageTypeRepository.InsertCageType(cageTypeToCreate)
	return updatedCageType, err
}
func (u *cageTypeService) UpdateCageType(cageType dto.UpdateCageType) (entity.CageType, error) {
	cageTypeToUpdate := entity.CageType{}
	errMap := smapping.FillStruct(&cageTypeToUpdate, smapping.MapFields(&cageType))
	if errMap != nil {
		return cageTypeToUpdate, errMap
	}
	updatedCageType, err := u.cageTypeRepository.UpdateCageType(cageTypeToUpdate)
	return updatedCageType, err
}
func (u *cageTypeService) DeleteCageType(cageTypeID string, userHotelID string) (error, interface{}) {
	cageType, err := u.cageTypeRepository.FindCageTypeByID(cageTypeID)
	if cageType.HotelID != userHotelID {
		return nil, false
	}
	if err != nil {
		return err, nil
	}
	errDel := u.cageTypeRepository.DeleteCageType(cageType)
	return errDel, nil
}
func (u *cageTypeService) ShowCageType(cageTypeID string, userHotelID string) (entity.CageType, error, interface{}) {
	result, err := u.cageTypeRepository.FindCageTypeByID(cageTypeID)
	if result.HotelID != userHotelID {
		return entity.CageType{}, nil, false
	}
	return result, err, nil
}
func (u *cageTypeService) GetAllCageType(hotelID string, filterPage dto.FilterPagination) ([]entity.CageType, dto.Pagination, error) {
	admins, pagination, err := u.cageTypeRepository.GetAllCageType(hotelID, filterPage)
	return admins, pagination, err
}
