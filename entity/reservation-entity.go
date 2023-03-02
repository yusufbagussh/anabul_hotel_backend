package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Reservation struct {
	IDReservation     string `gorm:"type:varchar(36);primary_key" json:"id_reservation"`
	Name              string `gorm:"type:varchar(255);default null" json:"name"`
	HotelID           string `gorm:"type:varchar(36);not null" json:"hotel_id"`
	UserID            string `gorm:"type:varchar(36);not null" json:"user_id"`
	StartDate         time.Time
	EndDate           time.Time
	TotalCost         float64 `gorm:"not null" json:"total_cost"`
	DPCost            float64 `gorm:"default:null" json:"dp_cost"`
	PaymentStatus     string  `gorm:"default:null;type:enum('Proses', 'Dibayar', 'Uang Muka')" json:"payment_status"`
	CheckInStatus     string  `gorm:"default:null;type:enum('Masuk', 'Keluar')" json:"check_in_status"`
	ReservationStatus string  `gorm:"default:null;type:enum('Proses','Diterima', 'Ditolak', 'Dibatalkan')" json:"reservation_status"`
	//CreatedBy         string  `gorm:"type:varchar(36);references:UserID;not null" json:"created_by"`
	//UpdatedBy         string  `gorm:"type:varchar(36);references:UserID;not null" json:"updated_by"`
	Base
	Hotel *Hotel `gorm:"foreignkey:HotelID;constraint:onUpdate:CASCADE;onDelete:CASCADE" json:"hotel,omitempty"`
	User  *User  `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE;onDelete:CASCADE" json:"user,omitempty"`
	//CreatedUser        *User                `gorm:"foreignKey:CreatedBy;constraint:onUpdate:CASCADE;onDelete:CASCADE" json:"created_user,omitempty"`
	//UpdatedUser        *User                `gorm:"foreignKey:UpdatedBy;constraint:onUpdate:CASCADE;onDelete:CASCADE" json:"updated_user,omitempty"`
	ReservationDetails *[]ReservationDetail `json:"reservation_details,omitempty"`
	Rate               *Rate                `json:"rate,omitempty"`
}

func (r *Reservation) BeforeCreate(tx *gorm.DB) (err error) {
	r.IDReservation = uuid.NewString()
	return
}
