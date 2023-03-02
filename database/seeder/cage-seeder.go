package seeder

import (
	"github.com/yusufbagussh/pet_hotel_backend/entity"
)

func CageSeeder(categoryTypes *[]entity.CageDetail, hotels *[]entity.Hotel) *[]entity.Cage {
	cageCategory1 := entity.Cage{
		CageDetailID: (*categoryTypes)[0].IDCageDetail,
		Name:         "Kandang 1",
		Status:       "Used",
		HotelID:      (*hotels)[0].IDHotel,
	}
	cageCategory2 := entity.Cage{
		CageDetailID: (*categoryTypes)[0].IDCageDetail,
		Name:         "Kandang 2",
		Status:       "Used",
		HotelID:      (*hotels)[0].IDHotel,
	}
	cageCategories := []entity.Cage{
		cageCategory1,
		cageCategory2,
	}
	return &cageCategories
}
