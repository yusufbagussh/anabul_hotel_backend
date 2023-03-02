package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReservationCondition struct {
	IDReservationCondition string `gorm:"type:varchar(36);primaryKey" json:"id_reservation_condition"`
	//Status                 string `gorm:"default:'Belum';type:enum('Sudah','Belum')" json:"status"`
	Title               string `gorm:"default:null;type:varchar(40)" json:"title"`
	Description         string `gorm:"default:null;type:text" json:"description"`
	Image               string `gorm:"default:null;type:varchar(40)" json:"image"`
	ReservationDetailID string `gorm:"type:varchar(36);not null" json:"reservation_detail_id"`
	Status              string `gorm:"default:null;type:enum('Makan', 'Bermain')" json:"status"`
	CreatedBy           string `gorm:"type:varchar(36);references:UserID;not null" json:"created_by"`
	UpdatedBy           string `gorm:"type:varchar(36);references:UserID;not null" json:"updated_by"`
	Base
	ReservationDetail *ReservationDetail `gorm:"foreignkey:ReservationDetailID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"reservation_detail,omitempty"`
	CreatedUser       *User              `gorm:"foreignKey:CreatedBy;constraint:onUpdate:CASCADE;onDelete:CASCADE" json:"created_user,omitempty"`
	UpdatedUser       *User              `gorm:"foreignKey:UpdatedBy;constraint:onUpdate:CASCADE;onDelete:CASCADE" json:"updated_user,omitempty"`
}

func (rc *ReservationCondition) BeforeCreate(tx *gorm.DB) (err error) {
	rc.IDReservationCondition = uuid.NewString()
	return
}
