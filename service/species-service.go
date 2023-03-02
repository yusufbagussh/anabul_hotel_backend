package service

import (
	"github.com/mashingan/smapping"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"github.com/yusufbagussh/pet_hotel_backend/repository"
)

// SpeciesService is a contract of what speciesService can do
type SpeciesService interface {
	GetAllSpecies(filterPagination dto.FilterPagination) ([]entity.Species, dto.Pagination, error)
	CreateSpecies(species dto.CreateSpecies) (entity.Species, error)
	UpdateSpecies(species dto.UpdateSpecies) (entity.Species, error)
	DeleteSpecies(speciesID string) error
	ShowSpecies(speciesID string) (entity.Species, error)
}

type speciesService struct {
	speciesRepository repository.SpeciesRepository
}

// NewSpeciesService creates a new instance of SpeciesService
func NewSpeciesService(speciesRepo repository.SpeciesRepository) SpeciesService {
	return &speciesService{
		speciesRepository: speciesRepo,
	}
}

func (u *speciesService) CreateSpecies(species dto.CreateSpecies) (entity.Species, error) {
	speciesToCreate := entity.Species{}
	errMap := smapping.FillStruct(&speciesToCreate, smapping.MapFields(&species))
	if errMap != nil {
		return speciesToCreate, errMap
	}
	updatedSpecies, err := u.speciesRepository.InsertSpecies(speciesToCreate)
	return updatedSpecies, err
}
func (u *speciesService) UpdateSpecies(species dto.UpdateSpecies) (entity.Species, error) {
	speciesToUpdate := entity.Species{}
	errMap := smapping.FillStruct(&speciesToUpdate, smapping.MapFields(&species))
	if errMap != nil {
		return speciesToUpdate, errMap
	}
	updatedSpecies, err := u.speciesRepository.UpdateSpecies(speciesToUpdate)
	return updatedSpecies, err
}
func (u *speciesService) DeleteSpecies(speciesID string) error {
	species, err := u.speciesRepository.FindSpeciesByID(speciesID)
	if err != nil {
		return err
	}
	errDel := u.speciesRepository.DeleteSpecies(species)
	return errDel
}
func (u *speciesService) ShowSpecies(speciesID string) (entity.Species, error) {
	result, err := u.speciesRepository.FindSpeciesByID(speciesID)
	return result, err
}
func (u *speciesService) GetAllSpecies(filterPagination dto.FilterPagination) ([]entity.Species, dto.Pagination, error) {
	admins, pagination, err := u.speciesRepository.GetAllSpecies(filterPagination)
	return admins, pagination, err
}
