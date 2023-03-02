package seeder

import (
	"github.com/yusufbagussh/pet_hotel_backend/entity"
)

func PetSeeder(species *[]entity.Species, users *[]entity.User) *[]entity.Pet {
	pet1 := entity.Pet{
		SpeciesID: (*species)[0].IDSpecies,
		UserID:    (*users)[3].ID,
		Name:      "Mike",
	}
	pet2 := entity.Pet{
		SpeciesID: (*species)[1].IDSpecies,
		UserID:    (*users)[3].ID,
		Name:      "Max",
	}
	cities := []entity.Pet{
		pet1,
		pet2,
	}
	return &cities
}
