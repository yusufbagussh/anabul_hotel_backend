package seeder

import (
	"github.com/yusufbagussh/pet_hotel_backend/entity"
)

func GroupSeeder(hotels *[]entity.Hotel) *[]entity.Group {
	group1 := entity.Group{
		HotelID: (*hotels)[0].IDHotel,
		Name:    "Small",
	}
	group2 := entity.Group{
		HotelID: (*hotels)[0].IDHotel,
		Name:    "Medium",
	}
	groups := []entity.Group{
		group1,
		group2,
	}
	return &groups
}
