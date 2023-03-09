package seeder

import (
	"github.com/yusufbagussh/pet_hotel_backend/entity"
)

func GroupDetailSeeder(groups *[]entity.Group, species *[]entity.Species, hotels *[]entity.Hotel) *[]entity.GroupDetail {
	groupDetail1 := entity.GroupDetail{
		GroupID:   (*groups)[0].IDGroup,
		SpeciesID: (*species)[0].IDSpecies,
		MaxWeight: 10,
		MinWeight: 1,
		HotelID:   (*hotels)[0].IDHotel,
	}
	groupDetail2 := entity.GroupDetail{
		GroupID:   (*groups)[1].IDGroup,
		SpeciesID: (*species)[1].IDSpecies,
		MaxWeight: 11,
		MinWeight: 20,
		HotelID:   (*hotels)[0].IDHotel,
	}
	groupDetails := []entity.GroupDetail{
		groupDetail1,
		groupDetail2,
	}
	return &groupDetails
}
