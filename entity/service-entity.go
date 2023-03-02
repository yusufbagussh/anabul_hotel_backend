package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service struct {
	IDService   string `gorm:"primary_key;type:varchar(36)" json:"id_service"`
	Name        string `gorm:"type:varchar(100);not null" json:"name"`
	Description string `gorm:"default:null" json:"description"`
	UnitType    string `gorm:"default:null" json:"unit_type"`
	HotelID     string `json:"hotel_id" gorm:"type:varchar(36);default:null"`
	Base
	Hotel *Hotel `gorm:"foreignkey:HotelID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"hotel,omitempty"`
}

func (s *Service) BeforeCreate(tx *gorm.DB) (err error) {
	s.IDService = uuid.NewString()
	return
}
