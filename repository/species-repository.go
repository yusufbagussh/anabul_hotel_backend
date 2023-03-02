package repository

import (
	"fmt"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"gorm.io/gorm"
	"math"
	"strings"
)

// SpeciesRepository is contract what speciesRepository can do to db
type SpeciesRepository interface {
	GetAllSpecies(filterPagination dto.FilterPagination) ([]entity.Species, dto.Pagination, error)
	InsertSpecies(species entity.Species) (entity.Species, error)
	UpdateSpecies(species entity.Species) (entity.Species, error)
	FindSpeciesByID(speciesID string) (entity.Species, error)
	DeleteSpecies(species entity.Species) error
}

// speciesConnection adalah func untuk melakukan query data ke tabel species
type speciesConnection struct {
	connection *gorm.DB
}

func (db *speciesConnection) DeleteSpecies(species entity.Species) error {
	err := db.connection.Where("id_species = ?", species.IDSpecies).Delete(&species).Error
	return err
}

func (db *speciesConnection) GetAllSpecies(filterPagination dto.FilterPagination) ([]entity.Species, dto.Pagination, error) {
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

	var species []entity.Species
	query := db.connection.Joins("JOIN Categories ON species.category_id = categories.id_category")

	whereClause := db.connection.Scopes(func(db *gorm.DB) *gorm.DB {
		if search != "" {
			keyword := strings.ToLower(search)
			if keyword != "" {
				db.Where("LOWER(species.name) LIKE ?", fmt.Sprintf("%%%s%%", keyword)).
					Or("LOWER(categories.name) LIKE ?", fmt.Sprintf("%%%s%%", keyword))
			}
		}
		return db
	})

	query.Where(whereClause).Scopes(func(db *gorm.DB) *gorm.DB {
		if filterPagination.CategoryID != "" {
			db.Where("species.category_id = ?", filterPagination.CategoryID)
		}
		if filterPagination.ClassID != "" {
			db.Where("species.class_id = ?", filterPagination.ClassID)
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

	query.Find(&species).Count(&total)

	err := query.Limit(perPage).Offset((page - 1) * perPage).Preload("Categories").Find(&species).Error

	totalPage := float64(total) / float64(perPage)

	pagination := dto.Pagination{
		Page:      uint(page),
		PerPage:   uint(perPage),
		TotalData: uint(total),
		TotalPage: uint(math.Ceil(totalPage)),
	}
	return species, pagination, err
}

// InsertSpecies is to add species in database
func (db *speciesConnection) InsertSpecies(species entity.Species) (entity.Species, error) {
	err := db.connection.Save(&species).Error
	db.connection.Preload("Category").Find(&species)
	return species, err
}

// UpdateSpecies is func to edit species in database
func (db *speciesConnection) UpdateSpecies(species entity.Species) (entity.Species, error) {
	err := db.connection.Where("id_species = ?", species.IDSpecies).Updates(&species).Error
	db.connection.Where("id_species = ?", species.IDSpecies).Preload("Category").Find(&species)
	return species, err
}

// FindSpeciesByID is func to get species by email
func (db *speciesConnection) FindSpeciesByID(speciesID string) (entity.Species, error) {
	var species entity.Species
	err := db.connection.Where("id_species = ?", speciesID).Preload("Category").Take(&species).Error
	return species, err
}

// NewSpeciesRepository is creates a new instance of SpeciesRepository
func NewSpeciesRepository(db *gorm.DB) SpeciesRepository {
	return &speciesConnection{
		connection: db,
	}
}
