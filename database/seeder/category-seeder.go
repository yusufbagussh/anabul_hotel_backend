package seeder

import (
	"github.com/yusufbagussh/pet_hotel_backend/entity"
)

func CategorySeeder(class *[]entity.Class) *[]entity.Category {
	category1 := entity.Category{
		ClassID: (*class)[0].IDClass,
		Name:    "Anjing",
	}
	category2 := entity.Category{
		ClassID: (*class)[0].IDClass,
		Name:    "Kucing",
	}
	category3 := entity.Category{
		ClassID: (*class)[1].IDClass,
		Name:    "Ayam",
	}
	category4 := entity.Category{
		ClassID: (*class)[1].IDClass,
		Name:    "Burung",
	}
	categories := []entity.Category{
		category1,
		category2,
		category3,
		category4,
	}
	return &categories
}
