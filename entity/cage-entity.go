package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Cage struct {
	IDCage       string `gorm:"type:varchar(36);primaryKey;not null" json:"id_cage"`
	Name         string `gorm:"type:varchar(100);not null" json:"name"`
	CageDetailID string `gorm:"type:varchar(36);not null" json:"cage_category_type_id"`
	Status       string `gorm:"default:'kosong';type:enum('kosong','terisi')" json:"status"`
	HotelID      string `json:"hotel_id" gorm:"type:varchar(36);default:null"`
	Base
	Hotel      *Hotel      `gorm:"foreignkey:HotelID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"hotel,omitempty"`
	CageDetail *CageDetail `gorm:"foreignkey:CageDetailID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"cage_detail,omitempty"`
}

func (c *Cage) BeforeCreate(tx *gorm.DB) (err error) {
	c.IDCage = uuid.NewString()
	return
}
