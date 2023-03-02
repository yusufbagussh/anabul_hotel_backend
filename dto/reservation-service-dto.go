package dto

type CreateReservationService struct {
	ReservationDetailID string `json:"reservation_detail_id" form:"reservation_detail_id"`
	ServiceID           string `json:"service_id" form:"service_id" binding:"required"`
	Quantity            uint8  `json:"quantity" form:"quantity" binding:"required"`
	Note                string `json:"note" form:"note"`
}
type UpdateReservationService struct {
	IDReservationService string `json:"id_reservation_service" form:"id_reservation_service"`
	ReservationDetailID  string `json:"reservation_detail_id" form:"reservation_detail_id" binding:"required"`
	ServiceID            string `json:"service_id" form:"service_id" binding:"required"`
	Quantity             uint8  `json:"quantity" form:"quantity" binding:"required"`
	Note                 string `json:"note" form:"note"`
}
