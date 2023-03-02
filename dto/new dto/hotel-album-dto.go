package entity

import "mime/multipart"

// CreateHotelAlbum represents users table in database
type CreateHotelAlbum struct {
	Name    string                `json:"name" form:"name"`
	HotelID string                `json:"hotel_id,omitempty" form:"hotel_id"`
	Image   *multipart.FileHeader `json:"image" form:"image"`
}
type UpdateHotelAlbum struct {
	IDHotelAlbum string                `json:"id_album_hotel" form:"id_hotel_album"`
	Name         string                `json:"name" form:"name"`
	HotelID      string                `json:"hotel_id,omitempty" form:"hotel_id"`
	Image        *multipart.FileHeader `json:"image" form:"image"`
}
