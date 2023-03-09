package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	IDProduct  string  `gorm:"primary_key;type:varchar(36)" json:"id_product"`
	Name       string  `gorm:"type:varchar(100);not null" json:"name"`
	Price      float64 `gorm:"not null" json:"price"`
	UnitType   string  `gorm:"default:null" json:"unit_type"`
	Status     string  `gorm:"default:null;type:enum('Tersedia', 'Kosong')" json:"reservation_status"`
	CategoryID string  `gorm:"type:varchar(36);default:null" json:"category_id"`
	HotelID    string  `gorm:"type:varchar(36);not null" json:"hotel_id"`
	Base
	Hotel    *Hotel    `gorm:"foreignkey:HotelID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"hotel,omitempty"`
	Category *Category `gorm:"foreignkey:CategoryID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"category,omitempty"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.IDProduct = uuid.NewString()
	return
}
