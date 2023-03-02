package service

import (
	"github.com/mashingan/smapping"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"github.com/yusufbagussh/pet_hotel_backend/repository"
)

// ServiceDetailService is a contract of what serviceDetailService can do
type ServiceDetailService interface {
	GetAllServiceDetail(hotelID string, filterPage dto.FilterPagination) ([]entity.ServiceDetail, dto.Pagination, error)
	CreateServiceDetail(serviceDetail dto.CreateServiceDetail) (entity.ServiceDetail, error)
	UpdateServiceDetail(serviceDetail dto.UpdateServiceDetail) (entity.ServiceDetail, error)
	DeleteServiceDetail(serviceDetailID string, userHotelID string) (error, interface{})
	ShowServiceDetail(serviceDetailID string, userHotelID string) (entity.ServiceDetail, error, interface{})
}

type serviceDetailService struct {
	serviceDetailRepository repository.ServiceDetailRepository
}

// NewServiceDetailService creates a new instance of ServiceDetailService
func NewServiceDetailService(serviceDetailRepo repository.ServiceDetailRepository) ServiceDetailService {
	return &serviceDetailService{
		serviceDetailRepository: serviceDetailRepo,
	}
}

func (u *serviceDetailService) CreateServiceDetail(serviceDetail dto.CreateServiceDetail) (entity.ServiceDetail, error) {
	serviceDetailToCreate := entity.ServiceDetail{}
	errMap := smapping.FillStruct(&serviceDetailToCreate, smapping.MapFields(&serviceDetail))
	if errMap != nil {
		return serviceDetailToCreate, errMap
	}
	updatedServiceDetail, err := u.serviceDetailRepository.InsertServiceDetail(serviceDetailToCreate)
	return updatedServiceDetail, err
}
func (u *serviceDetailService) UpdateServiceDetail(serviceDetail dto.UpdateServiceDetail) (entity.ServiceDetail, error) {
	serviceDetailToUpdate := entity.ServiceDetail{}
	errMap := smapping.FillStruct(&serviceDetailToUpdate, smapping.MapFields(&serviceDetail))
	if errMap != nil {
		return serviceDetailToUpdate, errMap
	}
	updatedServiceDetail, err := u.serviceDetailRepository.UpdateServiceDetail(serviceDetailToUpdate)
	return updatedServiceDetail, err
}
func (u *serviceDetailService) DeleteServiceDetail(serviceDetailID string, userHotelID string) (error, interface{}) {
	serviceDetail, err := u.serviceDetailRepository.FindServiceDetailByID(serviceDetailID)
	if serviceDetail.HotelID != userHotelID {
		return nil, false
	}
	if err != nil {
		return err, nil
	}
	errDel := u.serviceDetailRepository.DeleteServiceDetail(serviceDetail)
	return errDel, nil
}
func (u *serviceDetailService) ShowServiceDetail(serviceDetailID string, userHotelID string) (entity.ServiceDetail, error, interface{}) {
	result, err := u.serviceDetailRepository.FindServiceDetailByID(serviceDetailID)
	if result.HotelID != userHotelID {
		return entity.ServiceDetail{}, nil, false
	}
	return result, err, nil
}
func (u *serviceDetailService) GetAllServiceDetail(hotelID string, filterPage dto.FilterPagination) ([]entity.ServiceDetail, dto.Pagination, error) {
	admins, pagination, err := u.serviceDetailRepository.GetAllServiceDetail(hotelID, filterPage)
	return admins, pagination, err
}
