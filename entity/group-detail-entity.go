package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GroupDetail struct {
	IDGroupDetail string  `gorm:"primaryKey;type:varchar(36);not null" json:"id_category"`
	GroupID       string  `gorm:"type:varchar(36);default:null" json:"group_id"`
	SpeciesID     string  `gorm:"type:varchar(36);default:null" json:"species_id"`
	MinWeight     float64 `gorm:"default:null" json:"min_weight"`
	MaxWeight     float64 `gorm:"default:null" json:"max_weight"`
	HotelID       string  `json:"hotel_id" gorm:"type:varchar(36);not null"`
	Base
	Hotel   *Hotel   `gorm:"foreignkey:HotelID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"hotel"`
	Group   *Group   `gorm:"foreignkey:GroupID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"group"`
	Species *Species `gorm:"foreignkey:SpeciesID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"species"`
}

func (gt *GroupDetail) BeforeCreate(tx *gorm.DB) (err error) {
	gt.IDGroupDetail = uuid.NewString()
	return
}
