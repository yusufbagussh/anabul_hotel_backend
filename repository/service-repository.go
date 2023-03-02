package repository

import (
	"fmt"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"gorm.io/gorm"
	"math"
	"strings"
)

// ServiceRepository is contract what serviceRepository can do to db
type ServiceRepository interface {
	GetAllService(hotelID string, filterPagination dto.FilterPagination) ([]entity.Service, dto.Pagination, error)
	InsertService(service entity.Service) (entity.Service, error)
	UpdateService(service entity.Service) (entity.Service, error)
	FindServiceByID(serviceID string) (entity.Service, error)
	DeleteService(service entity.Service) error
}

// serviceConnection adalah func untuk melakukan query data ke tabel service
type serviceConnection struct {
	connection *gorm.DB
}

func (db *serviceConnection) DeleteService(service entity.Service) error {
	err := db.connection.Where("id_service = ?", service.IDService).Delete(&service).Error
	return err
}

func (db *serviceConnection) GetAllService(hotelID string, filterPagination dto.FilterPagination) ([]entity.Service, dto.Pagination, error) {
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

	var services []entity.Service
	query := db.connection

	if search != "" {
		keyword := strings.ToLower(search)
		if keyword != "" {
			query = query.Where("LOWER(services.name) LIKE ?", fmt.Sprintf("%%%s%%", keyword))
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

	err := query.Where("hotel_id = ?", hotelID).Limit(perPage).Offset((page - 1) * perPage).Preload("Hotel").Find(&services).Count(&total).Error

	totalPage := float64(total) / float64(perPage)

	pagination := dto.Pagination{
		Page:      uint(page),
		PerPage:   uint(perPage),
		TotalData: uint(total),
		TotalPage: uint(math.Ceil(totalPage)),
	}

	return services, pagination, err
}

// InsertService is to add service in database
func (db *serviceConnection) InsertService(service entity.Service) (entity.Service, error) {
	err := db.connection.Save(&service).Error
	db.connection.Find(&service)
	return service, err
}

// UpdateService is func to edit service in database
func (db *serviceConnection) UpdateService(service entity.Service) (entity.Service, error) {
	err := db.connection.Where("id_service = ?", service.IDService).Updates(&service).Error
	db.connection.Where("id_service = ?", service.IDService).Find(&service)
	return service, err
}

// FindServiceByID is func to get service by email
func (db *serviceConnection) FindServiceByID(serviceID string) (entity.Service, error) {
	var service entity.Service
	err := db.connection.Where("id_service = ?", serviceID).Take(&service).Error
	return service, err
}

// NewServiceRepository is creates a new instance of ServiceRepository
func NewServiceRepository(db *gorm.DB) ServiceRepository {
	return &serviceConnection{
		connection: db,
	}
}
