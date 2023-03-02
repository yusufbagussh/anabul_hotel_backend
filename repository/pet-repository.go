package repository

import (
	"fmt"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"gorm.io/gorm"
	"math"
	"strings"
)

// PetRepository is contract what petRepository can do to db
type PetRepository interface {
	GetAllPet(customerID string, filterPagination dto.FilterPagination) ([]entity.Pet, dto.Pagination, error)
	InsertPet(pet entity.Pet) (entity.Pet, error)
	UpdatePet(pet entity.Pet) (entity.Pet, error)
	FindPetByID(petID string) (entity.Pet, error)
	DeletePet(pet entity.Pet) error
}

// petConnection adalah func untuk melakukan query data ke tabel pet
type petConnection struct {
	connection *gorm.DB
}

func (db *petConnection) DeletePet(pet entity.Pet) error {
	err := db.connection.Where("id_pet = ?", pet.IDPet).Delete(&pet).Error
	return err
}

func (db *petConnection) GetAllPet(customerID string, filterPagination dto.FilterPagination) ([]entity.Pet, dto.Pagination, error) {
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

	var pets []entity.Pet
	query := db.connection.Joins("JOIN species ON pet.species_id = species.id_species").
		Joins("JOIN categories ON species.category_id = categories.id_category")

	whereClause := db.connection.Scopes(func(db *gorm.DB) *gorm.DB {
		if search != "" {
			keyword := strings.ToLower(search)
			if keyword != "" {
				db.Where("pets.name LIKE ?", fmt.Sprintf("%%%s%%", keyword)).
					Or("species.name LIKE ?", fmt.Sprintf("%%%s%%", keyword)).
					Or("categories.name LIKE ?", fmt.Sprintf("%%%s%%", keyword))
			}
		}
		return db
	})

	query.Where(whereClause).Scopes(func(db *gorm.DB) *gorm.DB {
		if filterPagination.ClassID != "" {
			db.Where("pet.class_id = ?", filterPagination.ClassID)
		}
		if filterPagination.CategoryID != "" {
			db.Where("pet.category_id = ?", filterPagination.CategoryID)
		}
		if filterPagination.SpeciesID != "" {
			db.Where("pet.species_id = ?", filterPagination.SpeciesID)
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

	err := query.Where("user_id = ?", customerID).Limit(perPage).Offset((page - 1) * perPage).Preload("Hotel").Find(&pets).Count(&total).Error

	totalPage := float64(total) / float64(perPage)

	pagination := dto.Pagination{
		Page:      uint(page),
		PerPage:   uint(perPage),
		TotalData: uint(total),
		TotalPage: uint(math.Ceil(totalPage)),
	}

	return pets, pagination, err
}

// InsertPet is to add pet in database
func (db *petConnection) InsertPet(pet entity.Pet) (entity.Pet, error) {
	err := db.connection.Save(&pet).Error
	db.connection.Preload("Hotel").Find(&pet)
	return pet, err
}

// UpdatePet is func to edit pet in database
func (db *petConnection) UpdatePet(pet entity.Pet) (entity.Pet, error) {
	err := db.connection.Where("id_pet = ?", pet.IDPet).Updates(&pet).Error
	db.connection.Where("id_pet = ?", pet.IDPet).Preload("User").Find(&pet)
	return pet, err
}

// FindPetByID is func to get pet by email
func (db *petConnection) FindPetByID(petID string) (entity.Pet, error) {
	var pet entity.Pet
	err := db.connection.Where("id_pet = ?", petID).Preload("User").Take(&pet).Error
	return pet, err
}

// NewPetRepository is creates a new instance of PetRepository
func NewPetRepository(db *gorm.DB) PetRepository {
	return &petConnection{
		connection: db,
	}
}
