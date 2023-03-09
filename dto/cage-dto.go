package dto

type CreateCage struct {
	Name         string `json:"name" form:"name" binding:"required"`
	CageDetailID string `json:"cage_detail_id" form:"cage_detail_id"`
	Status       string `json:"status" form:"status" binding:"required"`
	HotelID      string `json:"hotel_id" form:"hotel_id" binding:"required"`
}
type UpdateCage struct {
	IDCage       string `json:"id_cage" form:"id_cage" binding:"required"`
	Name         string `json:"name" form:"name" binding:"required"`
	CageDetailID string `json:"cage_detail_id" form:"cage_detail_id"`
	Status       string `json:"status" form:"status" binding:"required"`
	HotelID      string `json:"hotel_id" form:"hotel_id" binding:"required"`
}
