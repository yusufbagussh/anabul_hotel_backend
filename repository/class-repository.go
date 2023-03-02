package repository

import (
	"fmt"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"gorm.io/gorm"
	"math"
)

// ClassRepository is contract what classRepository can do to db
type ClassRepository interface {
	GetAllClass(filterPagination dto.FilterPagination) ([]entity.Class, dto.Pagination, error)
	InsertClass(class entity.Class) (entity.Class, error)
	UpdateClass(class entity.Class) (entity.Class, error)
	FindClassByID(classID string) (entity.Class, error)
	DeleteClass(class entity.Class) error
}

// classConnection adalah func untuk melakukan query data ke tabel class
type classConnection struct {
	connection *gorm.DB
}

func (db *classConnection) DeleteClass(class entity.Class) error {
	err := db.connection.Where("id_class = ?", class.IDClass).Delete(&class).Error
	return err
}

func (db *classConnection) GetAllClass(filterPagination dto.FilterPagination) ([]entity.Class, dto.Pagination, error) {
	var classes []entity.Class
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

	query := "SELECT * FROM classes"
	if search != "" {
		query = fmt.Sprintf("%s WHERE LOWER(name) LIKE '%%%s%%'", query, search)
	}
	query = fmt.Sprintf("%s ORDER BY %s %s", query, sortBy, orderBy)

	var total int64

	db.connection.Raw(query).Count(&total)
	query = fmt.Sprintf("%s LIMIT %d OFFSET %d", query, perPage, (page-1)*perPage)
	err := db.connection.Raw(query).Scan(&classes).Error

	totalPage := float64(total) / float64(perPage)

	pagination := dto.Pagination{
		Page:      uint(page),
		PerPage:   uint(perPage),
		TotalData: uint(total),
		TotalPage: uint(math.Ceil(totalPage)),
	}

	return classes, pagination, err
}

// InsertClass is to add class in database
func (db *classConnection) InsertClass(class entity.Class) (entity.Class, error) {
	err := db.connection.Save(&class).Error
	return class, err
}

// UpdateClass is func to edit class in database
func (db *classConnection) UpdateClass(class entity.Class) (entity.Class, error) {
	err := db.connection.Where("id_class = ?", class.IDClass).Updates(&class).Error
	return class, err
}

// FindClassByID is func to get class by email
func (db *classConnection) FindClassByID(classID string) (entity.Class, error) {
	var class entity.Class
	err := db.connection.Where("id_class = ?", classID).Take(&class).Error
	return class, err
}

// NewClassRepository is creates a new instance of ClassRepository
func NewClassRepository(db *gorm.DB) ClassRepository {
	return &classConnection{
		connection: db,
	}
}
