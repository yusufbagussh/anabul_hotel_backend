package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// User represents users table in database
type User struct {
	ID                 string    `gorm:"type:varchar(36);primaryKey;not null" json:"id"`
	Name               string    `gorm:"type:varchar(100);not null" json:"name"`
	Email              string    `gorm:"unique;not null;type:varchar(100)" json:"email"`
	Password           string    `gorm:"->;<-;not null" json:"-"`
	HotelID            string    `json:"hotel_id" gorm:"default:null;type:varchar(36)"`
	NIK                uint64    `gorm:"unique;default:null" json:"nik"`
	KTP                string    `gorm:"default:null;type:varchar(255)" json:"ktp"`
	Selfie             string    `gorm:"default:null;type:varchar(255)" json:"selfie"`
	BirthDate          time.Time `gorm:"default:null" json:"birth_date"`
	Gender             string    `gorm:"default:null;type:enum('Laki-laki', 'Perempuan')" json:"gender"`
	Address            string    `gorm:"default:null" json:"address"`
	Phone              uint64    `gorm:"default:null" json:"phone"`
	Image              string    `gorm:"default:null;type:varchar(255)" json:"image"`
	Role               string    `gorm:"default:'Customer';type:enum('Super Admin', 'Admin', 'Staff', 'Customer')" json:"role"`
	VerificationCode   string    `gorm:"default:null" json:"verification_code"`
	Verified           bool      `gorm:"default:false" json:"verified"`
	Status             bool      `gorm:"default:false" json:"status"`
	DeviceToken        string    `gorm:"default:null;type:varchar(255)" json:"device_token"`
	PasswordResetToken string    `gorm:"default:null" json:"password_reset_token"`
	PasswordResetAt    time.Time `gorm:"default:null" json:"password_reset_at"`
	Base
	Token string `gorm:"-" json:"token,omitempty"`
	Hotel Hotel  `gorm:"foreignkey:HotelID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"hotel,omitempty"`
	//Gender  string `gorm:"default:null;type:varchar(15)" json:"gender"`
	//Role               string `gorm:"default:'Customer;type:varchar(15)" json:"role"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewString()
	return
}
