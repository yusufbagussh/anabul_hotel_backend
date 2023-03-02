package seeder

import (
	"github.com/yusufbagussh/pet_hotel_backend/entity"
)

func ProductSeeder(hotels *[]entity.Hotel) *[]entity.Product {
	product1 := entity.Product{
		HotelID: (*hotels)[0].IDHotel,
		Name:    "Whiskas",
		Price:   33000,
	}
	product2 := entity.Product{
		HotelID: (*hotels)[0].IDHotel,
		Name:    "Bolt",
		Price:   33000,
	}
	products := []entity.Product{
		product1,
		product2,
	}
	return &products
}
