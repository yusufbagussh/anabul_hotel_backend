package seeder

import (
	"github.com/yusufbagussh/pet_hotel_backend/entity"
)

func CageDetailSeeder(categories *[]entity.CageCategory, types *[]entity.CageType, hotels *[]entity.Hotel) *[]entity.CageDetail {
	cageDetail1 := entity.CageDetail{
		CageCategoryID: (*categories)[0].IDCageCategory,
		CageTypeID:     (*types)[0].IDCageType,
		HotelID:        (*hotels)[0].IDHotel,
		Quantity:       10,
		Status:         "Tersedia",
		Price:          150000,
	}
	cageDetail2 := entity.CageDetail{
		CageCategoryID: (*categories)[0].IDCageCategory,
		CageTypeID:     (*types)[1].IDCageType,
		HotelID:        (*hotels)[0].IDHotel,
		Quantity:       10,
		Status:         "Tersedia",
		Price:          200000,
	}
	cageDetails := []entity.CageDetail{
		cageDetail1,
		cageDetail2,
	}
	return &cageDetails
}
