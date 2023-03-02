package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	IDCategory string `gorm:"type:varchar(36);primaryKey;not null" json:"id_category"`
	Name       string `gorm:"type:varchar(100);not null" json:"name"`
	ClassID    string `gorm:"type:varchar(36);not null" json:"class_id"`
	Base
	Class *Class `gorm:"foreignkey:ClassID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"class,omitempty"`
}

func (c *Category) BeforeCreate(tx *gorm.DB) (err error) {
	c.IDCategory = uuid.NewString()
	return
}
