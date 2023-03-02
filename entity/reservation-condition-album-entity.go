package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ReservationConditionAlbum represents users table in database
type ReservationConditionAlbum struct {
	IDReservationConditionAlbum string `gorm:"type:varchar(36);primaryKey;not null" json:"id_reservation_condition_album"`
	ReservationConditionID      string `json:"hotel_id,omitempty" gorm:"default:null;type:varchar(36)"`
	Image                       string `gorm:"default:null;type:varchar(40)" json:"image"`
	Base
	ReservationCondition *ReservationCondition `gorm:"foreignkey:ReservationConditionID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"hotel,omitempty"`
}

func (h *ReservationConditionAlbum) BeforeCreate(tx *gorm.DB) (err error) {
	h.IDReservationConditionAlbum = uuid.NewString()
	return
}
