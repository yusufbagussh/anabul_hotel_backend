package seeder

import (
	"github.com/yusufbagussh/pet_hotel_backend/entity"
)

func ResponseSeeder(rates *[]entity.Rate) *[]entity.Response {
	response1 := entity.Response{
		RateID:  (*rates)[0].IDRate,
		Comment: "Terima kasih kak",
	}
	//response2 := entity.Response{
	//	RateID:  (*rates)[1].IDRate,
	//	Comment: "Terima kasih kak",
	//}
	responses := []entity.Response{
		response1,
		//response2,
	}
	return &responses
}
