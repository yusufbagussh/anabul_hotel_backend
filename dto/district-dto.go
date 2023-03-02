package dto

type CreateDistrict struct {
	Name   string `json:"name" binding:"required" form:"name"`
	CityID string `json:"city_id" binding:"required" form:"city_id"`
}

type UpdateDistrict struct {
	IDDistrict string `json:"id_district" binding:"required" form:"id_district"`
	Name       string `json:"name" binding:"required" form:"name"`
	CityID     string `json:"city_id" binding:"required" form:"city_id"`
}
