package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReservationNotification struct {
	IDReservationNotification string `gorm:"type:varchar(36);primaryKey" json:"id_reservation_notification"`
	Title                     string `gorm:"default:null;type:varchar(40)" json:"title"`
	Description               string `gorm:"default:null;type:text" json:"description"`
	ReservationID             string `gorm:"type:varchar(36);not null" json:"reservation_id"`
	UserID                    string `gorm:"type:varchar(36);not null" json:"user_id"`
	Base
	Reservation *Reservation `gorm:"foreignkey:ReservationID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"reservation_detail,omitempty"`
	User        *User        `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"reservation_detail,omitempty"`
}

func (rc *ReservationNotification) BeforeCreate(tx *gorm.DB) (err error) {
	rc.IDReservationNotification = uuid.NewString()
	return
}
