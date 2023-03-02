package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Response struct {
	IDResponse string `gorm:"type:varchar(36);primary_key" json:"id_response"`
	RateID     string `gorm:"type:varchar(36);not null" json:"rate_id"`
	Comment    string `gorm:"default:null" json:"comment"`
	Base
	Rate *Rate `gorm:"foreignKey:RateID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"rate,omitempty"`
}

func (r *Response) BeforeCreate(tx *gorm.DB) (err error) {
	r.IDResponse = uuid.NewString()
	return
}
