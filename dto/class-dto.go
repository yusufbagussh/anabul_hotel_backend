package dto

type CreateClass struct {
	Name string `json:"name" form:"name" binding:"required"`
}
type UpdateClass struct {
	IDClass string `json:"id_class" form:"id_class" binding:"required"`
	Name    string `json:"name" form:"name" binding:"required"`
}
