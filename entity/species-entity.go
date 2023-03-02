package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Species struct {
	IDSpecies  string `gorm:"primaryKey;type:varchar(36)" json:"id_species"`
	Name       string `gorm:"type:varchar(100);not null" json:"name"`
	CategoryID string `json:"category_id" gorm:"type:varchar(36);not null"`
	//HotelID    string `json:"hotel_id" gorm:"type:varchar(36);not null"`
	Base
	//Hotel Hotel `gorm:"foreignkey:HotelID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"hotel"`
	Category *Category `gorm:"foreignkey:CategoryID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"category,omitempty"`
}

func (s *Species) BeforeCreate(tx *gorm.DB) (err error) {
	s.IDSpecies = uuid.NewString()
	return
}
