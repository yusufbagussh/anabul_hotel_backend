package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Pet struct {
	IDPet        string    `gorm:"type:varchar(36);primaryKey" json:"id_pet"`
	Name         string    `gorm:"type:varchar(100)" json:"name"`
	UserID       string    `gorm:"type:varchar(36);not null" json:"user_id"`
	SpeciesID    string    `gorm:"type:varchar(36);not null" json:"species_id"`
	BirthDate    time.Time `gorm:"default:null" json:"birth_date"`
	Gender       string    `gorm:"default:null" sql:"type:ENUM('Laki-laki', 'Perempuan')" json:"gender"`
	FavoriteFood string    `gorm:"default:null" json:"favorite_food"`
	Image        string    `gorm:"default:null;type:varchar(255)" json:"image"`
	Base
	User    *User    `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user,omitempty"`
	Species *Species `gorm:"foreignkey:SpeciesID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"species,omitempty"`
}

func (p *Pet) BeforeCreate(tx *gorm.DB) (err error) {
	p.IDPet = uuid.NewString()
	return
}
