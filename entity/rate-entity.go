package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Rate struct {
	IDRate        string `gorm:"type:varchar(36);primaryKey" json:"id_rate"`
	Star          uint8  `gorm:"type:varchar(100)" json:"start"`
	Comment       string `gorm:"default:null" json:"comment"`
	ReservationID string `gorm:"type:varchar(36);not null" json:"reservation_id"`
	UserID        string `gorm:"type:varchar(36);default:null" json:"user_id"`
	Base
	Reservation *Reservation `gorm:"foreignkey:ReservationID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"reservation,omitempty"`
	User        *User        `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user,omitempty"`
}

func (r *Rate) BeforeCreate(tx *gorm.DB) (err error) {
	r.IDRate = uuid.NewString()
	return
}
