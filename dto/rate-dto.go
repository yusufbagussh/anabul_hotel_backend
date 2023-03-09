package dto

type CreateRate struct {
	Star          uint8  `json:"start" form:"star" binding:"required"`
	Comment       string `json:"comment" form:"comment"`
	ReservationID string `json:"reservation_id" form:"reservation_id" binding:"required"`
	UserID        string `json:"user_id" form:"user_id"`
}
type UpdateRate struct {
	IDRate        string `json:"id_rate" form:"id_rate" binding:"required"`
	Star          uint8  `json:"start" form:"star" binding:"required"`
	Comment       string `json:"comment" form:"comment"`
	ReservationID string `json:"reservation_id" form:"reservation_id" binding:"required"`
	UserID        string `json:"user_id" form:"user_id"`
}
