package dto

type CreateProvince struct {
	Name string `json:"name" binding:"required" form:"name"`
}
type UpdateProvince struct {
	IDProvince string `json:"id_province" binding:"required" form:"id_province"`
	Name       string `json:"name" binding:"required" form:"name"`
}
