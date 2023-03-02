package repository

import (
	"fmt"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"gorm.io/gorm"
	"math"
	"strings"
)

// ServiceDetailRepository is contract what serviceDetailRepository can do to db
type ServiceDetailRepository interface {
	GetAllServiceDetail(hotelID string, filterPagination dto.FilterPagination) ([]entity.ServiceDetail, dto.Pagination, error)
	InsertServiceDetail(serviceDetail entity.ServiceDetail) (entity.ServiceDetail, error)
	UpdateServiceDetail(serviceDetail entity.ServiceDetail) (entity.ServiceDetail, error)
	FindServiceDetailByID(serviceDetailID string) (entity.ServiceDetail, error)
	DeleteServiceDetail(serviceDetail entity.ServiceDetail) error
}

// serviceDetailConnection adalah func untuk melakukan query data ke tabel serviceDetail
type serviceDetailConnection struct {
	connection *gorm.DB
}

func (db *serviceDetailConnection) DeleteServiceDetail(serviceDetail entity.ServiceDetail) error {
	err := db.connection.Where("id_serviceDetail = ?", serviceDetail.IDServiceDetail).Delete(&serviceDetail).Error
	return err
}

func (db *serviceDetailConnection) GetAllServiceDetail(hotelID string, filterPagination dto.FilterPagination) ([]entity.ServiceDetail, dto.Pagination, error) {
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

	var serviceDetails []entity.ServiceDetail
	query := db.connection.Joins("JOIN services ON service_details.service_id = services.id_service").
		Joins("JOIN groups ON services.group_id = groups.id_group")

	whereClause := db.connection.Scopes(func(db *gorm.DB) *gorm.DB {
		if search != "" {
			keyword := strings.ToLower(search)
			if keyword != "" {
				query = query.Where("LOWER(services.name) LIKE ?", fmt.Sprintf("%%%s%%", keyword)).
					Or("LOWER(groups.name) LIKE ?", fmt.Sprintf("%%%s%%", keyword))
			}
		}
		return db
	})

	query.Where(whereClause).Scopes(func(db *gorm.DB) *gorm.DB {
		if filterPagination.ServiceID != "" {
			db.Where("service_details.service_id = ?", filterPagination.ServiceID)
		}
		if filterPagination.GroupID != "" {
			db.Where("service_details.group_id = ?", filterPagination.GroupID)
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

	err := query.Where("hotel_id = ?", hotelID).Limit(perPage).Offset((page - 1) * perPage).
		Preload("Service").
		Preload("Group").
		Preload("Hotel").
		Find(&serviceDetails).Count(&total).Error

	totalPage := float64(total) / float64(perPage)

	pagination := dto.Pagination{
		Page:      uint(page),
		PerPage:   uint(perPage),
		TotalData: uint(total),
		TotalPage: uint(math.Ceil(totalPage)),
	}

	return serviceDetails, pagination, err
}

// InsertServiceDetail is to add serviceDetail in database
func (db *serviceDetailConnection) InsertServiceDetail(serviceDetail entity.ServiceDetail) (entity.ServiceDetail, error) {
	err := db.connection.Save(&serviceDetail).Error
	db.connection.Find(&serviceDetail)
	return serviceDetail, err
}

// UpdateServiceDetail is func to edit serviceDetail in database
func (db *serviceDetailConnection) UpdateServiceDetail(serviceDetail entity.ServiceDetail) (entity.ServiceDetail, error) {
	err := db.connection.Where("id_serviceDetail = ?", serviceDetail.IDServiceDetail).Updates(&serviceDetail).Error
	db.connection.Where("id_serviceDetail = ?", serviceDetail.IDServiceDetail).Find(&serviceDetail)
	return serviceDetail, err
}

// FindServiceDetailByID is func to get serviceDetail by email
func (db *serviceDetailConnection) FindServiceDetailByID(serviceDetailID string) (entity.ServiceDetail, error) {
	var serviceDetail entity.ServiceDetail
	err := db.connection.Where("id_serviceDetail = ?", serviceDetailID).Take(&serviceDetail).Error
	return serviceDetail, err
}

// NewServiceDetailRepository is creates a new instance of ServiceDetailRepository
func NewServiceDetailRepository(db *gorm.DB) ServiceDetailRepository {
	return &serviceDetailConnection{
		connection: db,
	}
}
