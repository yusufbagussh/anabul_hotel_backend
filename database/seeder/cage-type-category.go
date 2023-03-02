package seeder

import (
	"github.com/yusufbagussh/pet_hotel_backend/entity"
)

func CageTypeSeeder(hotels *[]entity.Hotel) *[]entity.CageType {
	cageType1 := entity.CageType{
		HotelID: (*hotels)[0].IDHotel,
		Name:    "Small",
		Length:  100,
		Width:   50,
		Height:  50,
	}
	cageType2 := entity.CageType{
		HotelID: (*hotels)[0].IDHotel,
		Name:    "Large",
		Length:  200,
		Width:   100,
		Height:  100,
	}
	cageCategories := []entity.CageType{
		cageType1,
		cageType2,
	}
	return &cageCategories
}
