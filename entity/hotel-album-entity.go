package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// HotelAlbum represents users table in database
type HotelAlbum struct {
	IDHotelAlbum string `gorm:"type:varchar(36);primaryKey;not null" json:"id_hotel_album"`
	//Name         string `gorm:"type:varchar(100);not null" json:"name"`
	HotelID string `json:"hotel_id,omitempty" gorm:"default:null;type:varchar(36)"`
	Image   string `gorm:"default:null;type:varchar(255)" json:"image"`
	Base
	Hotel *Hotel `gorm:"foreignkey:HotelID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"hotel,omitempty"`
}

func (h *HotelAlbum) BeforeCreate(tx *gorm.DB) (err error) {
	h.IDHotelAlbum = uuid.NewString()
	return
}
