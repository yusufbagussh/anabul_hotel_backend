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

// CageController is a contract of what cageController can do
type CageController interface {
	GetAllCage(ctx *gin.Context)
	CreateCage(ctx *gin.Context)
	DeleteCage(ctx *gin.Context)
	ShowCage(ctx *gin.Context)
	UpdateCage(ctx *gin.Context)
}

type cageController struct {
	cageService service.CageService
	jwtService  service.JWTService
}

// NewCageController is creating anew instance of CageControlller
func NewCageController(cageService service.CageService, jwtService service.JWTService) CageController {
	return &cageController{
		cageService: cageService,
		jwtService:  jwtService,
	}
}

func (u *cageController) CreateCage(ctx *gin.Context) {
	var CreateCage dto.CreateCage
	errDTO := ctx.ShouldBind(&CreateCage)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, errCreate := u.cageService.CreateCage(CreateCage)
	if errCreate != nil {
		res := helper.BuildErrorResponse("Failed to create cage", errCreate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Create success", result)
	ctx.JSON(http.StatusOK, res)
}
func (u *cageController) UpdateCage(ctx *gin.Context) {
	var UpdateCage dto.UpdateCage
	authHeader := ctx.GetHeader("Authorization")
	token, _ := u.jwtService.ValidateToken(authHeader)
	claims := token.Claims.(jwt.MapClaims)
	hotelID := fmt.Sprintf("%v", claims["hotel_id"])
	if UpdateCage.HotelID != hotelID {
		res := helper.BuildErrorResponse(
			"You don't have permission",
			"You are not the owner",
			helper.EmptyObj{},
		)
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}
	errDTO := ctx.ShouldBind(&UpdateCage)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, errCreate := u.cageService.UpdateCage(UpdateCage)
	if errCreate != nil {
		res := helper.BuildErrorResponse("Failed to update cage", errCreate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Update success", result)
	ctx.JSON(http.StatusOK, res)
}
func (u *cageController) DeleteCage(ctx *gin.Context) {
	cageID := ctx.Param("id")
	authHeader := ctx.GetHeader("Authorization")
	token, _ := u.jwtService.ValidateToken(authHeader)
	claims := token.Claims.(jwt.MapClaims)
	hotelID := fmt.Sprintf("%v", claims["hotel_id"])
	errDel, ok := u.cageService.DeleteCage(cageID, hotelID)
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
func (u *cageController) ShowCage(ctx *gin.Context) {
	cageID := ctx.Param("id")
	authHeader := ctx.GetHeader("Authorization")
	token, _ := u.jwtService.ValidateToken(authHeader)
	claims := token.Claims.(jwt.MapClaims)
	hotelID := fmt.Sprintf("%v", claims["hotel_id"])
	result, errShow, ok := u.cageService.ShowCage(cageID, hotelID)
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

func (u *cageController) GetAllCage(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, _ := u.jwtService.ValidateToken(authHeader)
	claims := token.Claims.(jwt.MapClaims)
	hotelID := fmt.Sprintf("%v", claims["hotel_id"])

	var filterPagination dto.FilterPagination
	errDTO := ctx.ShouldBind(&filterPagination)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var cages, pagination, err = u.cageService.GetAllCage(hotelID, filterPagination)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get all data", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponsePage(true, "Get all data success", cages, pagination)
	ctx.JSON(http.StatusOK, res)
}
