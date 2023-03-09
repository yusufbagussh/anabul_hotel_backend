package seeder

import (
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"time"
)

func ReservationSeeder(groups *[]entity.Hotel, users *[]entity.User) *[]entity.Reservation {
	reservation1 := entity.Reservation{
		HotelID:   (*groups)[0].IDHotel,
		UserID:    (*users)[3].ID,
		StartDate: time.Now(),
		EndDate:   time.Now(),
		TotalCost: 3500000,
		DPCost:    0,
		//PaymentStatus:     "Dibayar",
		//CheckInStatus:     "Keluar",
		//ReservationStatus: "Diterima",
		//CreatedBy:         (*users)[1].ID,
		//UpdatedBy:         (*users)[1].ID,
	}
	//reservation2 := entity.Reservation{
	//	HotelID:   (*groups)[0].IDHotel,
	//	UserID: (*users)[1].ID,
	//	StartDate: time.Now(),
	//	EndDate: time.Now(),
	//	TotalCost: 3500000,
	//	DPCost: 0,
	//	PaymentStatus: "Paid",
	//	CheckInStatus: "Out",
	//	ReservationStatus: "Completed",
	//}
	reservations := []entity.Reservation{
		reservation1,
		//reservation2,
	}
	return &reservations
}
