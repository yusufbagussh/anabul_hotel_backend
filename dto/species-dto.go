package dto

type CreateSpecies struct {
	Name       string `json:"name" form:"name" binding:"required"`
	CategoryID string `json:"category_id" form:"category_id" binding:"required"`
}
type UpdateSpecies struct {
	IDSpecies  string `json:"id_species" form:"id_species" binding:"required"`
	Name       string `json:"name" form:"name" binding:"required"`
	CategoryID string `json:"category_id" form:"category_id" binding:"required"`
}
