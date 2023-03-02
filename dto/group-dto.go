package dto

type CreateGroup struct {
	Name    string `json:"name" binding:"required" form:"name"`
	HotelID string `json:"hotel_id" binding:"required" form:"hotel_id"`
}
type UpdateGroup struct {
	IDGroup string `json:"id_category" binding:"required" form:"group_id"`
	Name    string `json:"name" binding:"required" form:"name"`
	HotelID string `json:"hotel_id" binding:"required" form:"hotel_id"`
}
