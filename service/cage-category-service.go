package service

import (
	"github.com/mashingan/smapping"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"github.com/yusufbagussh/pet_hotel_backend/repository"
)

// CageCategoryService is a contract of what cageCategoryService can do
type CageCategoryService interface {
	GetAllCageCategory(hotelID string, filterPage dto.FilterPagination) ([]entity.CageCategory, dto.Pagination, error)
	CreateCageCategory(cageCategory dto.CreateCageCategory) (entity.CageCategory, error)
	UpdateCageCategory(cageCategory dto.UpdateCageCategory) (entity.CageCategory, error)
	DeleteCageCategory(cageCategoryID string, userHotelID string) (error, interface{})
	ShowCageCategory(cageCategoryID string, userHotelID string) (entity.CageCategory, error, interface{})
}

type cageCategoryService struct {
	cageCategoryRepository repository.CageCategoryRepository
}

// NewCageCategoryService creates a new instance of CageCategoryService
func NewCageCategoryService(cageCategoryRepo repository.CageCategoryRepository) CageCategoryService {
	return &cageCategoryService{
		cageCategoryRepository: cageCategoryRepo,
	}
}

func (u *cageCategoryService) CreateCageCategory(cageCategory dto.CreateCageCategory) (entity.CageCategory, error) {
	cageCategoryToCreate := entity.CageCategory{}
	errMap := smapping.FillStruct(&cageCategoryToCreate, smapping.MapFields(&cageCategory))
	if errMap != nil {
		return cageCategoryToCreate, errMap
	}
	updatedCageCategory, err := u.cageCategoryRepository.InsertCageCategory(cageCategoryToCreate)
	return updatedCageCategory, err
}
func (u *cageCategoryService) UpdateCageCategory(cageCategory dto.UpdateCageCategory) (entity.CageCategory, error) {
	cageCategoryToUpdate := entity.CageCategory{}
	errMap := smapping.FillStruct(&cageCategoryToUpdate, smapping.MapFields(&cageCategory))
	if errMap != nil {
		return cageCategoryToUpdate, errMap
	}
	updatedCageCategory, err := u.cageCategoryRepository.UpdateCageCategory(cageCategoryToUpdate)
	return updatedCageCategory, err
}
func (u *cageCategoryService) DeleteCageCategory(cageCategoryID string, userHotelID string) (error, interface{}) {
	cageCategory, err := u.cageCategoryRepository.FindCageCategoryByID(cageCategoryID)
	if cageCategory.HotelID != userHotelID {
		return nil, false
	}
	if err != nil {
		return err, nil
	}
	errDel := u.cageCategoryRepository.DeleteCageCategory(cageCategory)
	return errDel, nil
}
func (u *cageCategoryService) ShowCageCategory(cageCategoryID string, userHotelID string) (entity.CageCategory, error, interface{}) {
	result, err := u.cageCategoryRepository.FindCageCategoryByID(cageCategoryID)
	if result.HotelID != userHotelID {
		return entity.CageCategory{}, nil, false
	}
	return result, err, nil
}
func (u *cageCategoryService) GetAllCageCategory(hotelID string, filterPage dto.FilterPagination) ([]entity.CageCategory, dto.Pagination, error) {
	admins, pagination, err := u.cageCategoryRepository.GetAllCageCategory(hotelID, filterPage)
	return admins, pagination, err
}
