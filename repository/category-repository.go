package repository

import (
	"fmt"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"gorm.io/gorm"
	"math"
	"strings"
)

// CategoryRepository is contract what categoryRepository can do to db
type CategoryRepository interface {
	GetAllCategory(filterPagination dto.FilterPagination) ([]entity.Category, dto.Pagination, error)
	InsertCategory(category entity.Category) (entity.Category, error)
	UpdateCategory(category entity.Category) (entity.Category, error)
	FindCategoryByID(categoryID string) (entity.Category, error)
	DeleteCategory(category entity.Category) error
}

// categoryConnection adalah func untuk melakukan query data ke tabel category
type categoryConnection struct {
	connection *gorm.DB
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func (db *categoryConnection) DeleteCategory(category entity.Category) error {
	err := db.connection.Where("id_category = ?", category.IDCategory).Delete(&category).Error
	return err
}

func (db *categoryConnection) GetAllCategory(filterPagination dto.FilterPagination) ([]entity.Category, dto.Pagination, error) {
	search := filterPagination.Search
	sortBy := filterPagination.SortBy
	orderBy := filterPagination.OrderBy
	perPage := int(filterPagination.PerPage)
	page := int(filterPagination.Page)

	var categories []entity.Category
	query := db.connection.Joins("JOIN classes ON categories.class_id = classes.id_class").
		Select("categories.id_category, categories.name, categories.class_id, categories.created_at, categories.updated_at, classes.id_class, classes.name as class_name")

	whereClause := db.connection.Scopes(func(db *gorm.DB) *gorm.DB {
		if search != "" {
			keyword := strings.ToLower(search)
			if keyword != "" {
				db.Where("LOWER(categories.name) LIKE ?", fmt.Sprintf("%%%s%%", keyword)).
					Or("LOWER(classes.name) LIKE ?", fmt.Sprintf("%%%s%%", keyword))
			}
		}
		return db
	})

	query.Where(whereClause).Scopes(func(db *gorm.DB) *gorm.DB {
		if filterPagination.ClassID != "" {
			db.Where("categories.class_id = ?", filterPagination.ClassID)
		}
		return db
	})

	listSortBy := []string{"name", "class_name"}
	listSortOrder := []string{"desc", "asc"}

	if sortBy != "" && contains(listSortBy, sortBy) == true && orderBy != "" && contains(listSortOrder, orderBy) {
		query = query.Order(fmt.Sprintf("%s %s", sortBy, orderBy))
	} else {
		sortBy = "created_at"
		orderBy = "desc"
		query = query.Order(fmt.Sprintf("%s %s", sortBy, orderBy))
	}

	var total int64
	if page == 0 {
		page = 1
	}
	if perPage == 0 {
		perPage = 10
	}
	query.Find(&categories).Count(&total)

	err := query.Limit(perPage).Offset((page - 1) * perPage).Preload("Class").Find(&categories).Error

	totalPage := float64(total) / float64(perPage)

	pagination := dto.Pagination{
		Page:      uint(page),
		PerPage:   uint(perPage),
		TotalData: uint(total),
		TotalPage: uint(math.Ceil(totalPage)),
	}

	//err := db.connection.Find(&categories).Error
	return categories, pagination, err
}

// InsertCategory is to add category in database
func (db *categoryConnection) InsertCategory(category entity.Category) (entity.Category, error) {
	err := db.connection.Save(&category).Error
	db.connection.Preload("Class").Find(&category)
	return category, err
}

// UpdateCategory is func to edit category in database
func (db *categoryConnection) UpdateCategory(category entity.Category) (entity.Category, error) {
	err := db.connection.Where("id_category = ?", category.IDCategory).Updates(&category).Error
	db.connection.Where("id_category = ?", category.IDCategory).Preload("Class").Find(&category)
	return category, err
}

// FindCategoryByID is func to get category by email
func (db *categoryConnection) FindCategoryByID(categoryID string) (entity.Category, error) {
	var category entity.Category
	err := db.connection.Where("id_category = ?", categoryID).Preload("Class").Take(&category).Error
	return category, err
}

// NewCategoryRepository is creates a new instance of CategoryRepository
func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryConnection{
		connection: db,
	}
}
