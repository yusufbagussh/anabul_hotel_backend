package dto

type CreateReservationNotification struct {
	Title         string `json:"title" form:"title"`
	Description   string `json:"description" form:"description"`
	ReservationID string `json:"reservation_id" form:"reservation_id"`
}

type UpdateReservationNotification struct {
	IDReservationNotification string `json:"id_reservation_notification" form:"id_reservation_notification"`
	Title                     string `json:"title" form:"title"`
	Description               string `json:"description" form:"description"`
	ReservationID             string `json:"reservation_id" form:"reservation_id"`
}
