package entity

import "mime/multipart"

// CreateReservationConditionAlbum represents users table in database
type CreateReservationConditionAlbum struct {
	ReservationConditionID string                `json:"hotel_id,omitempty" form:"hotel_id"`
	Image                  *multipart.FileHeader `json:"image" form:"image"`
}

type UpdateReservationConditionAlbum struct {
	IDReservationConditionAlbum string                `json:"id_reservation_condition_album" form:"id_reservation_condition_album" binding:"required"`
	ReservationConditionID      string                `json:"hotel_id,omitempty" form:"hotel_id"`
	Image                       *multipart.FileHeader `json:"image" form:"image"`
}
