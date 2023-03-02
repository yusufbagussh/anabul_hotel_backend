package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CageCategory struct {
	IDCageCategory string `gorm:"type:varchar(36);primaryKey;not null" json:"id_cage_category"`
	Name           string `gorm:"type:varchar(100);not null" json:"name"`
	Description    string `gorm:"type:text;default:null" json:"description"`
	HotelID        string `json:"hotel_id" gorm:"type:varchar(36);not null"`
	Base
	Hotel *Hotel `gorm:"foreignkey:HotelID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"hotel,omitempty"`
}

func (cc *CageCategory) BeforeCreate(tx *gorm.DB) (err error) {
	cc.IDCageCategory = uuid.NewString()
	return
}
