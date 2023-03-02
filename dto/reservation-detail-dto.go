package dto

type CreateReservationDetail struct {
	ReservationID string `json:"reservation_id" form:"reservation_id"`
	PetID         string `json:"pet_id" form:"pet_id" binding:"required"`
	SpeciesID     string `json:"species_id" form:"species_id" binding:"required"`
	CageID        string `json:"cage_id" form:"cage_id"`
	//ProductID            string                       `json:"product_id" form:"product_id" binding:"required"`
	Weight               float64                      `json:"weight" form:"weight" binding:"required"`
	CageDetailID         string                       `json:"cage_detail_id" form:"cage_detail_id" binding:"required"`
	CageCategoryID       string                       `json:"cage_category_id" form:"cage_category_id" binding:"required"`
	CageTypeID           string                       `json:"cage_type_id" form:"cage_type_id" binding:"required"`
	Vaccination          string                       `json:"vaccination" form:"vaccination"`
	FoodAllergy          string                       `json:"food_allergy" form:"food_allergy"`
	FleaDisease          string                       `json:"flea_disease" form:"flea_disease"`
	AnotherDisease       string                       `json:"another_disease" form:"another_disease"`
	ReservationService   []CreateReservationService   `json:"reservation_service"`
	ReservationInventory []CreateReservationInventory `json:"reservation_inventory"`
	ReservationProduct   []CreateReservationProduct   `json:"reservation_product"`
}

type UpdateReservationDetail struct {
	IDReservationDetail string `json:"id_reservation_detail" form:"id_reservation_detail"`
	ReservationID       string `json:"reservation_id" form:"reservation_id"`
	PetID               string `json:"pet_id" form:"pet_id" binding:"required"`
	SpeciesID           string `json:"species_id" form:"species_id" binding:"required"`
	CageID              string `json:"cage_id" form:"cage_id"`
	//ProductID            string                       `json:"product_id" form:"product_id" binding:"required"`
	Weight               float64                      `json:"weight" form:"weight" binding:"required"`
	CageDetailID         string                       `json:"cage_detail_id" form:"cage_detail_id" binding:"required"`
	CageCategoryID       string                       `json:"cage_category_id" form:"cage_category_id" binding:"required"`
	CageTypeID           string                       `json:"cage_type_id" form:"cage_type_id" binding:"required"`
	Vaccination          string                       `json:"vaccination" form:"vaccination"`
	FoodAllergy          string                       `json:"food_allergy" form:"food_allergy"`
	FleaDisease          string                       `json:"flea_disease" form:"flea_disease"`
	AnotherDisease       string                       `json:"another_disease" form:"another_disease"`
	ReservationService   []UpdateReservationService   `json:"reservation_service"`
	ReservationInventory []UpdateReservationInventory `json:"reservation_inventory"`
	ReservationProduct   []UpdateReservationProduct   `json:"reservation_product"`
}
