package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReservationService struct {
	IDReservationService string `gorm:"type:varchar(36);primaryKey" json:"id_reservation_service_detail"`
	ReservationDetailID  string `gorm:"type:varchar(36);default:null" json:"reservation_detail_id"`
	//ServiceDetailID            string `gorm:"type:varchar(36);default:null" json:"service_detail_id"`
	ServiceID string `gorm:"type:varchar(36);default:null" json:"service_id"`
	Quantity  uint8  `gorm:"default:null" json:"quantity"`
	Note      string `gorm:"default:null" json:"note"`
	Base
	Service           *Service           `gorm:"foreignKey:ServiceID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"service,omitempty"`
	ReservationDetail *ReservationDetail `gorm:"foreignKey:ReservationDetailID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"reservation-detail,omitempty"`
	//ServiceDetail     ServiceDetail     `gorm:"foreignKey:ServiceDetailID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"service_detail"`
}

func (rsd *ReservationService) BeforeCreate(tx *gorm.DB) (err error) {
	rsd.IDReservationService = uuid.NewString()
	return
}
