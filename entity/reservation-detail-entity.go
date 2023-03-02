package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReservationDetail struct {
	IDReservationDetail string `gorm:"type:varchar(36);primaryKey" json:"id_reservation_detail"`
	ReservationID       string `gorm:"type:varchar(36);not null" json:"reservation_id"`
	PetID               string `gorm:"type:varchar(36);not null" json:"pet_id"`
	CageID              string `gorm:"type:varchar(36);not null" json:"cage_id"`
	//ProductID           string  `gorm:"type:varchar(36);default:null" json:"product_id"`
	Weight         float64 `gorm:"not null" json:"weight"`
	CageDetailID   string  `gorm:"type:varchar(36);default:null" json:"cage_category_type_id"`
	CageCategoryID string  `gorm:"type:varchar(36);default:null" json:"cage_category_id"`
	CageTypeID     string  `gorm:"type:varchar(36);default:null" json:"cage_type_id"`
	Vaccination    string  `gorm:"default:'Belum';type:enum('Sudah', 'Belum')" json:"vaccination"`
	FoodAllergy    string  `gorm:"default:null" json:"food_allergy"`
	FleaDisease    string  `gorm:"default:'Tidak';type:enum('Ya', 'Tidak')" json:"flea_disease"`
	AnotherDisease string  `gorm:"default:null" json:"another_disease"`
	StaffID        string  `gorm:"type:varchar(36);references:UserID;default:null" json:"staff_id"`
	Base
	Reservation *Reservation `gorm:"foreignkey:ReservationID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"reservation,omitempty"`
	Pet         *Pet         `gorm:"foreignkey:PetID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"pet,omitempty"`
	Cage        *Cage        `gorm:"foreignkey:CageID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"cage,omitempty"`
	//Product                *Product                `gorm:"foreignkey:ProductID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"product,omitempty"`
	CageDetail             *CageDetail             `gorm:"foreignkey:CageDetailID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"cage_category_type,omitempty"`
	CageCategory           *CageCategory           `gorm:"foreignkey:CageCategoryID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"cage_category,omitempty"`
	CageType               *CageType               `gorm:"foreignkey:CageTypeID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"cage_type,omitempty"`
	Staff                  *User                   `gorm:"foreignKey:StaffID;constraint:onUpdate:CASCADE;onDelete:CASCADE" json:"staff,omitempty"`
	ReservationServices    *[]ReservationService   `json:"reservation_services,omitempty"`
	ReservationInventories *[]ReservationInventory `json:"reservationInventories,omitempty"`
	ReservationProducts    *[]ReservationProduct   `json:"reservation_products,omitempty"`
}

func (r *ReservationDetail) BeforeCreate(tx *gorm.DB) (err error) {
	r.IDReservationDetail = uuid.NewString()
	return
}
