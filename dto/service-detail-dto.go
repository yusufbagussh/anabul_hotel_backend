package dto

type CreateServiceDetail struct {
	ServiceID string  `json:"service_id" form:"service_id" binding:"required"`
	GroupID   string  `json:"group_id" form:"group_id"`
	Price     float64 `json:"price" form:"price" binding:"required"`
	HotelID   string  `json:"hotel_id" form:"hotel_id" binding:"required"`
}

type UpdateServiceDetail struct {
	IDServiceDetail string  `json:"id_service_detail" form:"id_service_detail" binding:"required"`
	ServiceID       string  `json:"service_id" form:"service_id" binding:"required"`
	GroupID         string  `json:"group_id" form:"group_id"`
	Price           float64 `json:"price" form:"price" binding:"required"`
	HotelID         string  `json:"hotel_id" form:"hotel_id" binding:"required"`
}
