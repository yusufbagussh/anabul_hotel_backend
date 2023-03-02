package service

import (
	"github.com/mashingan/smapping"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"github.com/yusufbagussh/pet_hotel_backend/repository"
)

// ServiceService is a contract of what serviceService can do
type ServiceService interface {
	GetAllService(hotelID string, filterPage dto.FilterPagination) ([]entity.Service, dto.Pagination, error)
	CreateService(service dto.CreateService) (entity.Service, error)
	UpdateService(service dto.UpdateService) (entity.Service, error)
	DeleteService(serviceID string, userHotelID string) (error, interface{})
	ShowService(serviceID string, userHotelID string) (entity.Service, error, interface{})
}

type serviceService struct {
	serviceRepository repository.ServiceRepository
}

// NewServiceService creates a new instance of ServiceService
func NewServiceService(serviceRepo repository.ServiceRepository) ServiceService {
	return &serviceService{
		serviceRepository: serviceRepo,
	}
}

func (u *serviceService) CreateService(service dto.CreateService) (entity.Service, error) {
	serviceToCreate := entity.Service{}
	errMap := smapping.FillStruct(&serviceToCreate, smapping.MapFields(&service))
	if errMap != nil {
		return serviceToCreate, errMap
	}
	updatedService, err := u.serviceRepository.InsertService(serviceToCreate)
	return updatedService, err
}
func (u *serviceService) UpdateService(service dto.UpdateService) (entity.Service, error) {
	serviceToUpdate := entity.Service{}
	errMap := smapping.FillStruct(&serviceToUpdate, smapping.MapFields(&service))
	if errMap != nil {
		return serviceToUpdate, errMap
	}
	updatedService, err := u.serviceRepository.UpdateService(serviceToUpdate)
	return updatedService, err
}
func (u *serviceService) DeleteService(serviceID string, userHotelID string) (error, interface{}) {
	service, err := u.serviceRepository.FindServiceByID(serviceID)
	if service.HotelID != userHotelID {
		return nil, false
	}
	if err != nil {
		return err, nil
	}
	errDel := u.serviceRepository.DeleteService(service)
	return errDel, nil
}
func (u *serviceService) ShowService(serviceID string, userHotelID string) (entity.Service, error, interface{}) {
	result, err := u.serviceRepository.FindServiceByID(serviceID)
	if result.HotelID != userHotelID {
		return entity.Service{}, nil, false
	}
	return result, err, nil
}
func (u *serviceService) GetAllService(hotelID string, filterPage dto.FilterPagination) ([]entity.Service, dto.Pagination, error) {
	admins, pagination, err := u.serviceRepository.GetAllService(hotelID, filterPage)
	return admins, pagination, err
}
