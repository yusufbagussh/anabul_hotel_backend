package entity

import (
	"gorm.io/gorm"
	"time"
)

type Base struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
