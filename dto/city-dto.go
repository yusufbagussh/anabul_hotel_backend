package dto

type CreateCity struct {
	Name       string `json:"name" form:"name" binding:"required"`
	ProvinceID string `json:"province_id" form:"province_id" binding:"required"`
}
type UpdateCity struct {
	IDCity     string `json:"id_city" form:"id_city" binding:"required"`
	Name       string `json:"name" form:"name" binding:"required"`
	ProvinceID string `json:"province_id" form:"province_id" binding:"required"`
}
