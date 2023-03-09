package dto

import (
	"mime/multipart"
)

type CreatePet struct {
	Name         string                `json:"name" form:"name" binding:"required"`
	UserID       string                `json:"user_id" form:"user_id" binding:"required"`
	SpeciesID    string                `json:"species_id" form:"species_id" binding:"required"`
	BirthDate    string                `json:"birth_date" form:"birth_date"`
	Gender       string                `json:"gender" form:"gender"`
	FavoriteFood string                `json:"favorite_food" form:"favorite_food"`
	Image        *multipart.FileHeader `json:"image" form:"image"`
}

type UpdatePet struct {
	IDPet        string                `json:"id_pet" form:"id_pet" binding:"required"`
	Name         string                `json:"name" form:"name" binding:"required"`
	UserID       string                `json:"user_id" form:"user_id" binding:"required"`
	SpeciesID    string                `json:"species_id" form:"species_id" binding:"required"`
	BirthDate    string                `json:"birth_date" form:"birth_date"`
	Gender       string                `json:"gender" form:"gender"`
	FavoriteFood string                `json:"favorite_food" form:"favorite_food"`
	Image        *multipart.FileHeader `json:"image" form:"image"`
}
