package service

import (
	"github.com/mashingan/smapping"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"github.com/yusufbagussh/pet_hotel_backend/repository"
)

// CageDetailService is a contract of what cageDetailService can do
type CageDetailService interface {
	GetAllCageDetail(hotelID string, filterPage dto.FilterPagination) ([]entity.CageDetail, dto.Pagination, error)
	CreateCageDetail(cageDetail dto.CreateCageDetail) (entity.CageDetail, error)
	UpdateCageDetail(cageDetail dto.UpdateCageDetail) (entity.CageDetail, error)
	DeleteCageDetail(cageDetailID string, userHotelID string) (error, interface{})
	ShowCageDetail(cageDetailID string, userHotelID string) (entity.CageDetail, error, interface{})
}

type cageDetailService struct {
	cageDetailRepository repository.CageDetailRepository
}

// NewCageDetailService creates a new instance of CageDetailService
func NewCageDetailService(cageDetailRepo repository.CageDetailRepository) CageDetailService {
	return &cageDetailService{
		cageDetailRepository: cageDetailRepo,
	}
}

func (u *cageDetailService) CreateCageDetail(cageDetail dto.CreateCageDetail) (entity.CageDetail, error) {
	cageDetailToCreate := entity.CageDetail{}
	errMap := smapping.FillStruct(&cageDetailToCreate, smapping.MapFields(&cageDetail))
	if errMap != nil {
		return cageDetailToCreate, errMap
	}
	updatedCageDetail, err := u.cageDetailRepository.InsertCageDetail(cageDetailToCreate)
	return updatedCageDetail, err
}
func (u *cageDetailService) UpdateCageDetail(cageDetail dto.UpdateCageDetail) (entity.CageDetail, error) {
	cageDetailToUpdate := entity.CageDetail{}
	errMap := smapping.FillStruct(&cageDetailToUpdate, smapping.MapFields(&cageDetail))
	if errMap != nil {
		return cageDetailToUpdate, errMap
	}
	updatedCageDetail, err := u.cageDetailRepository.UpdateCageDetail(cageDetailToUpdate)
	return updatedCageDetail, err
}
func (u *cageDetailService) DeleteCageDetail(cageDetailID string, userHotelID string) (error, interface{}) {
	cageDetail, err := u.cageDetailRepository.FindCageDetailByID(cageDetailID)
	if cageDetail.HotelID != userHotelID {
		return nil, false
	}
	if err != nil {
		return err, nil
	}
	errDel := u.cageDetailRepository.DeleteCageDetail(cageDetail)
	return errDel, nil
}
func (u *cageDetailService) ShowCageDetail(cageDetailID string, userHotelID string) (entity.CageDetail, error, interface{}) {
	result, err := u.cageDetailRepository.FindCageDetailByID(cageDetailID)
	if result.HotelID != userHotelID {
		return entity.CageDetail{}, nil, false
	}
	return result, err, nil
}
func (u *cageDetailService) GetAllCageDetail(hotelID string, filterPage dto.FilterPagination) ([]entity.CageDetail, dto.Pagination, error) {
	admins, pagination, err := u.cageDetailRepository.GetAllCageDetail(hotelID, filterPage)
	return admins, pagination, err
}
