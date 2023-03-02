package dto

type CreateRate struct {
	Star                uint8  `json:"start" form:"star" binding:"required"`
	Comment             string `json:"comment" form:"comment"`
	ReservationDetailID string `json:"reservation_detail_id" form:"reservation_detail_id" binding:"required"`
}
type UpdateRate struct {
	IDRate              string `json:"id_rate" form:"id_rate" binding:"required"`
	Star                uint8  `json:"start" form:"star" binding:"required"`
	Comment             string `json:"comment" form:"comment"`
	ReservationDetailID string `json:"reservation_detail_id" form:"reservation_detail_id" binding:"required"`
}
