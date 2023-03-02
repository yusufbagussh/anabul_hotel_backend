package dto

type CreateCageType struct {
	Name    string  `json:"name" form:"name" binding:"required"`
	Length  float32 `json:"length" form:"length" binding:"required"`
	Width   float32 `json:"width" form:"width" binding:"required"`
	Height  float32 `json:"height" form:"height" binding:"required"`
	HotelID string  `json:"hotel_id" form:"hotel_id" binding:"required"`
}
type UpdateCageType struct {
	IDCageType string  `json:"id_cage_type" form:"id_cage_type" binding:"required"`
	Name       string  `json:"name" form:"name" binding:"required"`
	Length     float32 `json:"length" form:"length" binding:"required"`
	Width      float32 `json:"width" form:"width" binding:"required"`
	Height     float32 `json:"height" form:"height" binding:"required"`
	HotelID    string  `json:"hotel_id" form:"hotel_id" binding:"required"`
}
