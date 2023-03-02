package seeder

import (
	"github.com/yusufbagussh/pet_hotel_backend/entity"
)

func CageCategorySeeder(hotels *[]entity.Hotel) *[]entity.CageCategory {
	cageCategory1 := entity.CageCategory{
		HotelID: (*hotels)[0].IDHotel,
		Name:    "Elite",
	}
	cageCategory2 := entity.CageCategory{
		HotelID: (*hotels)[0].IDHotel,
		Name:    "Premium",
	}
	cageCategories := []entity.CageCategory{
		cageCategory1,
		cageCategory2,
	}
	return &cageCategories
}
