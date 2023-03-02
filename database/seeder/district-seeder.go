package seeder

import (
	"github.com/yusufbagussh/pet_hotel_backend/entity"
)

func DistrictSeeder(city *[]entity.City) *[]entity.District {
	district1 := entity.District{
		CityID: (*city)[0].IDCity,
		Name:   "Jebres",
	}
	districts := []entity.District{
		district1,
	}
	return &districts
}
