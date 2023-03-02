package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReservationInventory struct {
	IDReservationInventory string `gorm:"type:varchar(36);primaryKey" json:"id_inventory"`
	Name                   string `gorm:"type:varchar(100)" json:"name"`
	Quantity               uint8  `gorm:"default:null" json:"quantity"`
	ReservationDetailID    string `gorm:"type:varchar(36);not null" json:"reservation_detail_id"`
	Base
	ReservationDetail *ReservationDetail `gorm:"foreignkey:ReservationDetailID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"reservation_detail,omitempty"`
}

func (ri *ReservationInventory) BeforeCreate(tx *gorm.DB) (err error) {
	ri.IDReservationInventory = uuid.NewString()
	return
}
