package dto

type CreateReservationProduct struct {
	ReservationDetailID string `json:"reservation_detail_id" form:"reservation_detail_id"`
	ProductID           string `json:"product_id" form:"product_id" binding:"required"`
	Quantity            uint8  `json:"quantity" form:"quantity" binding:"required"`
	Note                string `json:"note" form:"note"`
}
type UpdateReservationProduct struct {
	IDReservationProduct string `json:"id_reservation_product" form:"id_reservation_product"`
	ReservationDetailID  string `json:"reservation_detail_id" form:"reservation_detail_id" binding:"required"`
	ProductID            string `json:"product_id" form:"product_id" binding:"required"`
	Quantity             uint8  `json:"quantity" form:"quantity" binding:"required"`
	Note                 string `json:"note" form:"note"`
}
