package seeder

import (
	"github.com/yusufbagussh/pet_hotel_backend/entity"
)

func ReservationInventorySeeder(reserDetails *[]entity.ReservationDetail) *[]entity.ReservationInventory {
	reservationInventoryDetail1 := entity.ReservationInventory{
		ReservationDetailID: (*reserDetails)[0].IDReservationDetail,
		Name:                "bola bekel",
		Quantity:            1,
	}
	reservationInventoryDetail2 := entity.ReservationInventory{
		ReservationDetailID: (*reserDetails)[1].IDReservationDetail,
		Name:                "bola bulu",
		Quantity:            1,
	}
	reservationInventoryDetails := []entity.ReservationInventory{
		reservationInventoryDetail1,
		reservationInventoryDetail2,
	}
	return &reservationInventoryDetails
}
