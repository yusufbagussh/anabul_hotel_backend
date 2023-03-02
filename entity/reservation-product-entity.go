package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReservationProduct struct {
	IDReservationProduct string `gorm:"type:varchar(36);primaryKey" json:"id_reservation_product"`
	ReservationDetailID  string `gorm:"type:varchar(36);not null" json:"reservation_detail_id"`
	ProductID            string `gorm:"type:varchar(36)" json:"product_id"`
	Quantity             uint8  `gorm:"default:null" json:"quantity"`
	Note                 string `gorm:"default:null" json:"note"`
	Base
	Product           *Product           `gorm:"foreignkey:ProductID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"product,omitempty"`
	ReservationDetail *ReservationDetail `gorm:"foreignkey:ReservationDetailID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"reservation_detail,omitempty"`
}

func (ri *ReservationProduct) BeforeCreate(tx *gorm.DB) (err error) {
	ri.IDReservationProduct = uuid.NewString()
	return
}
