package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Class struct {
	IDClass string `gorm:"type:varchar(36);primaryKey;not null" json:"id_class"`
	Name    string `gorm:"type:varchar(100);not null" json:"name"`
	Base
}

func (c *Class) BeforeCreate(tx *gorm.DB) (err error) {
	c.IDClass = uuid.NewString()
	return
}
