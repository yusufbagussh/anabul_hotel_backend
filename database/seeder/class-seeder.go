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
	class3 := entity.Class{
		Name: "Reptil",
	}
	classes := []entity.Class{
		class1,
		class2,
		class3,
	}
	return &classes
}
