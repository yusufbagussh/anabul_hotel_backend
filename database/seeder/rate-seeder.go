package seeder

import (
	"github.com/yusufbagussh/pet_hotel_backend/entity"
)

func RateSeeder(reservDetail *[]entity.Reservation) *[]entity.Rate {
	rate1 := entity.Rate{
		ReservationID: (*reservDetail)[0].IDReservation,
		Star:          4,
		Comment:       "Pelayanannya cukup memuaskan",
	}
	//rate2 := entity.Rate{
	//	ReservationID: (*reservDetail)[1].IDReservation,
	//	Star:                3,
	//	Comment:             "Pelayanannya cukup memuaskan",
	//}
	rates := []entity.Rate{
		rate1,
		//rate2,
	}
	return &rates
}
