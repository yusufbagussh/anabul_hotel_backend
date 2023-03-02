package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Group struct {
	IDGroup string `gorm:"primary_key;not null;type:varchar(36)" json:"id_category"`
	Name    string `gorm:"type:varchar(100)" json:"name"`
	HotelID string `json:"hotel_id" gorm:"type:varchar(36);not null"`
	Base
	Hotel *Hotel `gorm:"foreignkey:HotelID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"hotel"`
}

func (group *Group) BeforeCreate(tx *gorm.DB) (err error) {
	group.IDGroup = uuid.NewString()
	return
}
