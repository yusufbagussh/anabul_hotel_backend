package repository

import (
	"fmt"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"gorm.io/gorm"
	"math"
	"strings"
)

// ReservationInventoryRepository is contract what reservationInventoryRepository can do to db
type ReservationInventoryRepository interface {
	GetAllReservationInventory(filterPagination dto.FilterPagination) ([]entity.ReservationInventory, dto.Pagination, error)
	InsertReservationInventory(reservationInventory entity.ReservationInventory) (entity.ReservationInventory, error)
	UpdateReservationInventory(reservationInventory entity.ReservationInventory) (entity.ReservationInventory, error)
	FindReservationInventoryByID(reservationInventoryID string) (entity.ReservationInventory, error)
	DeleteReservationInventory(reservationInventory entity.ReservationInventory) error
}

// reservationInventoryConnection adalah func untuk melakukan query data ke tabel reservationInventory
type reservationInventoryConnection struct {
	connection *gorm.DB
}

func (db *reservationInventoryConnection) DeleteReservationInventory(reservationInventory entity.ReservationInventory) error {
	err := db.connection.Where("id_reservationInventory = ?", reservationInventory.IDReservationInventory).Delete(&reservationInventory).Error
	return err
}

func (db *reservationInventoryConnection) GetAllReservationInventory(filterPagination dto.FilterPagination) ([]entity.ReservationInventory, dto.Pagination, error) {
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

	var reservationInventorys []entity.ReservationInventory
	query := db.connection

	if search != "" {
		keyword := strings.ToLower(search)
		if keyword != "" {
			query = query.Where("LOWER(reservationInventorys.name) LIKE ?", fmt.Sprintf("%%%s%%", keyword))
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

	err := query.Limit(perPage).Offset((page - 1) * perPage).Preload("Hotel").Find(&reservationInventorys).Count(&total).Error

	totalPage := float64(total) / float64(perPage)

	pagination := dto.Pagination{
		Page:      uint(page),
		PerPage:   uint(perPage),
		TotalData: uint(total),
		TotalPage: uint(math.Ceil(totalPage)),
	}

	return reservationInventorys, pagination, err
}

// InsertReservationInventory is to add reservationInventory in database
func (db *reservationInventoryConnection) InsertReservationInventory(reservationInventory entity.ReservationInventory) (entity.ReservationInventory, error) {
	err := db.connection.Save(&reservationInventory).Error
	db.connection.Find(&reservationInventory)
	return reservationInventory, err
}

// UpdateReservationInventory is func to edit reservationInventory in database
func (db *reservationInventoryConnection) UpdateReservationInventory(reservationInventory entity.ReservationInventory) (entity.ReservationInventory, error) {
	err := db.connection.Where("id_reservation_inventory = ?", reservationInventory.IDReservationInventory).Updates(&reservationInventory).Error
	db.connection.Where("id_reservation_inventory = ?", reservationInventory.IDReservationInventory).Find(&reservationInventory)
	return reservationInventory, err
}

// FindReservationInventoryByID is func to get reservationInventory by email
func (db *reservationInventoryConnection) FindReservationInventoryByID(reservationInventoryID string) (entity.ReservationInventory, error) {
	var reservationInventory entity.ReservationInventory
	err := db.connection.Where("id_reservationInventory = ?", reservationInventoryID).Take(&reservationInventory).Error
	return reservationInventory, err
}

// NewReservationInventoryRepository is creates a new instance of ReservationInventoryRepository
func NewReservationInventoryRepository(db *gorm.DB) ReservationInventoryRepository {
	return &reservationInventoryConnection{
		connection: db,
	}
}
