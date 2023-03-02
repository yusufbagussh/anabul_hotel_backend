package dto

type CreateCategory struct {
	Name    string `json:"name" form:"name" binding:"required"`
	ClassID string `json:"class_id" form:"class_id" binding:"required"`
}
type UpdateCategory struct {
	IDCategory string `json:"id_category" form:"id_category" binding:"required"`
	Name       string `json:"name" form:"name" binding:"required"`
	ClassID    string `json:"class_id" form:"class_id" binding:"required"`
}
