package service

import (
	"github.com/mashingan/smapping"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"github.com/yusufbagussh/pet_hotel_backend/repository"
)

// ResponseService is a contract of what responseService can do
type ResponseService interface {
	GetAllResponse(hotelID string, filterPagination dto.FilterPagination) ([]entity.Response, dto.Pagination, error)
	CreateResponse(response dto.CreateResponse) (entity.Response, error)
	UpdateResponse(response dto.UpdateResponse) (entity.Response, error)
	DeleteResponse(responseID string) error
	ShowResponse(responseID string) (entity.Response, error)
}

type responseService struct {
	responseRepository repository.ResponseRepository
}

// NewResponseService creates a new instance of ResponseService
func NewResponseService(responseRepo repository.ResponseRepository) ResponseService {
	return &responseService{
		responseRepository: responseRepo,
	}
}

func (u *responseService) CreateResponse(response dto.CreateResponse) (entity.Response, error) {
	responseToCreate := entity.Response{}
	errMap := smapping.FillStruct(&responseToCreate, smapping.MapFields(&response))
	if errMap != nil {
		return responseToCreate, errMap
	}
	updatedResponse, err := u.responseRepository.InsertResponse(responseToCreate)
	return updatedResponse, err
}
func (u *responseService) UpdateResponse(response dto.UpdateResponse) (entity.Response, error) {
	responseToUpdate := entity.Response{}
	errMap := smapping.FillStruct(&responseToUpdate, smapping.MapFields(&response))
	if errMap != nil {
		return responseToUpdate, errMap
	}
	updatedResponse, err := u.responseRepository.UpdateResponse(responseToUpdate)
	return updatedResponse, err
}
func (u *responseService) DeleteResponse(responseID string) error {
	response, err := u.responseRepository.FindResponseByID(responseID)
	if err != nil {
		return err
	}
	errDel := u.responseRepository.DeleteResponse(response)
	return errDel
}
func (u *responseService) ShowResponse(responseID string) (entity.Response, error) {
	result, err := u.responseRepository.FindResponseByID(responseID)
	return result, err
}
func (u *responseService) GetAllResponse(hotelID string, filterPaginate dto.FilterPagination) ([]entity.Response, dto.Pagination, error) {
	admins, pagination, err := u.responseRepository.GetAllResponse(hotelID, filterPaginate)
	return admins, pagination, err
}
