package repository

import (
	"fmt"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"gorm.io/gorm"
	"math"
	"strings"
)

// ReservationDetailRepository is contract what reservationDetailRepository can do to db
type ReservationDetailRepository interface {
	GetAllReservationDetail(filterPagination dto.FilterPagination) ([]entity.ReservationDetail, dto.Pagination, error)
	InsertReservationDetail(reservationDetail entity.ReservationDetail) (entity.ReservationDetail, error)
	UpdateReservationDetail(reservationDetail entity.ReservationDetail) (entity.ReservationDetail, error)
	FindReservationDetailByID(reservationDetailID string) (entity.ReservationDetail, error)
	DeleteReservationDetail(reservationDetail entity.ReservationDetail) error
}

// reservationDetailConnection adalah func untuk melakukan query data ke tabel reservationDetail
type reservationDetailConnection struct {
	connection *gorm.DB
}

func (db *reservationDetailConnection) DeleteReservationDetail(reservationDetail entity.ReservationDetail) error {
	err := db.connection.Where("id_reservationDetail = ?", reservationDetail.IDReservationDetail).Delete(&reservationDetail).Error
	return err
}

func (db *reservationDetailConnection) GetAllReservationDetail(filterPagination dto.FilterPagination) ([]entity.ReservationDetail, dto.Pagination, error) {
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

	var reservationDetails []entity.ReservationDetail
	query := db.connection

	if search != "" {
		keyword := strings.ToLower(search)
		if keyword != "" {
			query = query.Where("LOWER(reservationDetails.name) LIKE ?", fmt.Sprintf("%%%s%%", keyword))
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

	err := query.Limit(perPage).Offset((page - 1) * perPage).Preload("Hotel").Find(&reservationDetails).Count(&total).Error

	totalPage := float64(total) / float64(perPage)

	pagination := dto.Pagination{
		Page:      uint(page),
		PerPage:   uint(perPage),
		TotalData: uint(total),
		TotalPage: uint(math.Ceil(totalPage)),
	}

	return reservationDetails, pagination, err
}

// InsertReservationDetail is to add reservationDetail in database
func (db *reservationDetailConnection) InsertReservationDetail(reservationDetail entity.ReservationDetail) (entity.ReservationDetail, error) {
	err := db.connection.Save(&reservationDetail).Error
	db.connection.Find(&reservationDetail)
	return reservationDetail, err
}

// UpdateReservationDetail is func to edit reservationDetail in database
func (db *reservationDetailConnection) UpdateReservationDetail(reservationDetail entity.ReservationDetail) (entity.ReservationDetail, error) {
	err := db.connection.Where("id_reservationDetail = ?", reservationDetail.IDReservationDetail).Updates(&reservationDetail).Error
	db.connection.Where("id_reservationDetail = ?", reservationDetail.IDReservationDetail).Find(&reservationDetail)
	return reservationDetail, err
}

// FindReservationDetailByID is func to get reservationDetail by email
func (db *reservationDetailConnection) FindReservationDetailByID(reservationDetailID string) (entity.ReservationDetail, error) {
	var reservationDetail entity.ReservationDetail
	err := db.connection.Where("id_reservationDetail = ?", reservationDetailID).Take(&reservationDetail).Error
	return reservationDetail, err
}

// NewReservationDetailRepository is creates a new instance of ReservationDetailRepository
func NewReservationDetailRepository(db *gorm.DB) ReservationDetailRepository {
	return &reservationDetailConnection{
		connection: db,
	}
}
