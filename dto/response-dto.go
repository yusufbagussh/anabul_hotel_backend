package dto

type CreateResponse struct {
	RateID  string `json:"rate_id" form:"rate_id" binding:"required"`
	Comment string `json:"comment" form:"comment" binding:"required"`
}
type UpdateResponse struct {
	IDResponse string `json:"id_response" form:"id_response" binding:"required"`
	RateID     string `json:"rate_id" form:"rate_id" binding:"required"`
	Comment    string `json:"comment" form:"comment" binding:"required"`
}
