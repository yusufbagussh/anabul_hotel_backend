package seeder

import (
	"github.com/yusufbagussh/pet_hotel_backend/entity"
)

func ReservationDetailSeeder(reservation *[]entity.Reservation, pets *[]entity.Pet, cages *[]entity.Cage, cageDetails *[]entity.CageDetail, cageCategories *[]entity.CageCategory, cageTypes *[]entity.CageType, users *[]entity.User) *[]entity.ReservationDetail {
	reservationDetail1 := entity.ReservationDetail{
		ReservationID:  (*reservation)[0].IDReservation,
		PetID:          (*pets)[0].IDPet,
		CageID:         (*cages)[0].IDCage,
		StaffID:        (*users)[3].ID,
		CageDetailID:   (*cageDetails)[0].IDCageDetail,
		CageCategoryID: (*cageCategories)[0].IDCageCategory,
		CageTypeID:     (*cageTypes)[0].IDCageType,
		Weight:         5,
		//ProductID:     (*products)[0].IDProduct,
	}
	reservationDetail2 := entity.ReservationDetail{
		ReservationID:  (*reservation)[0].IDReservation,
		PetID:          (*pets)[1].IDPet,
		CageID:         (*cages)[1].IDCage,
		StaffID:        (*users)[3].ID,
		CageDetailID:   (*cageDetails)[1].IDCageDetail,
		CageCategoryID: (*cageCategories)[0].IDCageCategory,
		CageTypeID:     (*cageTypes)[1].IDCageType,
		Weight:         11,
		//ProductID:     (*products)[0].IDProduct,
	}
	reservationDetails := []entity.ReservationDetail{
		reservationDetail1,
		reservationDetail2,
	}
	return &reservationDetails
}
