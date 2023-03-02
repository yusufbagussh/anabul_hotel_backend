package repository

import (
	"fmt"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"gorm.io/gorm"
	"math"
	"strings"
)

// GroupRepository is contract what groupRepository can do to db
type GroupRepository interface {
	GetAllGroup(hotelID string, filterPagination dto.FilterPagination) ([]entity.Group, dto.Pagination, error)
	InsertGroup(group entity.Group) (entity.Group, error)
	UpdateGroup(group entity.Group) (entity.Group, error)
	FindGroupByID(groupID string) (entity.Group, error)
	DeleteGroup(group entity.Group) error
}

// groupConnection adalah func untuk melakukan query data ke tabel group
type groupConnection struct {
	connection *gorm.DB
}

func (db *groupConnection) DeleteGroup(group entity.Group) error {
	err := db.connection.Where("id_group = ?", group.IDGroup).Delete(&group).Error
	return err
}

func (db *groupConnection) GetAllGroup(hotelID string, filterPagination dto.FilterPagination) ([]entity.Group, dto.Pagination, error) {
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

	var groups []entity.Group
	query := db.connection

	if search != "" {
		keyword := strings.ToLower(search)
		if keyword != "" {
			query = query.Where("LOWER(groups.name) LIKE ?", fmt.Sprintf("%%%s%%", keyword))
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

	err := query.Where("hotel_id = ?", hotelID).Limit(perPage).Offset((page - 1) * perPage).Preload("Hotel").Find(&groups).Count(&total).Error

	totalPage := float64(total) / float64(perPage)

	pagination := dto.Pagination{
		Page:      uint(page),
		PerPage:   uint(perPage),
		TotalData: uint(total),
		TotalPage: uint(math.Ceil(totalPage)),
	}

	return groups, pagination, err
}

// InsertGroup is to add group in database
func (db *groupConnection) InsertGroup(group entity.Group) (entity.Group, error) {
	err := db.connection.Save(&group).Error
	db.connection.Find(&group)
	return group, err
}

// UpdateGroup is func to edit group in database
func (db *groupConnection) UpdateGroup(group entity.Group) (entity.Group, error) {
	err := db.connection.Where("id_group = ?", group.IDGroup).Updates(&group).Error
	db.connection.Where("id_group = ?", group.IDGroup).Find(&group)
	return group, err
}

// FindGroupByID is func to get group by email
func (db *groupConnection) FindGroupByID(groupID string) (entity.Group, error) {
	var group entity.Group
	err := db.connection.Where("id_group = ?", groupID).Take(&group).Error
	return group, err
}

// NewGroupRepository is creates a new instance of GroupRepository
func NewGroupRepository(db *gorm.DB) GroupRepository {
	return &groupConnection{
		connection: db,
	}
}
