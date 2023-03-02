package service

import (
	"github.com/mashingan/smapping"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"github.com/yusufbagussh/pet_hotel_backend/repository"
)

// GroupService is a contract of what groupService can do
type GroupService interface {
	GetAllGroup(hotelID string, filterPage dto.FilterPagination) ([]entity.Group, dto.Pagination, error)
	CreateGroup(group dto.CreateGroup) (entity.Group, error)
	UpdateGroup(group dto.UpdateGroup) (entity.Group, error)
	DeleteGroup(groupID string, userHotelID string) (error, interface{})
	ShowGroup(groupID string, userHotelID string) (entity.Group, error, interface{})
}

type groupService struct {
	groupRepository repository.GroupRepository
}

// NewGroupService creates a new instance of GroupService
func NewGroupService(groupRepo repository.GroupRepository) GroupService {
	return &groupService{
		groupRepository: groupRepo,
	}
}

func (u *groupService) CreateGroup(group dto.CreateGroup) (entity.Group, error) {
	groupToCreate := entity.Group{}
	errMap := smapping.FillStruct(&groupToCreate, smapping.MapFields(&group))
	if errMap != nil {
		return groupToCreate, errMap
	}
	updatedGroup, err := u.groupRepository.InsertGroup(groupToCreate)
	return updatedGroup, err
}
func (u *groupService) UpdateGroup(group dto.UpdateGroup) (entity.Group, error) {
	groupToUpdate := entity.Group{}
	errMap := smapping.FillStruct(&groupToUpdate, smapping.MapFields(&group))
	if errMap != nil {
		return groupToUpdate, errMap
	}
	updatedGroup, err := u.groupRepository.UpdateGroup(groupToUpdate)
	return updatedGroup, err
}
func (u *groupService) DeleteGroup(groupID string, userHotelID string) (error, interface{}) {
	group, err := u.groupRepository.FindGroupByID(groupID)
	if group.HotelID != userHotelID {
		return nil, false
	}
	if err != nil {
		return err, nil
	}
	errDel := u.groupRepository.DeleteGroup(group)
	return errDel, nil
}
func (u *groupService) ShowGroup(groupID string, userHotelID string) (entity.Group, error, interface{}) {
	result, err := u.groupRepository.FindGroupByID(groupID)
	if result.HotelID != userHotelID {
		return entity.Group{}, nil, false
	}
	return result, err, nil
}
func (u *groupService) GetAllGroup(hotelID string, filterPage dto.FilterPagination) ([]entity.Group, dto.Pagination, error) {
	admins, pagination, err := u.groupRepository.GetAllGroup(hotelID, filterPage)
	return admins, pagination, err
}
