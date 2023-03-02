package service

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"github.com/yusufbagussh/pet_hotel_backend/repository"
	"os"
	"path/filepath"
)

// PetService is a contract of what petService can do
type PetService interface {
	GetAllPet(customerID string, pagePagination dto.FilterPagination) ([]entity.Pet, dto.Pagination, error)
	CreatePet(pet dto.CreatePet, ctx *gin.Context) (entity.Pet, error)
	UpdatePet(pet dto.UpdatePet, ctx *gin.Context) (entity.Pet, error)
	DeletePet(petID string, userID string) (error, interface{})
	ShowPet(petID string, userID string) (entity.Pet, error, interface{})
	//IsAllowedPet(petID string, userID string) bool
}

type petService struct {
	petRepository repository.PetRepository
}

// NewPetService creates a new instance of PetService
func NewPetService(petRepo repository.PetRepository) PetService {
	return &petService{
		petRepository: petRepo,
	}
}

func (u *petService) CreatePet(pet dto.CreatePet, ctx *gin.Context) (entity.Pet, error) {
	petToCreate := entity.Pet{}
	if pet.Image != nil {
		extension := filepath.Ext(pet.Image.Filename) // Generate random file name for the new uploaded file, so it doesn't override the old file with same name
		newImage := uuid.New().String() + extension
		petToCreate.Image = newImage

		err := ctx.SaveUploadedFile(pet.Image, "uploads/user-profiles/"+newImage)
		if err != nil {
			return petToCreate, err
		}
	}
	petToCreate.Name = pet.Name
	petToCreate.SpeciesID = pet.SpeciesID
	petToCreate.FavoriteFood = pet.FavoriteFood
	petToCreate.BirthDate = pet.BirthDate
	petToCreate.Gender = pet.Gender
	petToCreate.UserID = pet.UserID
	updatedPet, err := u.petRepository.InsertPet(petToCreate)
	return updatedPet, err
}
func (u *petService) UpdatePet(pet dto.UpdatePet, ctx *gin.Context) (entity.Pet, error) {
	petToUpdate := entity.Pet{}
	if pet.Image != nil {
		extension := filepath.Ext(pet.Image.Filename) // Generate random file name for the new uploaded file, so it doesn't override the old file with same name
		newImage := uuid.New().String() + extension
		petToUpdate.Image = newImage

		err := ctx.SaveUploadedFile(pet.Image, "uploads/user-profiles/"+newImage)
		if err != nil {
			return petToUpdate, err
		} else {
			petByID, _ := u.petRepository.FindPetByID(pet.IDPet)
			if petByID.Image != "" {
				errRemove := os.Remove("./uploads/user-profiles/" + petByID.Image)
				if errRemove != nil {
					return petToUpdate, errRemove
				}
			}
		}
	}
	petToUpdate.IDPet = pet.IDPet
	petToUpdate.Name = pet.Name
	petToUpdate.SpeciesID = pet.SpeciesID
	petToUpdate.FavoriteFood = pet.FavoriteFood
	petToUpdate.BirthDate = pet.BirthDate
	petToUpdate.Gender = pet.Gender
	petToUpdate.UserID = pet.UserID
	updatedPet, err := u.petRepository.UpdatePet(petToUpdate)
	return updatedPet, err
}
func (u *petService) DeletePet(petID string, userID string) (error, interface{}) {
	pet, err := u.petRepository.FindPetByID(petID)
	if pet.UserID != userID {
		return nil, false
	}
	if err != nil {
		return err, nil
	}
	if pet.Image != "" {
		errRemove := os.Remove("./uploads/user-profiles/" + pet.Image)
		if errRemove != nil {
			return errRemove, nil
		}
	}
	errDel := u.petRepository.DeletePet(pet)
	return errDel, nil
}
func (u *petService) ShowPet(petID string, userID string) (entity.Pet, error, interface{}) {
	result, err := u.petRepository.FindPetByID(petID)
	if result.UserID != userID {
		return entity.Pet{}, nil, false
	}
	return result, err, nil
}
func (u *petService) GetAllPet(customerID string, pagePagination dto.FilterPagination) ([]entity.Pet, dto.Pagination, error) {
	admins, pagination, err := u.petRepository.GetAllPet(customerID, pagePagination)
	return admins, pagination, err
}
