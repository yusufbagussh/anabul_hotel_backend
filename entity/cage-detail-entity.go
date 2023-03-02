package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CageDetail struct {
	IDCageDetail   string  `gorm:"type:varchar(36);primaryKey;not null" json:"id_cage_category_type"`
	CageCategoryID string  `gorm:"type:varchar(36);default:null" json:"cage_category_id"`
	CageTypeID     string  `gorm:"type:varchar(36);default:null" json:"cage_type_id"`
	HotelID        string  `gorm:"type:varchar(36);not null" json:"hotel_id"`
	Price          float64 `gorm:"not null" json:"price"`
	Quantity       uint8   `gorm:"default:null" json:"quantity"`
	Status         string  `gorm:"default:null;type:enum('Tersedia', 'Penuh')" json:"reservation_status"`
	Base
	CageCategory *CageCategory `gorm:"foreignkey:CageCategoryID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"cage_category,omitempty"`
	CageType     *CageType     `gorm:"foreignkey:CageTypeID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"cage_type,omitempty"`
	Hotel        *Hotel        `gorm:"foreignkey:HotelID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"hotel,omitempty"`
}

func (cct *CageDetail) BeforeCreate(tx *gorm.DB) (err error) {
	cct.IDCageDetail = uuid.NewString()
	return
}
