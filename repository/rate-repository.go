package repository

import (
	"fmt"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"gorm.io/gorm"
	"math"
	"strings"
)

// RateRepository is contract what rateRepository can do to db
type RateRepository interface {
	GetAllRate(hotelID string, filterPagination dto.FilterPagination) ([]entity.Rate, dto.Pagination, error)
	InsertRate(rate entity.Rate) (entity.Rate, error)
	UpdateRate(rate entity.Rate) (entity.Rate, error)
	FindRateByID(rateID string) (entity.Rate, error)
	DeleteRate(rate entity.Rate) error
}

// rateConnection adalah func untuk melakukan query data ke tabel rate
type rateConnection struct {
	connection *gorm.DB
}

func (db *rateConnection) DeleteRate(rate entity.Rate) error {
	err := db.connection.Where("id_rate = ?", rate.IDRate).Delete(&rate).Error
	return err
}

func (db *rateConnection) GetAllRate(hotelID string, filterPagination dto.FilterPagination) ([]entity.Rate, dto.Pagination, error) {
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

	var rates []entity.Rate
	query := db.connection

	if search != "" {
		keyword := strings.ToLower(search)
		if keyword != "" {
			query = query.Where("LOWER(rates.name) LIKE ?", fmt.Sprintf("%%%s%%", keyword))
		}
	}

	listSortBy := []string{"name"}
	listSortOrder := []string{"desc", "asc"}

	if sortBy != "" && contains(listSortBy, sortBy) == true && orderBy != "" && contains(listSortOrder, orderBy) {
		query = query.Order(fmt.Sprintf("%s %s", sortBy, orderBy))
	} else {
		sortBy = "created_at"
		orderBy = "desc"
		query = query.Order(fmt.Sprintf("%s %s", sortBy, orderBy))
	}

	err := query.Where("hotel_id = ?", hotelID).Limit(perPage).Offset((page - 1) * perPage).Preload("Hotel").Find(&rates).Count(&total).Error

	totalPage := float64(total) / float64(perPage)

	pagination := dto.Pagination{
		Page:      uint(page),
		PerPage:   uint(perPage),
		TotalData: uint(total),
		TotalPage: uint(math.Ceil(totalPage)),
	}

	return rates, pagination, err
}

// InsertRate is to add rate in database
func (db *rateConnection) InsertRate(rate entity.Rate) (entity.Rate, error) {
	err := db.connection.Save(&rate).Error
	db.connection.Find(&rate)
	return rate, err
}

// UpdateRate is func to edit rate in database
func (db *rateConnection) UpdateRate(rate entity.Rate) (entity.Rate, error) {
	err := db.connection.Where("id_rate = ?", rate.IDRate).Updates(&rate).Error
	db.connection.Where("id_rate = ?", rate.IDRate).Find(&rate)
	return rate, err
}

// FindRateByID is func to get rate by email
func (db *rateConnection) FindRateByID(rateID string) (entity.Rate, error) {
	var rate entity.Rate
	err := db.connection.Where("id_rate = ?", rateID).Take(&rate).Error
	return rate, err
}

// NewRateRepository is creates a new instance of RateRepository
func NewRateRepository(db *gorm.DB) RateRepository {
	return &rateConnection{
		connection: db,
	}
}
