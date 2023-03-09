package dto

import "mime/multipart"

// UserHotelCreateDTO is a model that client use when create a new book
type UserHotelCreateDTO struct {
	Name      string                `json:"name" form:"name" binding:"required"`
	Email     string                `json:"email" form:"email" binding:"required"`
	Password  string                `json:"password,omitempty" form:"password,omitempty"`
	NIK       uint64                `json:"nik" form:"nik" binding:"required"`
	KTP       *multipart.FileHeader `json:"ktp" form:"ktp"`
	Selfie    *multipart.FileHeader `json:"selfie" form:"selfie"`
	Role      string                `json:"role" form:"role"`
	BirthDate string                `json:"birth_date" form:"birth_date"`
	Address   string                `json:"address" form:"address"`
	Phone     uint64                `json:"phone" form:"phone"`
	HotelID   string                `json:"hotel_id" form:"hotel_id"`
	Image     *multipart.FileHeader `json:"image" form:"image"`
	Gender    string                `json:"gender" form:"gender"`
}

type UserHotelUpdateDTO struct {
	ID        string                `json:"id" form:"id" binding:"required"`
	Name      string                `json:"name" form:"name" binding:"required"`
	Email     string                `json:"email" form:"email" binding:"required"`
	BirthDate string                `json:"birth_date" form:"birth_date"`
	Gender    string                `json:"gender" form:"gender"`
	NIK       uint64                `json:"nik" form:"nik"`
	KTP       *multipart.FileHeader `json:"ktp" form:"ktp"`
	Selfie    *multipart.FileHeader `json:"selfie" form:"selfie"`
	Address   string                `json:"address" form:"address"`
	Phone     uint64                `json:"phone" form:"phone"`
	Role      string                `json:"role" form:"role"`
	Image     *multipart.FileHeader `json:"image,omitempty" form:"image,omitempty"`
	HotelID   string                `json:"hotel_id" form:"hotel_id"`
}
