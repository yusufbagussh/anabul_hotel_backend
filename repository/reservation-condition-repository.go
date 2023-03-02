package repository

import (
	"fmt"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"gorm.io/gorm"
	"math"
	"strings"
)

// ReservationConditionRepository is contract what reservationConditionRepository can do to db
type ReservationConditionRepository interface {
	GetAllReservationCondition(hotelID string, filterPagination dto.FilterPagination) ([]entity.ReservationCondition, dto.Pagination, error)
	InsertReservationCondition(reservationCondition entity.ReservationCondition) (entity.ReservationCondition, error)
	UpdateReservationCondition(reservationCondition entity.ReservationCondition) (entity.ReservationCondition, error)
	FindReservationConditionByID(reservationConditionID string) (entity.ReservationCondition, error)
	DeleteReservationCondition(reservationCondition entity.ReservationCondition) error
}

// reservationConditionConnection adalah func untuk melakukan query data ke tabel reservationCondition
type reservationConditionConnection struct {
	connection *gorm.DB
}

func (db *reservationConditionConnection) DeleteReservationCondition(reservationCondition entity.ReservationCondition) error {
	err := db.connection.Where("id_reservationCondition = ?", reservationCondition.IDReservationCondition).Delete(&reservationCondition).Error
	return err
}

func (db *reservationConditionConnection) GetAllReservationCondition(hotelID string, filterPagination dto.FilterPagination) ([]entity.ReservationCondition, dto.Pagination, error) {
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

	var reservationConditions []entity.ReservationCondition
	query := db.connection

	if search != "" {
		keyword := strings.ToLower(search)
		if keyword != "" {
			query = query.Where("LOWER(reservationConditions.name) LIKE ?", fmt.Sprintf("%%%s%%", keyword))
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

	err := query.Where("hotel_id = ?", hotelID).Limit(perPage).Offset((page - 1) * perPage).Preload("Hotel").Find(&reservationConditions).Count(&total).Error

	totalPage := float64(total) / float64(perPage)

	pagination := dto.Pagination{
		Page:      uint(page),
		PerPage:   uint(perPage),
		TotalData: uint(total),
		TotalPage: uint(math.Ceil(totalPage)),
	}

	return reservationConditions, pagination, err
}

// InsertReservationCondition is to add reservationCondition in database
func (db *reservationConditionConnection) InsertReservationCondition(reservationCondition entity.ReservationCondition) (entity.ReservationCondition, error) {
	err := db.connection.Save(&reservationCondition).Error
	db.connection.Find(&reservationCondition)
	return reservationCondition, err
}

// UpdateReservationCondition is func to edit reservationCondition in database
func (db *reservationConditionConnection) UpdateReservationCondition(reservationCondition entity.ReservationCondition) (entity.ReservationCondition, error) {
	err := db.connection.Where("id_reservationCondition = ?", reservationCondition.IDReservationCondition).Updates(&reservationCondition).Error
	db.connection.Where("id_reservationCondition = ?", reservationCondition.IDReservationCondition).Find(&reservationCondition)
	return reservationCondition, err
}

// FindReservationConditionByID is func to get reservationCondition by email
func (db *reservationConditionConnection) FindReservationConditionByID(reservationConditionID string) (entity.ReservationCondition, error) {
	var reservationCondition entity.ReservationCondition
	err := db.connection.Where("id_reservationCondition = ?", reservationConditionID).Take(&reservationCondition).Error
	return reservationCondition, err
}

// NewReservationConditionRepository is creates a new instance of ReservationConditionRepository
func NewReservationConditionRepository(db *gorm.DB) ReservationConditionRepository {
	return &reservationConditionConnection{
		connection: db,
	}
}
