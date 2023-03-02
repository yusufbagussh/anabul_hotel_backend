package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CageType struct {
	IDCageType string  `gorm:"type:varchar(36);primaryKey;not null" json:"id_cage_type"`
	Name       string  `gorm:"type:varchar(100);not null" json:"name"`
	Length     float32 `gorm:"not null" json:"length"`
	Width      float32 `gorm:"not null" json:"width"`
	Height     float32 `gorm:"not null" json:"height"`
	HotelID    string  `json:"hotel_id" gorm:"type:varchar(36);not null"`
	Base
	Hotel *Hotel `gorm:"foreignkey:HotelID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"hotel,omitempty"`
}

func (ct *CageType) BeforeCreate(tx *gorm.DB) (err error) {
	ct.IDCageType = uuid.NewString()
	return
}
