package dto

import "github.com/yusufbagussh/pet_hotel_backend/entity"

type TransactionDetail struct {
	ReservationDetail      entity.ReservationDetail
	ReservationServices    []entity.ReservationService   `json:"reservation_services"`
	ReservationInventories []entity.ReservationInventory `json:"reservation_inventories"`
	ReservationProducts    []entity.ReservationProduct   `json:"reservation_products"`
}
