package seeder

import (
	"github.com/yusufbagussh/pet_hotel_backend/entity"
)

func ProvinceSeeder() *[]entity.Province {
	province1 := entity.Province{
		Name: "Jawa Tengah",
	}
	province2 := entity.Province{
		Name: "Jawa Timur",
	}
	provinces := []entity.Province{
		province1,
		province2,
	}
	return &provinces
}
