package repository

import (
	"fmt"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"gorm.io/gorm"
	"math"
	"strings"
)

// ReservationServiceRepository is contract what reservationServiceRepository can do to db
type ReservationServiceRepository interface {
	GetAllReservationService(filterPagination dto.FilterPagination) ([]entity.ReservationService, dto.Pagination, error)
	InsertReservationService(reservationService entity.ReservationService) (entity.ReservationService, error)
	UpdateReservationService(reservationService entity.ReservationService) (entity.ReservationService, error)
	FindReservationServiceByID(reservationServiceID string) (entity.ReservationService, error)
	DeleteReservationService(reservationService entity.ReservationService) error
}

// reservationServiceConnection adalah func untuk melakukan query data ke tabel reservationService
type reservationServiceConnection struct {
	connection *gorm.DB
}

func (db *reservationServiceConnection) DeleteReservationService(reservationService entity.ReservationService) error {
	err := db.connection.Where("id_reservationService = ?", reservationService.IDReservationService).Delete(&reservationService).Error
	return err
}

func (db *reservationServiceConnection) GetAllReservationService(filterPagination dto.FilterPagination) ([]entity.ReservationService, dto.Pagination, error) {
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

	var reservationServices []entity.ReservationService
	query := db.connection

	if search != "" {
		keyword := strings.ToLower(search)
		if keyword != "" {
			query = query.Where("LOWER(reservationServices.name) LIKE ?", fmt.Sprintf("%%%s%%", keyword))
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

	err := query.Limit(perPage).Offset((page - 1) * perPage).Preload("Hotel").Find(&reservationServices).Count(&total).Error

	totalPage := float64(total) / float64(perPage)

	pagination := dto.Pagination{
		Page:      uint(page),
		PerPage:   uint(perPage),
		TotalData: uint(total),
		TotalPage: uint(math.Ceil(totalPage)),
	}

	return reservationServices, pagination, err
}

// InsertReservationService is to add reservationService in database
func (db *reservationServiceConnection) InsertReservationService(reservationService entity.ReservationService) (entity.ReservationService, error) {
	err := db.connection.Save(&reservationService).Error
	db.connection.Find(&reservationService)
	return reservationService, err
}

// UpdateReservationService is func to edit reservationService in database
func (db *reservationServiceConnection) UpdateReservationService(reservationService entity.ReservationService) (entity.ReservationService, error) {
	err := db.connection.Where("id_reservationService = ?", reservationService.IDReservationService).Updates(&reservationService).Error
	db.connection.Where("id_reservationService = ?", reservationService.IDReservationService).Find(&reservationService)
	return reservationService, err
}

// FindReservationServiceByID is func to get reservationService by email
func (db *reservationServiceConnection) FindReservationServiceByID(reservationServiceID string) (entity.ReservationService, error) {
	var reservationService entity.ReservationService
	err := db.connection.Where("id_reservationService = ?", reservationServiceID).Take(&reservationService).Error
	return reservationService, err
}

// NewReservationServiceRepository is creates a new instance of ReservationServiceRepository
func NewReservationServiceRepository(db *gorm.DB) ReservationServiceRepository {
	return &reservationServiceConnection{
		connection: db,
	}
}
