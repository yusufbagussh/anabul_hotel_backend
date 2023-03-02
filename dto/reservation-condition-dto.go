package dto

import "mime/multipart"

type CreateReservationCondition struct {
	//Status              string                `json:"status" form:"status"`
	Title               string                `json:"title" form:"title"`
	Description         string                `json:"description" form:"description"`
	Image               *multipart.FileHeader `json:"image" form:"image"`
	Category            string                `json:"category" form:"category"`
	ReservationDetailID string                `json:"reservation_detail_id" form:"reservation_detail_id" binding:"required"`
	CreatedBy           string                `json:"created_by" form:"created_by" binding:"required"`
	UpdatedBy           string                `json:"updated_by" form:"updated_by"`
}

type UpdateReservationCondition struct {
	IDReservationCondition string `json:"id_reservation_condition" form:"id_reservation_condition" binding:"required"`
	//Status                 string                `json:"status" form:"status"`
	Title               string                `json:"title" form:"title"`
	Description         string                `json:"description" form:"description"`
	Image               *multipart.FileHeader `json:"image" form:"image"`
	Category            string                `json:"category" form:"category"`
	ReservationDetailID string                `json:"reservation_detail_id" form:"reservation_detail_id" binding:"required"`
	CreatedBy           string                `json:"created_by" form:"created_by" binding:"required"`
	UpdatedBy           string                `json:"updated_by" form:"updated_by" binding:"required"`
}
