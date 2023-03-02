package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PetDetail struct {
	IDPetDetail    string `gorm:"type:varchar(36);primaryKey" json:"id_pet_detail"`
	PetID          string `gorm:"type:varchar(36);not null" json:"pet_id"`
	Vaccination    string `gorm:"default:null" sql:"type:ENUM('Pernah','Belum Pernah')" json:"vaccination"`
	FoodAllergy    string `gorm:"default:null" json:"food_allergy"`
	FleaDisease    string `gorm:"default:null" sql:"type:ENUM('Ya','Tidak')" json:"flea_disease"`
	AnotherDisease string `gorm:"default:null" json:"another_disease"`
	Base
	Pet *Pet `gorm:"foreignkey:PetID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user,omitempty"`
}

func (p *PetDetail) BeforeCreate(tx *gorm.DB) (err error) {
	p.IDPetDetail = uuid.NewString()
	return
}
