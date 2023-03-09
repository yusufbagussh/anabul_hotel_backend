package repository

import (
	"fmt"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"gorm.io/gorm"
	"math"
	"strings"
)

// CageDetailRepository is contract what cageDetailRepository can do to db
type CageDetailRepository interface {
	GetAllCageDetail(hotelID string, filterPagination dto.FilterPagination) ([]entity.CageDetail, dto.Pagination, error)
	InsertCageDetail(cageDetail entity.CageDetail) (entity.CageDetail, error)
	UpdateCageDetail(cageDetail entity.CageDetail) (entity.CageDetail, error)
	FindCageDetailByID(cageDetailID string) (entity.CageDetail, error)
	DeleteCageDetail(cageDetail entity.CageDetail) error
	UpdateCageDetailStatus(productStatus dto.UpdateCageDetailStatus) (entity.CageDetail, error)
}

// cageDetailConnection adalah func untuk melakukan query data ke tabel cageDetail
type cageDetailConnection struct {
	connection *gorm.DB
}

func (db *cageDetailConnection) DeleteCageDetail(cageDetail entity.CageDetail) error {
	err := db.connection.Where("id_cage_detail = ?", cageDetail.IDCageDetail).Delete(&cageDetail).Error
	return err
}

func (db *cageDetailConnection) UpdateCageDetailStatus(productStatus dto.UpdateCageDetailStatus) (entity.CageDetail, error) {
	var cageDetail entity.CageDetail
	err := db.connection.Model(&cageDetail).Where("id_cage_detail = ?", productStatus.IDCageDetail).Updates(&entity.CageDetail{Status: productStatus.Status}).Error
	db.connection.Find(&cageDetail)
	return cageDetail, err
}

func (db *cageDetailConnection) GetAllCageDetail(hotelID string, filterPagination dto.FilterPagination) ([]entity.CageDetail, dto.Pagination, error) {
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

	var cageDetails []entity.CageDetail
	query := db.connection.Model(&cageDetails).Joins("LEFT JOIN cage_categories ON cage_details.cage_category_id = cage_categories.id_cage_category").
		Joins("LEFT JOIN cage_types ON cage_details.cage_type_id = cage_types.id_cage_type")

	whereClause := db.connection.Scopes(func(db *gorm.DB) *gorm.DB {
		if search != "" {
			keyword := strings.ToLower(search)
			if keyword != "" {
				db.Where("LOWER(cage_categories.name) LIKE ?", fmt.Sprintf("%%%s%%", keyword)).
					Or("LOWER(cage_types.name) LIKE ?", fmt.Sprintf("%%%s%%", keyword))
			}
		}
		return db
	})

	query.Where(whereClause).Scopes(func(db *gorm.DB) *gorm.DB {
		if filterPagination.CageCategoryID != "" {
			db.Where("cage_details.cage_category_id = ?", filterPagination.CageCategoryID)
		}
		if filterPagination.CageTypeID != "" {
			db.Where("cage_details.cage_type_id = ?", filterPagination.CageTypeID)
		}
		if filterPagination.HotelID != "" {
			db.Where("cage_details.hotel_id = ?", filterPagination.HotelID)
		} else {
			db.Where("cage_details.hotel_id = ?", hotelID)
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

	err := query.Count(&total).Limit(perPage).Offset((page - 1) * perPage).
		Preload("CageCategory").
		Preload("CageType").
		Preload("Hotel").
		Find(&cageDetails).Count(&total).Error

	totalPage := float64(total) / float64(perPage)

	pagination := dto.Pagination{
		Page:      uint(page),
		PerPage:   uint(perPage),
		TotalData: uint(total),
		TotalPage: uint(math.Ceil(totalPage)),
	}

	return cageDetails, pagination, err
}

// InsertCageDetail is to add cageDetail in database
func (db *cageDetailConnection) InsertCageDetail(cageDetail entity.CageDetail) (entity.CageDetail, error) {
	err := db.connection.Save(&cageDetail).Error
	db.connection.Find(&cageDetail)
	return cageDetail, err
}

// UpdateCageDetail is func to edit cageDetail in database
func (db *cageDetailConnection) UpdateCageDetail(cageDetail entity.CageDetail) (entity.CageDetail, error) {
	err := db.connection.Where("id_cage_detail = ?", cageDetail.IDCageDetail).Updates(&cageDetail).Error
	db.connection.Where("id_cage_detail = ?", cageDetail.IDCageDetail).Find(&cageDetail)
	return cageDetail, err
}

// FindCageDetailByID is func to get cageDetail by email
func (db *cageDetailConnection) FindCageDetailByID(cageDetailID string) (entity.CageDetail, error) {
	var cageDetail entity.CageDetail
	err := db.connection.Where("id_cage_detail = ?", cageDetailID).Take(&cageDetail).Error
	return cageDetail, err
}

// NewCageDetailRepository is creates a new instance of CageDetailRepository
func NewCageDetailRepository(db *gorm.DB) CageDetailRepository {
	return &cageDetailConnection{
		connection: db,
	}
}
