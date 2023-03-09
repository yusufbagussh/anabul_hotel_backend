package repository

import (
	"fmt"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"gorm.io/gorm"
	"math"
	"strings"
)

// CageCategoryRepository is contract what cageCategoryRepository can do to db
type CageCategoryRepository interface {
	GetAllCageCategory(hotelID string, filterPagination dto.FilterPagination) ([]entity.CageCategory, dto.Pagination, error)
	InsertCageCategory(cageCategory entity.CageCategory) (entity.CageCategory, error)
	UpdateCageCategory(cageCategory entity.CageCategory) (entity.CageCategory, error)
	FindCageCategoryByID(cageCategoryID string) (entity.CageCategory, error)
	DeleteCageCategory(cageCategory entity.CageCategory) error
}

// cageCategoryConnection adalah func untuk melakukan query data ke tabel cageCategory
type cageCategoryConnection struct {
	connection *gorm.DB
}

func (db *cageCategoryConnection) DeleteCageCategory(cageCategory entity.CageCategory) error {
	err := db.connection.Where("id_cage_category = ?", cageCategory.IDCageCategory).Delete(&cageCategory).Error
	return err
}

func (db *cageCategoryConnection) GetAllCageCategory(hotelID string, filterPagination dto.FilterPagination) ([]entity.CageCategory, dto.Pagination, error) {
	search := filterPagination.Search
	sortBy := filterPagination.SortBy
	orderBy := filterPagination.OrderBy
	perPage := int(filterPagination.PerPage)
	page := int(filterPagination.Page)

	if page == 0 {
		page = 1
	}
	if perPage == 0 {
		perPage = 10
	}
	var total int64

	var cageCategories []entity.CageCategory
	query := db.connection.Model(&cageCategories)

	whereClause := db.connection.Scopes(func(db *gorm.DB) *gorm.DB {
		if search != "" {
			keyword := strings.ToLower(search)
			if keyword != "" {
				db.Where("LOWER(cage_categories.name) LIKE ?", fmt.Sprintf("%%%s%%", keyword))
			}
		}
		return db
	})

	query.Where(whereClause).Scopes(func(db *gorm.DB) *gorm.DB {
		if filterPagination.HotelID != "" {
			db.Where("cage_categories.hotel_id = ?", filterPagination.HotelID)
		} else {
			db.Where("cage_categories.hotel_id = ?", hotelID)
		}
		return db
	})

	listSortBy := []string{"name"}
	listSortOrder := []string{"desc", "asc"}

	if sortBy != "" && contains(listSortBy, sortBy) == true && orderBy != "" && contains(listSortOrder, orderBy) {
		query = query.Order(fmt.Sprintf("%s %s", sortBy, orderBy))
	} else {
		sortBy = "created_at"
		orderBy = "desc"
		query = query.Order(fmt.Sprintf("%s %s", sortBy, orderBy))
	}

	err := query.Count(&total).Limit(perPage).Offset((page - 1) * perPage).Preload("Hotel").Find(&cageCategories).Error

	totalPage := float64(total) / float64(perPage)

	pagination := dto.Pagination{
		Page:      uint(page),
		PerPage:   uint(perPage),
		TotalData: uint(total),
		TotalPage: uint(math.Ceil(totalPage)),
	}

	return cageCategories, pagination, err
}

// InsertCageCategory is to add cageCategory in database
func (db *cageCategoryConnection) InsertCageCategory(cageCategory entity.CageCategory) (entity.CageCategory, error) {
	err := db.connection.Save(&cageCategory).Error
	db.connection.Find(&cageCategory)
	return cageCategory, err
}

// UpdateCageCategory is func to edit cageCategory in database
func (db *cageCategoryConnection) UpdateCageCategory(cageCategory entity.CageCategory) (entity.CageCategory, error) {
	err := db.connection.Where("id_cage_category = ?", cageCategory.IDCageCategory).Updates(&cageCategory).Error
	db.connection.Where("id_cage_category = ?", cageCategory.IDCageCategory).Find(&cageCategory)
	return cageCategory, err
}

// FindCageCategoryByID is func to get cageCategory by email
func (db *cageCategoryConnection) FindCageCategoryByID(cageCategoryID string) (entity.CageCategory, error) {
	var cageCategory entity.CageCategory
	err := db.connection.Where("id_cage_category = ?", cageCategoryID).Take(&cageCategory).Error
	return cageCategory, err
}

// NewCageCategoryRepository is creates a new instance of CageCategoryRepository
func NewCageCategoryRepository(db *gorm.DB) CageCategoryRepository {
	return &cageCategoryConnection{
		connection: db,
	}
}
