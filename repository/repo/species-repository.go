package repository

import (
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"gorm.io/gorm"
)

// SpeciesRepository is contract what speciesRepository can do to db
type SpeciesRepository interface {
	GetAllSpecies(hotelID string) ([]entity.Species, error)
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

func (db *speciesConnection) GetAllSpecies(hotelID string) ([]entity.Species, error) {
	var species []entity.Species
	err := db.connection.Where("hotel_id = ?", hotelID).Preload("Category").Find(&species).Error
	return species, err
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
