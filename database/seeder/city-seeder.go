package seeder

import (
	"github.com/yusufbagussh/pet_hotel_backend/entity"
)

func CitySeeder(province *[]entity.Province) *[]entity.City {
	city1 := entity.City{
		ProvinceID: (*province)[0].IDProvince,
		Name:       "Surakarta",
	}
	cities := []entity.City{
		city1,
	}
	return &cities
}
