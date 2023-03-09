package seeder

import (
	"github.com/yusufbagussh/pet_hotel_backend/entity"
)

func CageSeeder(cageDetails *[]entity.CageDetail, hotels *[]entity.Hotel) *[]entity.Cage {
	cageCategory1 := entity.Cage{
		CageDetailID: (*cageDetails)[0].IDCageDetail,
		Name:         "Kandang 1",
		Status:       "Terisi",
		HotelID:      (*hotels)[0].IDHotel,
	}
	cageCategory2 := entity.Cage{
		CageDetailID: (*cageDetails)[1].IDCageDetail,
		Name:         "Kandang A",
		Status:       "Terisi",
		HotelID:      (*hotels)[0].IDHotel,
	}
	cageCategories := []entity.Cage{
		cageCategory1,
		cageCategory2,
	}
	return &cageCategories
}
