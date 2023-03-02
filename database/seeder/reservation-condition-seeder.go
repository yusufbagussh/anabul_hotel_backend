package seeder

import (
	"github.com/yusufbagussh/pet_hotel_backend/entity"
)

func ReservationConditionSeeder(reserDetails *[]entity.ReservationDetail, users *[]entity.User) *[]entity.ReservationCondition {
	reservationCondition1 := entity.ReservationCondition{
		ReservationDetailID: (*reserDetails)[0].IDReservationDetail,
		//Status:              "Sudah",
		Title:       "Pemberian Makan",
		Description: "Peliharaan Anda Sudah Selesai Makan",
		CreatedBy:   (*users)[2].ID,
		UpdatedBy:   (*users)[2].ID,
	}
	reservationCondition2 := entity.ReservationCondition{
		ReservationDetailID: (*reserDetails)[0].IDReservationDetail,
		//Status:              "Sudah",
		Title:       "Kegiatan Bermain",
		Description: "Peliharaan Anda Sudah Selesai Bermain",
		CreatedBy:   (*users)[2].ID,
		UpdatedBy:   (*users)[2].ID,
	}
	reserConditions := []entity.ReservationCondition{
		reservationCondition1,
		reservationCondition2,
	}
	return &reserConditions
}
