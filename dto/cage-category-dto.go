package dto

type CreateCageCategory struct {
	Name        string `json:"name" binding:"required" form:"name"`
	Description string `json:"description" form:"description"`
	HotelID     string `json:"hotel_id" binding:"required" form:"hotel_id"`
}
type UpdateCageCategory struct {
	IDCageCategory string `json:"id_cage_category" binding:"required" form:"id_cage_category"`
	Name           string `json:"name" binding:"required" form:"name"`
	Description    string `json:"description" form:"description"`
	HotelID        string `json:"hotel_id" binding:"required" form:"hotel_id"`
}
