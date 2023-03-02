package dto

type CreateCity struct {
	Name       string `json:"name" form:"" binding:"required"`
	ProvinceID string `json:"province_id" form:"" binding:"required"`
}
type UpdateCity struct {
	IDCity     string `json:"id_city" form:"" binding:"required"`
	Name       string `json:"name" form:"" binding:"required"`
	ProvinceID string `json:"province_id" form:"" binding:"required"`
}
