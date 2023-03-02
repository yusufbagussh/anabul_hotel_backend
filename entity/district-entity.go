package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type District struct {
	IDDistrict string `gorm:"type:varchar(36);primaryKey;not null" json:"id_district"`
	Name       string `gorm:"type:varchar(100)" json:"name"`
	CityID     string `json:"city_id" gorm:"type:varchar(36);not null"`
	Base
	City *City `gorm:"foreignkey:CityID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"city,omitempty"`
}

func (d *District) BeforeCreate(tx *gorm.DB) (err error) {
	d.IDDistrict = uuid.NewString()
	return
}
