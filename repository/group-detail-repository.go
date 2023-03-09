package repository

import (
	"fmt"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"gorm.io/gorm"
	"math"
	"strings"
)

// GroupDetailRepository is contract what groupDetailRepository can do to db
type GroupDetailRepository interface {
	GetAllGroupDetail(hotelID string, filterPagination dto.FilterPagination) ([]entity.GroupDetail, dto.Pagination, error)
	InsertGroupDetail(groupDetail entity.GroupDetail) (entity.GroupDetail, error)
	UpdateGroupDetail(groupDetail entity.GroupDetail) (entity.GroupDetail, error)
	FindGroupDetailByID(groupDetailID string) (entity.GroupDetail, error)
	DeleteGroupDetail(groupDetail entity.GroupDetail) error
}

// groupDetailConnection adalah func untuk melakukan query data ke tabel groupDetail
type groupDetailConnection struct {
	connection *gorm.DB
}

func (db *groupDetailConnection) DeleteGroupDetail(groupDetail entity.GroupDetail) error {
	err := db.connection.Where("id_group_detail = ?", groupDetail.IDGroupDetail).Delete(&groupDetail).Error
	return err
}

func (db *groupDetailConnection) GetAllGroupDetail(hotelID string, filterPagination dto.FilterPagination) ([]entity.GroupDetail, dto.Pagination, error) {
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

	var groupDetails []entity.GroupDetail
	query := db.connection.Model(&groupDetails).Joins("LEFT JOIN groups ON group_details.group_id = groups.id_group").
		Joins("LEFT JOIN species ON group_details.species_id = species.id_species")

	whereClause := db.connection.Scopes(func(db *gorm.DB) *gorm.DB {
		if search != "" {
			keyword := strings.ToLower(search)
			if keyword != "" {
				query = query.Where("LOWER(groups.name) LIKE ?", fmt.Sprintf("%%%s%%", keyword)).
					Or("LOWER(species.name) LIKE ?", fmt.Sprintf("%%%s%%", keyword))
			}
		}
		return db
	})

	query.Where(whereClause).Scopes(func(db *gorm.DB) *gorm.DB {
		if filterPagination.GroupID != "" {
			db.Where("group_details.group_id = ?", filterPagination.GroupID)
		}
		if filterPagination.SpeciesID != "" {
			db.Where("group_details.species_id = ?", filterPagination.SpeciesID)
		}
		if filterPagination.HotelID != "" {
			db.Where("group_details.hotel_id = ?", filterPagination.HotelID)
		} else {
			db.Where("group_details.hotel_id = ?", hotelID)
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
		Preload("Group").
		Preload("Species").
		Preload("Hotel").
		Find(&groupDetails).Count(&total).Error

	totalPage := float64(total) / float64(perPage)

	pagination := dto.Pagination{
		Page:      uint(page),
		PerPage:   uint(perPage),
		TotalData: uint(total),
		TotalPage: uint(math.Ceil(totalPage)),
	}

	return groupDetails, pagination, err
}

// InsertGroupDetail is to add groupDetail in database
func (db *groupDetailConnection) InsertGroupDetail(groupDetail entity.GroupDetail) (entity.GroupDetail, error) {
	err := db.connection.Save(&groupDetail).Error
	db.connection.Find(&groupDetail)
	return groupDetail, err
}

// UpdateGroupDetail is func to edit groupDetail in database
func (db *groupDetailConnection) UpdateGroupDetail(groupDetail entity.GroupDetail) (entity.GroupDetail, error) {
	err := db.connection.Where("id_group_detail = ?", groupDetail.IDGroupDetail).Updates(&groupDetail).Error
	db.connection.Where("id_group_detail = ?", groupDetail.IDGroupDetail).Find(&groupDetail)
	return groupDetail, err
}

// FindGroupDetailByID is func to get groupDetail by email
func (db *groupDetailConnection) FindGroupDetailByID(groupDetailID string) (entity.GroupDetail, error) {
	var groupDetail entity.GroupDetail
	err := db.connection.Where("id_group_detail = ?", groupDetailID).Take(&groupDetail).Error
	return groupDetail, err
}

// NewGroupDetailRepository is creates a new instance of GroupDetailRepository
func NewGroupDetailRepository(db *gorm.DB) GroupDetailRepository {
	return &groupDetailConnection{
		connection: db,
	}
}
