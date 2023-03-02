package service

import (
	"github.com/mashingan/smapping"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"github.com/yusufbagussh/pet_hotel_backend/repository"
)

// GroupDetailService is a contract of what groupDetailService can do
type GroupDetailService interface {
	GetAllGroupDetail(hotelID string, filterPage dto.FilterPagination) ([]entity.GroupDetail, dto.Pagination, error)
	CreateGroupDetail(groupDetail dto.CreateGroupDetail) (entity.GroupDetail, error)
	UpdateGroupDetail(groupDetail dto.UpdateGroupDetail) (entity.GroupDetail, error)
	DeleteGroupDetail(groupDetailID string, userHotelID string) (error, interface{})
	ShowGroupDetail(groupDetailID string, userHotelID string) (entity.GroupDetail, error, interface{})
}

type groupDetailService struct {
	groupDetailRepository repository.GroupDetailRepository
}

// NewGroupDetailService creates a new instance of GroupDetailService
func NewGroupDetailService(groupDetailRepo repository.GroupDetailRepository) GroupDetailService {
	return &groupDetailService{
		groupDetailRepository: groupDetailRepo,
	}
}

func (u *groupDetailService) CreateGroupDetail(groupDetail dto.CreateGroupDetail) (entity.GroupDetail, error) {
	groupDetailToCreate := entity.GroupDetail{}
	errMap := smapping.FillStruct(&groupDetailToCreate, smapping.MapFields(&groupDetail))
	if errMap != nil {
		return groupDetailToCreate, errMap
	}
	updatedGroupDetail, err := u.groupDetailRepository.InsertGroupDetail(groupDetailToCreate)
	return updatedGroupDetail, err
}
func (u *groupDetailService) UpdateGroupDetail(groupDetail dto.UpdateGroupDetail) (entity.GroupDetail, error) {
	groupDetailToUpdate := entity.GroupDetail{}
	errMap := smapping.FillStruct(&groupDetailToUpdate, smapping.MapFields(&groupDetail))
	if errMap != nil {
		return groupDetailToUpdate, errMap
	}
	updatedGroupDetail, err := u.groupDetailRepository.UpdateGroupDetail(groupDetailToUpdate)
	return updatedGroupDetail, err
}
func (u *groupDetailService) DeleteGroupDetail(groupDetailID string, userHotelID string) (error, interface{}) {
	groupDetail, err := u.groupDetailRepository.FindGroupDetailByID(groupDetailID)
	if groupDetail.HotelID != userHotelID {
		return nil, false
	}
	if err != nil {
		return err, nil
	}
	errDel := u.groupDetailRepository.DeleteGroupDetail(groupDetail)
	return errDel, nil
}
func (u *groupDetailService) ShowGroupDetail(groupDetailID string, userHotelID string) (entity.GroupDetail, error, interface{}) {
	result, err := u.groupDetailRepository.FindGroupDetailByID(groupDetailID)
	if result.HotelID != userHotelID {
		return entity.GroupDetail{}, nil, false
	}
	return result, err, nil
}
func (u *groupDetailService) GetAllGroupDetail(hotelID string, filterPage dto.FilterPagination) ([]entity.GroupDetail, dto.Pagination, error) {
	admins, pagination, err := u.groupDetailRepository.GetAllGroupDetail(hotelID, filterPage)
	return admins, pagination, err
}
