package seeder

import (
	"github.com/yusufbagussh/pet_hotel_backend/entity"
)

func ProductSeeder(hotels *[]entity.Hotel) *[]entity.Product {
	product1 := entity.Product{
		HotelID: (*hotels)[0].IDHotel,
		Name:    "Whiskas",
		Price:   20000,
		Status:  "Tersedia",
	}
	product2 := entity.Product{
		HotelID: (*hotels)[0].IDHotel,
		Name:    "Bolt",
		Price:   30000,
		Status:  "Tersedia",
	}
	products := []entity.Product{
		product1,
		product2,
	}
	return &products
}
