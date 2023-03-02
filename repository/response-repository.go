package repository

import (
	"fmt"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"gorm.io/gorm"
	"math"
	"strings"
)

// ResponseRepository is contract what responseRepository can do to db
type ResponseRepository interface {
	GetAllResponse(hotelID string, filterPagination dto.FilterPagination) ([]entity.Response, dto.Pagination, error)
	InsertResponse(response entity.Response) (entity.Response, error)
	UpdateResponse(response entity.Response) (entity.Response, error)
	FindResponseByID(responseID string) (entity.Response, error)
	DeleteResponse(response entity.Response) error
}

// responseConnection adalah func untuk melakukan query data ke tabel response
type responseConnection struct {
	connection *gorm.DB
}

func (db *responseConnection) DeleteResponse(response entity.Response) error {
	err := db.connection.Where("id_response = ?", response.IDResponse).Delete(&response).Error
	return err
}

func (db *responseConnection) GetAllResponse(hotelID string, filterPagination dto.FilterPagination) ([]entity.Response, dto.Pagination, error) {
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

	var responses []entity.Response
	query := db.connection

	if search != "" {
		keyword := strings.ToLower(search)
		if keyword != "" {
			query = query.Where("LOWER(responses.name) LIKE ?", fmt.Sprintf("%%%s%%", keyword))
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

	err := query.Where("hotel_id = ?", hotelID).Limit(perPage).Offset((page - 1) * perPage).Preload("Hotel").Find(&responses).Count(&total).Error

	totalPage := float64(total) / float64(perPage)

	pagination := dto.Pagination{
		Page:      uint(page),
		PerPage:   uint(perPage),
		TotalData: uint(total),
		TotalPage: uint(math.Ceil(totalPage)),
	}

	return responses, pagination, err
}

// InsertResponse is to add response in database
func (db *responseConnection) InsertResponse(response entity.Response) (entity.Response, error) {
	err := db.connection.Save(&response).Error
	db.connection.Find(&response)
	return response, err
}

// UpdateResponse is func to edit response in database
func (db *responseConnection) UpdateResponse(response entity.Response) (entity.Response, error) {
	err := db.connection.Where("id_response = ?", response.IDResponse).Updates(&response).Error
	db.connection.Where("id_response = ?", response.IDResponse).Find(&response)
	return response, err
}

// FindResponseByID is func to get response by email
func (db *responseConnection) FindResponseByID(responseID string) (entity.Response, error) {
	var response entity.Response
	err := db.connection.Where("id_response = ?", responseID).Take(&response).Error
	return response, err
}

// NewResponseRepository is creates a new instance of ResponseRepository
func NewResponseRepository(db *gorm.DB) ResponseRepository {
	return &responseConnection{
		connection: db,
	}
}
