package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Hotel represents users table in database
type Hotel struct {
	IDHotel     string  `gorm:"primaryKey;type:varchar(36)" json:"id_hotel"`
	Name        string  `gorm:"type:varchar(100);not null" json:"name"`
	Email       string  `gorm:"unique;not null;type:varchar(100)" json:"email"`
	ProvinceID  string  `json:"province_id" gorm:"type:varchar(36);default:null"`
	CityID      string  `json:"city_id" gorm:"type:varchar(36);default:null"`
	DistrictID  string  `json:"district_id" gorm:"type:varchar(36);default:null"`
	Description string  `gorm:"default:null;type:text" json:"description"`
	Address     string  `gorm:"default:;type:text" json:"address"`
	Phone       uint64  `gorm:"default:null" json:"no_hp"`
	Image       string  `gorm:"default:null;type:varchar(255)" json:"image"`
	NPWP        string  `gorm:"default:null;type:varchar(255)" json:"npwp"`     //NPWP
	Document    string  `gorm:"default:null;type:varchar(255)" json:"document"` //SIUP,AKTA, NIB
	OpenTime    string  `gorm:"default:null;type:varchar(5)" json:"open_time"`
	CloseTime   string  `gorm:"default:null;type:varchar(5)" json:"close_time"`
	Latitude    float64 `gorm:"default:null" json:"latitude"`
	Longitude   float64 `gorm:"default:null" json:"longitude"`
	MapLink     string  `gorm:"default:null;type:varchar(255)" json:"map_link"`
	Requirement string  `gorm:"default:null" json:"requirement"`
	Regulation  string  `gorm:"default:null" json:"regulation"`
	Base
	Province *Province `gorm:"foreignkey:ProvinceID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"province,omitempty"`
	City     *City     `gorm:"foreignkey:CityID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"city,omitempty"`
	District *District `gorm:"foreignkey:DistrictID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"district,omitempty"`
}

func (h *Hotel) BeforeCreate(tx *gorm.DB) (err error) {
	h.IDHotel = uuid.NewString()
	return
}
