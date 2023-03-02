package dto

type CreateReservationInventory struct {
	Name                string `json:"name" form:"name" binding:"required"`
	Quantity            uint8  `json:"quantity" form:"quantity" binding:"required"`
	ReservationDetailID string `json:"reservation_detail_id" form:"reservation_detail_id"`
}

type UpdateReservationInventory struct {
	IDReservationInventory string `json:"id_reservation_inventory" form:"id_reservation_inventory"`
	Quantity               uint8  `json:"quantity" form:"quantity" binding:"required"`
	Name                   string `json:"name" form:"name" binding:"required"`
	ReservationDetailID    string `json:"reservation_detail_id" form:"reservation_detail_id"`
}
