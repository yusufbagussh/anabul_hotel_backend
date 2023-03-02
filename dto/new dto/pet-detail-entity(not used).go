package entity

type CreatePetDetail struct {
	PetID          string `gorm:"type:varchar(36);not null" json:"pet_id" form:"pet_id" binding:"required"`
	Vaccination    string `json:"vaccination" form:"vaccination"`
	FoodAllergy    string `json:"food_allergy" form:"food_allergy"`
	FleaDisease    string `json:"flea_disease" form:"flea_disease"`
	AnotherDisease string `json:"another_disease" form:"another_disease"`
}
type UpdatePetDetail struct {
	IDPetDetail    string `json:"id_pet_detail" form:"id_pet_detail" binding:"required"`
	PetID          string `gorm:"type:varchar(36);not null" json:"pet_id" form:"pet_id" binding:"required"`
	Vaccination    string `json:"vaccination" form:"vaccination"`
	FoodAllergy    string `json:"food_allergy" form:"food_allergy"`
	FleaDisease    string `json:"flea_disease" form:"flea_disease"`
	AnotherDisease string `json:"another_disease" form:"another_disease"`
}
