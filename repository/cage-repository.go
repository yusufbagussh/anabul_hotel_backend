package repository

import (
	"fmt"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"gorm.io/gorm"
	"math"
	"strings"
)

// CageRepository is contract what cageRepository can do to db
type CageRepository interface {
	GetAllCage(hotelID string, filterPagination dto.FilterPagination) ([]entity.Cage, dto.Pagination, error)
	InsertCage(cage entity.Cage) (entity.Cage, error)
	UpdateCage(cage entity.Cage) (entity.Cage, error)
	FindCageByID(cageID string) (entity.Cage, error)
	DeleteCage(cage entity.Cage) error
}

// cageConnection adalah func untuk melakukan query data ke tabel cage
type cageConnection struct {
	connection *gorm.DB
}

func (db *cageConnection) DeleteCage(cage entity.Cage) error {
	err := db.connection.Where("id_cage = ?", cage.IDCage).Delete(&cage).Error
	return err
}

func (db *cageConnection) GetAllCage(hotelID string, filterPagination dto.FilterPagination) ([]entity.Cage, dto.Pagination, error) {
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

	var cages []entity.Cage
	query := db.connection.Model(&cages).Joins("LEFT JOIN cage_details ON cages.cage_detail_id = cage_details.id_cage_detail").
		Joins("LEFT JOIN cage_categories ON cage_details.cage_category_id = cage_categories.id_cage_category").
		Joins("LEFT JOIN cage_types ON cage_details.cage_type_id = cage_types.id_cage_type")

	whereClause := db.connection.Scopes(func(db *gorm.DB) *gorm.DB {
		if search != "" {
			keyword := strings.ToLower(search)
			if keyword != "" {
				db.Where("LOWER(cages.name) LIKE ?", fmt.Sprintf("%%%s%%", keyword)).
					Or("LOWER(cage_categories.name) LIKE ?", fmt.Sprintf("%%%s%%", keyword)).
					Or("LOWER(cage_types.name) LIKE ?", fmt.Sprintf("%%%s%%", keyword))
			}
		}
		return db
	})

	query.Where(whereClause).Scopes(func(db *gorm.DB) *gorm.DB {
		if filterPagination.CageDetailID != "" {
			db.Where("cages.cage_detail_id = ?", filterPagination.CageDetailID)
		}
		if filterPagination.CageCategoryID != "" {
			db.Where("cage_details.cage_category_id = ?", filterPagination.CageCategoryID)
		}
		if filterPagination.CageTypeID != "" {
			db.Where("cage_details.cage_type_id = ?", filterPagination.CageTypeID)
		}
		if filterPagination.HotelID != "" {
			db.Where("cages.hotel_id = ?", filterPagination.HotelID)
		} else {
			db.Where("cages.hotel_id = ?", hotelID)
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
		Preload("CageDetail").
		//Preload("CageCategory").
		//Preload("CageCategory.CageDetail").
		//Preload("CageType").
		//Preload("CageType.CageDetail").
		Preload("Hotel").
		Find(&cages).Error

	totalPage := float64(total) / float64(perPage)

	pagination := dto.Pagination{
		Page:      uint(page),
		PerPage:   uint(perPage),
		TotalData: uint(total),
		TotalPage: uint(math.Ceil(totalPage)),
	}

	return cages, pagination, err
}

// InsertCage is to add cage in database
func (db *cageConnection) InsertCage(cage entity.Cage) (entity.Cage, error) {
	err := db.connection.Save(&cage).Error
	db.connection.Find(&cage)
	return cage, err
}

// UpdateCage is func to edit cage in database
func (db *cageConnection) UpdateCage(cage entity.Cage) (entity.Cage, error) {
	err := db.connection.Where("id_cage = ?", cage.IDCage).Updates(&cage).Error
	db.connection.Where("id_cage = ?", cage.IDCage).Find(&cage)
	return cage, err
}

// FindCageByID is func to get cage by email
func (db *cageConnection) FindCageByID(cageID string) (entity.Cage, error) {
	var cage entity.Cage
	err := db.connection.Where("id_cage = ?", cageID).Take(&cage).Error
	return cage, err
}

// NewCageRepository is creates a new instance of CageRepository
func NewCageRepository(db *gorm.DB) CageRepository {
	return &cageConnection{
		connection: db,
	}
}
