package dto

type CreateCageDetail struct {
	CageCategoryID string  `json:"cage_category_id" form:"cage_category_id"`
	CageTypeID     string  `json:"cage_type_id" form:"cage_type_id"`
	Quantity       uint8   `json:"quantity" form:"quantity" binding:"required"`
	Price          float64 `json:"price" form:"price" binding:"required"`
	Status         string  `json:"status" form:"status"`
	HotelID        string  `json:"hotel_id" form:"hotel_id" binding:"required"`
}
type UpdateCageDetail struct {
	IDCageDetail   string  `json:"id_cage_detail" form:"id_cage_detail" binding:"required"`
	CageCategoryID string  `json:"cage_category_id" form:"cage_category_id"`
	CageTypeID     string  `json:"cage_type_id" form:"cage_type_id"`
	Quantity       uint8   `json:"quantity" form:"quantity" binding:"required"`
	Price          float64 `json:"price" form:"price" binding:"required"`
	Status         string  `json:"status" form:"status"`
	HotelID        string  `json:"hotel_id" form:"hotel_id" binding:"required"`
}

type UpdateCageDetailStatus struct {
	IDCageDetail string `json:"id_cage_detail" form:"id_cage_detail" binding:"required"`
	Status       string `json:"status" form:"status" binding:"required"`
}
