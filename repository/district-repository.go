package repository

import (
	"fmt"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"gorm.io/gorm"
	"math"
	"strings"
)

// DistrictRepository is contract what districtRepository can do to db
type DistrictRepository interface {
	GetAllDistrict(filterPagination dto.FilterPagination) ([]entity.District, dto.Pagination, error)
	InsertDistrict(district entity.District) (entity.District, error)
	UpdateDistrict(district entity.District) (entity.District, error)
	FindDistrictByID(districtID string) (entity.District, error)
	DeleteDistrict(district entity.District) error
}

// districtConnection adalah func untuk melakukan query data ke tabel district
type districtConnection struct {
	connection *gorm.DB
}

func (db *districtConnection) DeleteDistrict(district entity.District) error {
	err := db.connection.Where("id_district = ?", district.IDDistrict).Delete(&district).Error
	return err
}

func (db *districtConnection) GetAllDistrict(filterPagination dto.FilterPagination) ([]entity.District, dto.Pagination, error) {
	search := filterPagination.Search
	sortBy := filterPagination.SortBy
	orderBy := filterPagination.OrderBy
	perPage := int(filterPagination.PerPage)
	page := int(filterPagination.Page)

	var categories []entity.District
	query := db.connection.Joins("JOIN classes ON categories.class_id = classes.id_class").
		Select("categories.id_district, categories.name, categories.class_id, categories.created_at, categories.updated_at, classes.id_class, classes.name as class_name")

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

// InsertDistrict is to add district in database
func (db *districtConnection) InsertDistrict(district entity.District) (entity.District, error) {
	err := db.connection.Save(&district).Error
	db.connection.Preload("Class").Find(&district)
	return district, err
}

// UpdateDistrict is func to edit district in database
func (db *districtConnection) UpdateDistrict(district entity.District) (entity.District, error) {
	err := db.connection.Where("id_district = ?", district.IDDistrict).Updates(&district).Error
	db.connection.Where("id_district = ?", district.IDDistrict).Preload("Class").Find(&district)
	return district, err
}

// FindDistrictByID is func to get district by email
func (db *districtConnection) FindDistrictByID(districtID string) (entity.District, error) {
	var district entity.District
	err := db.connection.Where("id_district = ?", districtID).Preload("Class").Take(&district).Error
	return district, err
}

// NewDistrictRepository is creates a new instance of DistrictRepository
func NewDistrictRepository(db *gorm.DB) DistrictRepository {
	return &districtConnection{
		connection: db,
	}
}
