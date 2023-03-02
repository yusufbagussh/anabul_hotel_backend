package seeder

import (
	"github.com/yusufbagussh/pet_hotel_backend/entity"
)

func SpeciesSeeder(categories *[]entity.Category) *[]entity.Species {
	species1 := entity.Species{
		CategoryID: (*categories)[0].IDCategory,
		//HotelID:    (*hotels)[0].IDHotel,
		Name: "Dalmation",
	}
	species2 := entity.Species{
		CategoryID: (*categories)[0].IDCategory,
		//HotelID:    (*hotels)[0].IDHotel,
		Name: "Anggora",
	}
	species := []entity.Species{
		species1,
		species2,
	}
	return &species
}
