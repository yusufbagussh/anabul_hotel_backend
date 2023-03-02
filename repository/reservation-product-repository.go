package repository

import (
	"fmt"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"gorm.io/gorm"
	"math"
	"strings"
)

// ReservationProductRepository is contract what reservationProductRepository can do to db
type ReservationProductRepository interface {
	GetAllReservationProduct(filterPagination dto.FilterPagination) ([]entity.ReservationProduct, dto.Pagination, error)
	InsertReservationProduct(reservationProduct entity.ReservationProduct) (entity.ReservationProduct, error)
	UpdateReservationProduct(reservationProduct entity.ReservationProduct) (entity.ReservationProduct, error)
	FindReservationProductByID(reservationProductID string) (entity.ReservationProduct, error)
	DeleteReservationProduct(reservationProduct entity.ReservationProduct) error
}

// reservationProductConnection adalah func untuk melakukan query data ke tabel reservationProduct
type reservationProductConnection struct {
	connection *gorm.DB
}

func (db *reservationProductConnection) DeleteReservationProduct(reservationProduct entity.ReservationProduct) error {
	err := db.connection.Where("id_reservationProduct = ?", reservationProduct.IDReservationProduct).Delete(&reservationProduct).Error
	return err
}

func (db *reservationProductConnection) GetAllReservationProduct(filterPagination dto.FilterPagination) ([]entity.ReservationProduct, dto.Pagination, error) {
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

	var reservationProducts []entity.ReservationProduct
	query := db.connection

	if search != "" {
		keyword := strings.ToLower(search)
		if keyword != "" {
			query = query.Where("LOWER(reservationProducts.name) LIKE ?", fmt.Sprintf("%%%s%%", keyword))
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

	err := query.Limit(perPage).Offset((page - 1) * perPage).Preload("Hotel").Find(&reservationProducts).Count(&total).Error

	totalPage := float64(total) / float64(perPage)

	pagination := dto.Pagination{
		Page:      uint(page),
		PerPage:   uint(perPage),
		TotalData: uint(total),
		TotalPage: uint(math.Ceil(totalPage)),
	}

	return reservationProducts, pagination, err
}

// InsertReservationProduct is to add reservationProduct in database
func (db *reservationProductConnection) InsertReservationProduct(reservationProduct entity.ReservationProduct) (entity.ReservationProduct, error) {
	err := db.connection.Save(&reservationProduct).Error
	db.connection.Find(&reservationProduct)
	return reservationProduct, err
}

// UpdateReservationProduct is func to edit reservationProduct in database
func (db *reservationProductConnection) UpdateReservationProduct(reservationProduct entity.ReservationProduct) (entity.ReservationProduct, error) {
	err := db.connection.Where("id_reservationProduct = ?", reservationProduct.IDReservationProduct).Updates(&reservationProduct).Error
	db.connection.Where("id_reservationProduct = ?", reservationProduct.IDReservationProduct).Find(&reservationProduct)
	return reservationProduct, err
}

// FindReservationProductByID is func to get reservationProduct by email
func (db *reservationProductConnection) FindReservationProductByID(reservationProductID string) (entity.ReservationProduct, error) {
	var reservationProduct entity.ReservationProduct
	err := db.connection.Where("id_reservationProduct = ?", reservationProductID).Take(&reservationProduct).Error
	return reservationProduct, err
}

// NewReservationProductRepository is creates a new instance of ReservationProductRepository
func NewReservationProductRepository(db *gorm.DB) ReservationProductRepository {
	return &reservationProductConnection{
		connection: db,
	}
}
