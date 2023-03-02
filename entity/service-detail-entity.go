package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ServiceDetail struct {
	IDServiceDetail string  `gorm:"type:varchar(36);primary_key" json:"id_service_detail"`
	ServiceID       string  `gorm:"type:varchar(36);not null" json:"service_id"`
	GroupID         string  `gorm:"type:varchar(36);not null" json:"group_id"`
	HotelID         string  `gorm:"type:varchar(36);not null" json:"hotel_id"`
	Price           float64 `gorm:"not null" json:"price"`
	Base
	Service *Service `gorm:"foreignKey:ServiceID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"service,omitempty"`
	Group   *Group   `gorm:"foreignKey:GroupID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"group,omitempty"`
	Hotel   *Hotel   `gorm:"foreignkey:HotelID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"hotel,omitempty"`
}

func (sd *ServiceDetail) BeforeCreate(tx *gorm.DB) (err error) {
	sd.IDServiceDetail = uuid.NewString()
	return
}
