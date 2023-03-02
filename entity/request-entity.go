package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Request represents users table in database
type Request struct {
	IDRequest  string `gorm:"type:varchar(36);primaryKey;not null" json:"id_request"`
	HotelName  string `gorm:"type:varchar(100);not null" json:"hotel_name"`
	HotelEmail string `gorm:"unique;not null;type:varchar(100)" json:"hotel_email"`
	HotelPhone uint64 `gorm:"default:null" json:"hotel_phone"`
	NPWP       string `gorm:"default:null;type:varchar(255)" json:"npwp"`     //NPWP
	Document   string `gorm:"default:null;type:varchar(255)" json:"document"` //SIUP,AKTA, NIB
	AdminName  string `gorm:"type:varchar(100);not null" json:"admin_name"`
	AdminPhone uint64 `gorm:"default:null" json:"phone"`
	NIK        uint64 `gorm:"unique;default:null" json:"nik"`
	KTP        string `gorm:"default:null;type:varchar(255)" json:"ktp"`
	Selfie     string `gorm:"default:null;type:varchar(255)" json:"selfie"`
	Status     string `gorm:"default:Proses;type:enum('Tolak', 'Terima', 'Proses')" json:"status"`
	Base
}

func (h *Request) BeforeCreate(tx *gorm.DB) (err error) {
	h.IDRequest = uuid.NewString()
	return
}
