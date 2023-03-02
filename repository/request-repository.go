package repository

import (
	"fmt"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"gorm.io/gorm"
	"math"
)

// RequestRepository is contract what requestRepository can do to db
type RequestRepository interface {
	GetAllRequest(filterPagination dto.FilterPagination) ([]entity.Request, dto.Pagination, error)
	InsertRequest(request entity.Request) (entity.Request, error)
	UpdateRequest(request entity.Request) (entity.Request, error)
	FindRequestByID(requestID string) (entity.Request, error)
	DeleteRequest(request entity.Request) error
}

// requestConnection adalah func untuk melakukan query data ke tabel request
type requestConnection struct {
	connection *gorm.DB
}

func (db *requestConnection) DeleteRequest(request entity.Request) error {
	err := db.connection.Where("id_request = ?", request.IDRequest).Delete(&request).Error
	return err
}

func (db *requestConnection) GetAllRequest(filterPagination dto.FilterPagination) ([]entity.Request, dto.Pagination, error) {
	var requestes []entity.Request
	search := filterPagination.Search
	sortBy := filterPagination.SortBy
	orderBy := filterPagination.OrderBy
	perPage := filterPagination.PerPage
	page := filterPagination.Page
	if sortBy == "" {
		sortBy = "created_at"
	}
	if orderBy == "" {
		orderBy = "desc"
	}
	if page == 0 {
		page = 1
	}
	if perPage == 0 {
		perPage = 10
	}

	query := "SELECT * FROM requestes"
	if search != "" {
		query = fmt.Sprintf("%s WHERE name LIKE '%%%s%%'", query, search)
	}
	query = fmt.Sprintf("%s ORDER BY %s %s", query, sortBy, orderBy)

	var total int64

	db.connection.Raw(query).Count(&total)
	query = fmt.Sprintf("%s LIMIT %d OFFSET %d", query, perPage, (page-1)*perPage)
	err := db.connection.Raw(query).Scan(&requestes).Error

	totalPage := float64(total) / float64(perPage)

	pagination := dto.Pagination{
		Page:      uint(page),
		PerPage:   uint(perPage),
		TotalData: uint(total),
		TotalPage: uint(math.Ceil(totalPage)),
	}

	return requestes, pagination, err
}

// InsertRequest is to add request in database
func (db *requestConnection) InsertRequest(request entity.Request) (entity.Request, error) {
	err := db.connection.Save(&request).Error
	return request, err
}

// UpdateRequest is func to edit request in database
func (db *requestConnection) UpdateRequest(request entity.Request) (entity.Request, error) {
	err := db.connection.Where("id_request = ?", request.IDRequest).Updates(&request).Error
	db.connection.Where("id_request = ?", request.IDRequest).Find(&request)
	return request, err
}

// FindRequestByID is func to get request by email
func (db *requestConnection) FindRequestByID(requestID string) (entity.Request, error) {
	var request entity.Request
	err := db.connection.Where("id_request = ?", requestID).Take(&request).Error
	return request, err
}

// NewRequestRepository is creates a new instance of RequestRepository
func NewRequestRepository(db *gorm.DB) RequestRepository {
	return &requestConnection{
		connection: db,
	}
}
