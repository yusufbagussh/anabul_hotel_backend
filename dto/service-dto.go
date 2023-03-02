package dto

type CreateService struct {
	Name        string `json:"name" form:"name" binding:"required"`
	Description string `json:"description" form:"description"`
	Unit        string `json:"unit" form:"unit"`
	HotelID     string `json:"hotel_id" form:"hotel_id" binding:"required"`
}

type UpdateService struct {
	IDService   string `json:"id_service" form:"id_service" binding:"required"`
	Name        string `json:"name" form:"name" binding:"required"`
	Description string `json:"description" form:"description"`
	Unit        string `json:"unit" form:"unit"`
	HotelID     string `json:"hotel_id" form:"hotel_id" binding:"required"`
}
