package repository

import (
	"fmt"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"gorm.io/gorm"
	"math"
	"strings"
)

// CageTypeRepository is contract what cageTypeRepository can do to db
type CageTypeRepository interface {
	GetAllCageType(hotelID string, filterPagination dto.FilterPagination) ([]entity.CageType, dto.Pagination, error)
	InsertCageType(cageType entity.CageType) (entity.CageType, error)
	UpdateCageType(cageType entity.CageType) (entity.CageType, error)
	FindCageTypeByID(cageTypeID string) (entity.CageType, error)
	DeleteCageType(cageType entity.CageType) error
}

// cageTypeConnection adalah func untuk melakukan query data ke tabel cageType
type cageTypeConnection struct {
	connection *gorm.DB
}

func (db *cageTypeConnection) DeleteCageType(cageType entity.CageType) error {
	err := db.connection.Where("id_cage_type = ?", cageType.IDCageType).Delete(&cageType).Error
	return err
}

func (db *cageTypeConnection) GetAllCageType(hotelID string, filterPagination dto.FilterPagination) ([]entity.CageType, dto.Pagination, error) {
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

	var cageTypes []entity.CageType
	query := db.connection.Model(&cageTypes)

	whereClause := db.connection.Scopes(func(db *gorm.DB) *gorm.DB {
		if search != "" {
			keyword := strings.ToLower(search)
			if keyword != "" {
				db.Where("LOWER(cageTypes.name) LIKE ?", fmt.Sprintf("%%%s%%", keyword))
			}
		}
		return db
	})

	query.Where(whereClause).Scopes(func(db *gorm.DB) *gorm.DB {
		if filterPagination.HotelID != "" {
			db.Where("cage_types.hotel_id = ?", filterPagination.HotelID)
		} else {
			db.Where("cage_types.hotel_id = ?", hotelID)
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

	err := query.Count(&total).Limit(perPage).Offset((page - 1) * perPage).Preload("Hotel").Find(&cageTypes).Error

	totalPage := float64(total) / float64(perPage)

	pagination := dto.Pagination{
		Page:      uint(page),
		PerPage:   uint(perPage),
		TotalData: uint(total),
		TotalPage: uint(math.Ceil(totalPage)),
	}

	return cageTypes, pagination, err
}

// InsertCageType is to add cageType in database
func (db *cageTypeConnection) InsertCageType(cageType entity.CageType) (entity.CageType, error) {
	err := db.connection.Save(&cageType).Error
	db.connection.Find(&cageType)
	return cageType, err
}

// UpdateCageType is func to edit cageType in database
func (db *cageTypeConnection) UpdateCageType(cageType entity.CageType) (entity.CageType, error) {
	err := db.connection.Where("id_cage_type = ?", cageType.IDCageType).Updates(&cageType).Error
	db.connection.Where("id_cage_type = ?", cageType.IDCageType).Find(&cageType)
	return cageType, err
}

// FindCageTypeByID is func to get cageType by email
func (db *cageTypeConnection) FindCageTypeByID(cageTypeID string) (entity.CageType, error) {
	var cageType entity.CageType
	err := db.connection.Where("id_cage_type = ?", cageTypeID).Take(&cageType).Error
	return cageType, err
}

// NewCageTypeRepository is creates a new instance of CageTypeRepository
func NewCageTypeRepository(db *gorm.DB) CageTypeRepository {
	return &cageTypeConnection{
		connection: db,
	}
}
