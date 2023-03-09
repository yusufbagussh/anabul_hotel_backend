package seeder

import (
	"github.com/yusufbagussh/pet_hotel_backend/entity"
)

func ReservationProductSeeder(reserDetails *[]entity.ReservationDetail, service *[]entity.Product) *[]entity.ReservationProduct {
	reservationProductDetail1 := entity.ReservationProduct{
		ReservationDetailID: (*reserDetails)[0].IDReservationDetail,
		ProductID:           (*service)[0].IDProduct,
		Quantity:            2,
	}
	reservationProductDetail2 := entity.ReservationProduct{
		ReservationDetailID: (*reserDetails)[1].IDReservationDetail,
		ProductID:           (*service)[0].IDProduct,
		Quantity:            2,
	}
	reservationProductDetails := []entity.ReservationProduct{
		reservationProductDetail1,
		reservationProductDetail2,
	}
	return &reservationProductDetails
}
