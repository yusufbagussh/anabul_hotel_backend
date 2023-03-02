package dto

type CreateGroupDetail struct {
	GroupID   string  `json:"group_id" form:"group_id" binding:"required"`
	SpeciesID string  `json:"species_id" form:"species_id" binding:"required"`
	HotelID   string  `json:"hotel_id" form:"hotel_id" binding:"required"`
	MinWeight float64 `json:"min_weight" form:"min_weight"`
	MaxWeight float64 `json:"max_weight" form:"max_weight"`
}

type UpdateGroupDetail struct {
	IDGroupDetail string  `json:"id_category" form:"id_group_detail" binding:"required"`
	GroupID       string  `json:"group_id" form:"group_id" binding:"required"`
	SpeciesID     string  `json:"species_id" form:"species_id" binding:"required"`
	HotelID       string  `json:"hotel_id" form:"hotel_id" binding:"required"`
	MinWeight     float64 `json:"min_weight" form:"min_weight"`
	MaxWeight     float64 `json:"max_weight" form:"max_weight"`
}
