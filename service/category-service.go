package service

import (
	"github.com/mashingan/smapping"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"github.com/yusufbagussh/pet_hotel_backend/repository"
)

// CategoryService is a contract of what categoryService can do
type CategoryService interface {
	GetAllCategory(filterPagination dto.FilterPagination) ([]entity.Category, dto.Pagination, error)
	CreateCategory(category dto.CreateCategory) (entity.Category, error)
	UpdateCategory(category dto.UpdateCategory) (entity.Category, error)
	DeleteCategory(categoryID string) error
	ShowCategory(categoryID string) (entity.Category, error)
}

type categoryService struct {
	categoryRepository repository.CategoryRepository
}

// NewCategoryService creates a new instance of CategoryService
func NewCategoryService(categoryRepo repository.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepository: categoryRepo,
	}
}

func (u *categoryService) CreateCategory(category dto.CreateCategory) (entity.Category, error) {
	categoryToCreate := entity.Category{}
	errMap := smapping.FillStruct(&categoryToCreate, smapping.MapFields(&category))
	if errMap != nil {
		return categoryToCreate, errMap
	}
	updatedCategory, err := u.categoryRepository.InsertCategory(categoryToCreate)
	return updatedCategory, err
}
func (u *categoryService) UpdateCategory(category dto.UpdateCategory) (entity.Category, error) {
	categoryToUpdate := entity.Category{}
	errMap := smapping.FillStruct(&categoryToUpdate, smapping.MapFields(&category))
	if errMap != nil {
		return categoryToUpdate, errMap
	}
	updatedCategory, err := u.categoryRepository.UpdateCategory(categoryToUpdate)
	return updatedCategory, err
}
func (u *categoryService) DeleteCategory(categoryID string) error {
	category, err := u.categoryRepository.FindCategoryByID(categoryID)
	if err != nil {
		return err
	}
	errDel := u.categoryRepository.DeleteCategory(category)
	return errDel
}
func (u *categoryService) ShowCategory(categoryID string) (entity.Category, error) {
	result, err := u.categoryRepository.FindCategoryByID(categoryID)
	return result, err
}
func (u *categoryService) GetAllCategory(filterPaginate dto.FilterPagination) ([]entity.Category, dto.Pagination, error) {
	admins, pagination, err := u.categoryRepository.GetAllCategory(filterPaginate)
	return admins, pagination, err
}
