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

// CageCategoryController is a contract of what cageCategoryController can do
type CageCategoryController interface {
	GetAllCageCategory(ctx *gin.Context)
	CreateCageCategory(ctx *gin.Context)
	DeleteCageCategory(ctx *gin.Context)
	ShowCageCategory(ctx *gin.Context)
	UpdateCageCategory(ctx *gin.Context)
}

type cageCategoryController struct {
	cageCategoryService service.CageCategoryService
	jwtService          service.JWTService
}

// NewCageCategoryController is creating anew instance of CageCategoryControlller
func NewCageCategoryController(cageCategoryService service.CageCategoryService, jwtService service.JWTService) CageCategoryController {
	return &cageCategoryController{
		cageCategoryService: cageCategoryService,
		jwtService:          jwtService,
	}
}

func (u *cageCategoryController) CreateCageCategory(ctx *gin.Context) {
	var CreateCageCategory dto.CreateCageCategory
	errDTO := ctx.ShouldBind(&CreateCageCategory)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, errCreate := u.cageCategoryService.CreateCageCategory(CreateCageCategory)
	if errCreate != nil {
		res := helper.BuildErrorResponse("Failed to create cageCategory", errCreate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Create success", result)
	ctx.JSON(http.StatusOK, res)
}
func (u *cageCategoryController) UpdateCageCategory(ctx *gin.Context) {
	var UpdateCageCategory dto.UpdateCageCategory
	authHeader := ctx.GetHeader("Authorization")
	token, _ := u.jwtService.ValidateToken(authHeader)
	claims := token.Claims.(jwt.MapClaims)
	hotelID := fmt.Sprintf("%v", claims["hotel_id"])
	if UpdateCageCategory.HotelID != hotelID {
		res := helper.BuildErrorResponse(
			"You don't have permission",
			"You are not the owner",
			helper.EmptyObj{},
		)
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}
	errDTO := ctx.ShouldBind(&UpdateCageCategory)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, errCreate := u.cageCategoryService.UpdateCageCategory(UpdateCageCategory)
	if errCreate != nil {
		res := helper.BuildErrorResponse("Failed to update cageCategory", errCreate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Update success", result)
	ctx.JSON(http.StatusOK, res)
}
func (u *cageCategoryController) DeleteCageCategory(ctx *gin.Context) {
	cageCategoryID := ctx.Param("id")
	authHeader := ctx.GetHeader("Authorization")
	token, _ := u.jwtService.ValidateToken(authHeader)
	claims := token.Claims.(jwt.MapClaims)
	hotelID := fmt.Sprintf("%v", claims["hotel_id"])
	errDel, ok := u.cageCategoryService.DeleteCageCategory(cageCategoryID, hotelID)
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
func (u *cageCategoryController) ShowCageCategory(ctx *gin.Context) {
	cageCategoryID := ctx.Param("id")
	authHeader := ctx.GetHeader("Authorization")
	token, _ := u.jwtService.ValidateToken(authHeader)
	claims := token.Claims.(jwt.MapClaims)
	hotelID := fmt.Sprintf("%v", claims["hotel_id"])
	result, errShow, ok := u.cageCategoryService.ShowCageCategory(cageCategoryID, hotelID)
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

func (u *cageCategoryController) GetAllCageCategory(ctx *gin.Context) {
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

	var cageCategorys, pagination, err = u.cageCategoryService.GetAllCageCategory(hotelID, filterPagination)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get all data", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponsePage(true, "Get all data success", cageCategorys, pagination)
	ctx.JSON(http.StatusOK, res)
}
