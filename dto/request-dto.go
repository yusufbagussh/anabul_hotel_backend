package dto

import "mime/multipart"

// CreateRequest represents users table in database
type CreateRequest struct {
	HotelName  string                `json:"hotel_name" form:"hotel_name"`
	HotelEmail string                `json:"hotel_email" form:"hotel_email"`
	HotelPhone uint64                `json:"hotel_phone" form:"hotel_phone"`
	NPWP       *multipart.FileHeader `json:"npwp" form:"npwp"`         //NPWP
	Document   *multipart.FileHeader `json:"document" form:"document"` //SIUP,AKTA, NIB
	AdminName  string                `json:"admin_name" form:"admin_name"`
	AdminPhone uint64                `json:"phone" form:"admin_phone"`
	NIK        uint64                `json:"nik" form:"nik"`
	KTP        *multipart.FileHeader `json:"ktp" form:"ktp"`
	Selfie     *multipart.FileHeader `json:"selfie" form:"selfie"`
	Status     string                `json:"status" form:"status"`
}
type UpdateRequest struct {
	IDRequest  string                `json:"id_request" form:"id_request" binding:"required"`
	HotelName  string                `json:"hotel_name" form:"hotel_name"`
	HotelEmail string                `json:"hotel_email" form:"hotel_email"`
	HotelPhone uint64                `json:"hotel_phone" form:"hotel_phone"`
	NPWP       *multipart.FileHeader `json:"npwp" form:"npwp"`         //NPWP
	Document   *multipart.FileHeader `json:"document" form:"document"` //SIUP,AKTA, NIB
	AdminName  string                `json:"admin_name" form:"admin_name"`
	AdminPhone uint64                `json:"phone" form:"admin_phone"`
	NIK        uint64                `json:"nik" form:"nik"`
	KTP        *multipart.FileHeader `json:"ktp" form:"ktp"`
	Selfie     *multipart.FileHeader `json:"selfie" form:"selfie"`
	Status     string                `json:"status" form:"status"`
}
