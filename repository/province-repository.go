package repository

import (
	"fmt"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"gorm.io/gorm"
	"math"
)

// ProvinceRepository is contract what provinceRepository can do to db
type ProvinceRepository interface {
	GetAllProvince(filterPagination dto.FilterPagination) ([]entity.Province, dto.Pagination, error)
	InsertProvince(province entity.Province) (entity.Province, error)
	UpdateProvince(province entity.Province) (entity.Province, error)
	FindProvinceByID(provinceID string) (entity.Province, error)
	DeleteProvince(province entity.Province) error
}

// provinceConnection adalah func untuk melakukan query data ke tabel province
type provinceConnection struct {
	connection *gorm.DB
}

func (db *provinceConnection) DeleteProvince(province entity.Province) error {
	err := db.connection.Where("id_province = ?", province.IDProvince).Delete(&province).Error
	return err
}

func (db *provinceConnection) GetAllProvince(filterPagination dto.FilterPagination) ([]entity.Province, dto.Pagination, error) {
	var provinces []entity.Province
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

	query := "SELECT * FROM provinces"
	if search != "" {
		query = fmt.Sprintf("%s WHERE name LIKE '%%%s%%'", query, search)
	}
	query = fmt.Sprintf("%s ORDER BY %s %s", query, sortBy, orderBy)

	var total int64

	db.connection.Raw(query).Count(&total)
	query = fmt.Sprintf("%s LIMIT %d OFFSET %d", query, perPage, (page-1)*perPage)
	err := db.connection.Raw(query).Scan(&provinces).Error

	totalPage := float64(total) / float64(perPage)

	pagination := dto.Pagination{
		Page:      uint(page),
		PerPage:   uint(perPage),
		TotalData: uint(total),
		TotalPage: uint(math.Ceil(totalPage)),
	}

	return provinces, pagination, err
}

// InsertProvince is to add province in database
func (db *provinceConnection) InsertProvince(province entity.Province) (entity.Province, error) {
	err := db.connection.Save(&province).Error
	return province, err
}

// UpdateProvince is func to edit province in database
func (db *provinceConnection) UpdateProvince(province entity.Province) (entity.Province, error) {
	err := db.connection.Where("id_province = ?", province.IDProvince).Updates(&province).Error
	return province, err
}

// FindProvinceByID is func to get province by email
func (db *provinceConnection) FindProvinceByID(provinceID string) (entity.Province, error) {
	var province entity.Province
	err := db.connection.Where("id_province = ?", provinceID).Take(&province).Error
	return province, err
}

// NewProvinceRepository is creates a new instance of ProvinceRepository
func NewProvinceRepository(db *gorm.DB) ProvinceRepository {
	return &provinceConnection{
		connection: db,
	}
}
