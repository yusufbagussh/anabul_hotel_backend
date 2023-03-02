package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Province struct {
	IDProvince string `gorm:"type:varchar(36);primaryKey;not null;" json:"id_province"`
	Name       string `gorm:"type:varchar(100)" json:"province_name"`
	Base
}

func (p *Province) BeforeCreate(tx *gorm.DB) (err error) {
	p.IDProvince = uuid.NewString()
	return
}
