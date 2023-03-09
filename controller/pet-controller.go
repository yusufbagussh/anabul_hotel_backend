package controller

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/helper"
	"github.com/yusufbagussh/pet_hotel_backend/service"
	"net/http"
)

// PetController is a contract of what petController can do
type PetController interface {
	GetPet(ctx *gin.Context)
	CreatePet(ctx *gin.Context)
	DeletePet(ctx *gin.Context)
	ShowPet(ctx *gin.Context)
	UpdatePet(ctx *gin.Context)
}

type petController struct {
	petService service.PetService
	jwtService service.JWTService
}

// NewPetController is creating anew instance of PetControlller
func NewPetController(petService service.PetService, jwtService service.JWTService) PetController {
	return &petController{
		petService: petService,
		jwtService: jwtService,
	}
}

func (u *petController) CreatePet(ctx *gin.Context) {
	var CreatePet dto.CreatePet
	errDTO := ctx.ShouldBind(&CreatePet)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, errCreate := u.petService.CreatePet(CreatePet, ctx)
	if errCreate != nil {
		res := helper.BuildErrorResponse("Failed to create pet", errCreate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Create success", result)
	ctx.JSON(http.StatusOK, res)
}

func (u *petController) GetPet(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, _ := u.jwtService.ValidateToken(authHeader)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	var filterPagination dto.FilterPagination
	errDTO := ctx.ShouldBind(&filterPagination)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var staffs, pagination, errGet = u.petService.GetAllPet(userID, filterPagination)
	if errGet != nil {
		res := helper.BuildErrorResponse("Failed to get all data", errGet.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponsePage(true, "OK", staffs, pagination)
	ctx.JSON(http.StatusOK, res)
}

func (u *petController) UpdatePet(ctx *gin.Context) {
	var UpdatePet dto.UpdatePet
	errDTO := ctx.ShouldBind(&UpdatePet)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	authHeader := ctx.GetHeader("Authorization")
	token, _ := u.jwtService.ValidateToken(authHeader)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if UpdatePet.UserID != userID {
		res := helper.BuildErrorResponse(
			"You don't have permission",
			"You are not the owner",
			helper.EmptyObj{},
		)
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}
	result, errCreate := u.petService.UpdatePet(UpdatePet, ctx)
	if errCreate != nil {
		res := helper.BuildErrorResponse("Failed to update pet", errCreate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Update success", result)
	ctx.JSON(http.StatusOK, res)
}
func (u *petController) DeletePet(ctx *gin.Context) {
	petID := ctx.Param("id")
	authHeader := ctx.GetHeader("Authorization")
	token, _ := u.jwtService.ValidateToken(authHeader)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	errDel, ok := u.petService.DeletePet(petID, userID)
	if ok == false {
		res := helper.BuildErrorResponse(
			"You don't have permission",
			"You are not the owner",
			helper.EmptyObj{},
		)
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}
	if errDel != nil {
		res := helper.BuildErrorResponse("Failed to delete", errDel.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Delete success", nil)
	ctx.JSON(http.StatusOK, res)
}
func (u *petController) ShowPet(ctx *gin.Context) {
	petID := ctx.Param("id")
	authHeader := ctx.GetHeader("Authorization")
	token, _ := u.jwtService.ValidateToken(authHeader)
	claims := token.Claims.(jwt.MapClaims)
	hotelID := fmt.Sprintf("%v", claims["hotel_id"])
	result, errShow, ok := u.petService.ShowPet(petID, hotelID)
	if ok == false {
		res := helper.BuildErrorResponse(
			"You don't have permission",
			"You are not the owner",
			helper.EmptyObj{},
		)
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}
	if errShow != nil {
		res := helper.BuildErrorResponse("Failed to show data", errShow.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Show data success", result)
	ctx.JSON(http.StatusOK, res)
}
