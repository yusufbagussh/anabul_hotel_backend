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

	var districts []entity.District
	query := db.connection.Model(&districts).Joins("LEFT JOIN cities ON districts.city_id = cities.id_city")

	whereClause := db.connection.Scopes(func(db *gorm.DB) *gorm.DB {
		if search != "" {
			keyword := strings.ToLower(search)
			if keyword != "" {
				db.Where("LOWER(districts.name) LIKE ?", fmt.Sprintf("%%%s%%", keyword)).
					Or("LOWER(cities.name) LIKE ?", fmt.Sprintf("%%%s%%", keyword))
			}
		}
		return db
	})

	query.Where(whereClause).Scopes(func(db *gorm.DB) *gorm.DB {
		if filterPagination.CityID != "" {
			db.Where("districts.city_id = ?", filterPagination.CityID)
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

	var total int64
	if page == 0 {
		page = 1
	}
	if perPage == 0 {
		perPage = 10
	}
	//query.Find(&districts).Count(&total)

	err := query.Count(&total).Limit(perPage).Offset((page - 1) * perPage).Preload("City").Find(&districts).Error

	totalPage := float64(total) / float64(perPage)

	pagination := dto.Pagination{
		Page:      uint(page),
		PerPage:   uint(perPage),
		TotalData: uint(total),
		TotalPage: uint(math.Ceil(totalPage)),
	}

	//err := db.connection.Find(&districts).Error
	return districts, pagination, err
}

// InsertDistrict is to add district in database
func (db *districtConnection) InsertDistrict(district entity.District) (entity.District, error) {
	err := db.connection.Save(&district).Error
	db.connection.Preload("City").Find(&district)
	return district, err
}

// UpdateDistrict is func to edit district in database
func (db *districtConnection) UpdateDistrict(district entity.District) (entity.District, error) {
	err := db.connection.Where("id_district = ?", district.IDDistrict).Updates(&district).Error
	db.connection.Where("id_district = ?", district.IDDistrict).Preload("City").Find(&district)
	return district, err
}

// FindDistrictByID is func to get district by email
func (db *districtConnection) FindDistrictByID(districtID string) (entity.District, error) {
	var district entity.District
	err := db.connection.Where("id_district = ?", districtID).Preload("City").Take(&district).Error
	return district, err
}

// NewDistrictRepository is creates a new instance of DistrictRepository
func NewDistrictRepository(db *gorm.DB) DistrictRepository {
	return &districtConnection{
		connection: db,
	}
}
