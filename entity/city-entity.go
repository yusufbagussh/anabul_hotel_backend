package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type City struct {
	IDCity     string `gorm:"type:varchar(36);primaryKey;not null" json:"id_city"`
	Name       string `gorm:"type:varchar(100)" json:"name"`
	ProvinceID string `json:"province_id" gorm:"type:varchar(36);not null"`
	Base
	Province *Province `gorm:"foreignkey:ProvinceID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"province,omitempty"`
}

func (c *City) BeforeCreate(tx *gorm.DB) (err error) {
	c.IDCity = uuid.NewString()
	return
}
