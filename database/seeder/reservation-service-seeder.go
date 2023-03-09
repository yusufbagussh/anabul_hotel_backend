package seeder

import (
	"github.com/yusufbagussh/pet_hotel_backend/entity"
)

func ReservationServiceSeeder(reserDetails *[]entity.ReservationDetail, service *[]entity.Service) *[]entity.ReservationService {
	reservationServiceDetail1 := entity.ReservationService{
		ReservationDetailID: (*reserDetails)[0].IDReservationDetail,
		ServiceID:           (*service)[0].IDService,
		Quantity:            1,
	}
	reservationServiceDetail2 := entity.ReservationService{
		ReservationDetailID: (*reserDetails)[1].IDReservationDetail,
		ServiceID:           (*service)[0].IDService,
		Quantity:            1,
	}
	reservationServiceDetails := []entity.ReservationService{
		reservationServiceDetail1,
		reservationServiceDetail2,
	}
	return &reservationServiceDetails
}
