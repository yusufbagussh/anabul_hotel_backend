package dto

type CreateProduct struct {
	Name       string  `json:"name" form:"name" binding:"required"`
	Price      float64 `json:"price" form:"price" binding:"required"`
	Status     string  `json:"status" form:"status"`
	UnitType   string  `json:"unit_type" form:"unit_type"`
	HotelID    string  `json:"hotel_id" form:"hotel_id" binding:"required"`
	CategoryID string  `json:"category_id" form:"category_id" binding:"required"`
}
type UpdateProduct struct {
	IDProduct  string  `json:"id_product" form:"id_product" binding:"required"`
	Name       string  `json:"name" form:"name" binding:"required"`
	Price      float64 `json:"price" form:"price" binding:"required"`
	Status     string  `json:"status" form:"status"`
	UnitType   string  `json:"unit_type" form:"unit_type"`
	HotelID    string  `json:"hotel_id" form:"hotel_id" binding:"required"`
	CategoryID string  `json:"category_id" form:"category_id" binding:"required"`
}
