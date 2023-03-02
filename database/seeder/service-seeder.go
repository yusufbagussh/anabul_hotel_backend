package seeder

import (
	"github.com/yusufbagussh/pet_hotel_backend/entity"
)

func ServiceSeeder(hotels *[]entity.Hotel) *[]entity.Service {
	service1 := entity.Service{
		HotelID: (*hotels)[0].IDHotel,
		Name:    "Penitipan",
	}
	service2 := entity.Service{
		HotelID: (*hotels)[0].IDHotel,
		Name:    "Grooming",
	}
	services := []entity.Service{
		service1,
		service2,
	}
	return &services
}
