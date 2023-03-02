package dto

type CreateSpecies struct {
	Name       string `json:"name" form:"name" binding:"name"`
	CategoryID string `json:"category_id" form:"category_id" binding:"category_id"`
	HotelID    string `json:"hotel_id" form:"hotel_id" binding:"hotel_id"`
}
type UpdateSpecies struct {
	IDSpecies  string `json:"id_species" form:"id_species" binding:"id_species"`
	Name       string `json:"name" form:"name" binding:"name"`
	CategoryID string `json:"category_id" form:"category_id" binding:"category_id"`
	HotelID    string `json:"hotel_id" form:"hotel_id" binding:"hotel_id"`
}
