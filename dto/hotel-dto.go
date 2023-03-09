package dto

import "mime/multipart"

type CreateHotel struct {
	Name               string                `json:"name" form:"name" binding:"required"`
	Email              string                `json:"email" form:"email" binding:"required"`
	ProvinceID         string                `json:"province_id" form:"province_id" binding:"required"`
	CityID             string                `json:"city_id" form:"city_id" binding:"required"`
	DistrictID         string                `json:"district_id" form:"district_id" binding:"required"`
	Address            string                `json:"address" form:"address"`
	Phone              uint64                `json:"no_hp" form:"phone"`
	Image              *multipart.FileHeader `json:"image" form:"image"`
	Document           *multipart.FileHeader `json:"document" form:"document"`
	NPWP               *multipart.FileHeader `json:"npwp" form:"npwp"`
	Latitude           float64               `json:"latitude" form:"latitude"`
	Longitude          float64               `json:"longitude" form:"longitude"`
	MapLink            float64               `json:"map_link" form:"map_link"`
	Requirement        string                `json:"requirement" form:"requirement"`
	Regulation         string                `json:"regulation" form:"regulation"`
	UserHotelCreateDTO UserHotelCreateDTO    `json:"admin_hotel" form:"admin_hotel"`
}

type UpdateHotel struct {
	IDHotel     string                `json:"id_hotel" form:"id_hotel" binding:"required"`
	Name        string                `json:"name" form:"name" binding:"required"`
	Email       string                `json:"email" form:"email" binding:"required"`
	ProvinceID  string                `json:"province_id" form:"province_id" binding:"required"`
	CityID      string                `json:"city_id" form:"city_id" binding:"required"`
	DistrictID  string                `json:"district_id" form:"district_id" binding:"required"`
	Address     string                `json:"address" form:"address"`
	Phone       uint64                `json:"no_hp" form:"phone"`
	Image       *multipart.FileHeader `json:"image" form:"image"`
	NPWP        *multipart.FileHeader `json:"npwp" form:"npwp"`
	Document    *multipart.FileHeader `json:"document" form:"document"`
	Latitude    float64               `json:"latitude" form:"latitude"`
	Longitude   float64               `json:"longitude" form:"longitude"`
	MapLink     string                `json:"map_link" form:"map_link"`
	Requirement string                `json:"requirement" form:"requirement"`
	Description string                `json:"description" form:"description"`
	Regulation  string                `json:"regulation" form:"regulation"`
	OpenTime    string                `json:"open_time" form:"open_time"`
	CloseTime   string                `json:"close_time" form:"close_time"`
}
