package repository

import (
	"fmt"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"gorm.io/gorm"
	"math"
	"strings"
)

// CityRepository is contract what cityRepository can do to db
type CityRepository interface {
	GetAllCity(filterPagination dto.FilterPagination) ([]entity.City, dto.Pagination, error)
	InsertCity(city entity.City) (entity.City, error)
	UpdateCity(city entity.City) (entity.City, error)
	FindCityByID(cityID string) (entity.City, error)
	DeleteCity(city entity.City) error
}

// cityConnection adalah func untuk melakukan query data ke tabel city
type cityConnection struct {
	connection *gorm.DB
}

func (db *cityConnection) DeleteCity(city entity.City) error {
	err := db.connection.Where("id_city = ?", city.IDCity).Delete(&city).Error
	return err
}

func (db *cityConnection) GetAllCity(filterPagination dto.FilterPagination) ([]entity.City, dto.Pagination, error) {
	search := filterPagination.Search
	sortBy := filterPagination.SortBy
	orderBy := filterPagination.OrderBy
	perPage := int(filterPagination.PerPage)
	page := int(filterPagination.Page)

	var cities []entity.City
	query := db.connection.Model(&cities).Joins("LEFT JOIN provinces ON cities.province_id = provinces.id_province")

	whereClause := db.connection.Scopes(func(db *gorm.DB) *gorm.DB {
		if search != "" {
			keyword := strings.ToLower(search)
			if keyword != "" {
				db.Where("LOWER(cities.name) LIKE ?", fmt.Sprintf("%%%s%%", keyword)).
					Or("LOWER(provinces.name) LIKE ?", fmt.Sprintf("%%%s%%", keyword))
			}
		}
		return db
	})

	query.Where(whereClause).Scopes(func(db *gorm.DB) *gorm.DB {
		if filterPagination.ProvinceID != "" {
			db.Where("classes.province_id = ?", filterPagination.ProvinceID)
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
	//query.Find(&cities).Count(&total)

	err := query.Count(&total).Limit(perPage).Offset((page - 1) * perPage).Preload("Province").Find(&cities).Error

	totalPage := float64(total) / float64(perPage)

	pagination := dto.Pagination{
		Page:      uint(page),
		PerPage:   uint(perPage),
		TotalData: uint(total),
		TotalPage: uint(math.Ceil(totalPage)),
	}

	//err := db.connection.Find(&cities).Error
	return cities, pagination, err
}

// InsertCity is to add city in database
func (db *cityConnection) InsertCity(city entity.City) (entity.City, error) {
	err := db.connection.Save(&city).Error
	db.connection.Preload("Province").Find(&city)
	return city, err
}

// UpdateCity is func to edit city in database
func (db *cityConnection) UpdateCity(city entity.City) (entity.City, error) {
	err := db.connection.Where("id_city = ?", city.IDCity).Updates(&city).Error
	db.connection.Where("id_city = ?", city.IDCity).Preload("Province").Find(&city)
	return city, err
}

// FindCityByID is func to get city by email
func (db *cityConnection) FindCityByID(cityID string) (entity.City, error) {
	var city entity.City
	err := db.connection.Where("id_city = ?", cityID).Preload("Province").Take(&city).Error
	return city, err
}

// NewCityRepository is creates a new instance of CityRepository
func NewCityRepository(db *gorm.DB) CityRepository {
	return &cityConnection{
		connection: db,
	}
}
