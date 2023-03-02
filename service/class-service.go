package service

import (
	"github.com/mashingan/smapping"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"github.com/yusufbagussh/pet_hotel_backend/repository"
)

// ClassService is a contract of what classService can do
type ClassService interface {
	GetAllClass(filterPage dto.FilterPagination) ([]entity.Class, dto.Pagination, error)
	CreateClass(class dto.CreateClass) (entity.Class, error)
	UpdateClass(class dto.UpdateClass) (entity.Class, error)
	DeleteClass(classID string) error
	ShowClass(classID string) (entity.Class, error)
}

type classService struct {
	classRepository repository.ClassRepository
}

// NewClassService creates a new instance of ClassService
func NewClassService(classRepo repository.ClassRepository) ClassService {
	return &classService{
		classRepository: classRepo,
	}
}

func (u *classService) CreateClass(class dto.CreateClass) (entity.Class, error) {
	classToCreate := entity.Class{}
	errMap := smapping.FillStruct(&classToCreate, smapping.MapFields(&class))
	if errMap != nil {
		return classToCreate, errMap
	}
	updatedClass, err := u.classRepository.InsertClass(classToCreate)
	return updatedClass, err
}
func (u *classService) UpdateClass(class dto.UpdateClass) (entity.Class, error) {
	classToUpdate := entity.Class{}
	errMap := smapping.FillStruct(&classToUpdate, smapping.MapFields(&class))
	if errMap != nil {
		return classToUpdate, errMap
	}
	updatedClass, err := u.classRepository.UpdateClass(classToUpdate)
	return updatedClass, err
}
func (u *classService) DeleteClass(classID string) error {
	class, err := u.classRepository.FindClassByID(classID)
	if err != nil {
		return err
	}
	errDel := u.classRepository.DeleteClass(class)
	return errDel
}
func (u *classService) ShowClass(classID string) (entity.Class, error) {
	result, err := u.classRepository.FindClassByID(classID)
	return result, err
}
func (u *classService) GetAllClass(filterPage dto.FilterPagination) ([]entity.Class, dto.Pagination, error) {
	admins, pagination, err := u.classRepository.GetAllClass(filterPage)
	return admins, pagination, err
}
