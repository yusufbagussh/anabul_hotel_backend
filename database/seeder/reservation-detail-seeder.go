package seeder

import (
	"github.com/yusufbagussh/pet_hotel_backend/entity"
)

func ReservationDetailSeeder(reservation *[]entity.Reservation, pets *[]entity.Pet, cages *[]entity.Cage, products *[]entity.Product) *[]entity.ReservationDetail {
	reservationDetail1 := entity.ReservationDetail{
		ReservationID: (*reservation)[0].IDReservation,
		PetID:         (*pets)[0].IDPet,
		CageID:        (*cages)[0].IDCage,
		//ProductID:     (*products)[0].IDProduct,
	}
	reservationDetail2 := entity.ReservationDetail{
		ReservationID: (*reservation)[0].IDReservation,
		PetID:         (*pets)[1].IDPet,
		CageID:        (*cages)[1].IDCage,
		//ProductID:     (*products)[0].IDProduct,
	}
	reservationDetails := []entity.ReservationDetail{
		reservationDetail1,
		reservationDetail2,
	}
	return &reservationDetails
}
