package seeder

import (
	"github.com/yusufbagussh/pet_hotel_backend/entity"
)

func ClassSeeder() *[]entity.Class {
	class1 := entity.Class{
		Name: "Mamalia",
	}
	class2 := entity.Class{
		Name: "Aves",
	}
	classes := []entity.Class{
		class1,
		class2,
	}
	return &classes
}
