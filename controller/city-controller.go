package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/helper"
	"github.com/yusufbagussh/pet_hotel_backend/service"
	"net/http"
)

// CityController is a contract of what cityController can do
type CityController interface {
	GetAllCity(ctx *gin.Context)
	CreateCity(ctx *gin.Context)
	DeleteCity(ctx *gin.Context)
	ShowCity(ctx *gin.Context)
	UpdateCity(ctx *gin.Context)
}

type cityController struct {
	cityService service.CityService
	jwtService  service.JWTService
}

// NewCityController is creating anew instance of CityControlller
func NewCityController(cityService service.CityService, jwtService service.JWTService) CityController {
	return &cityController{
		cityService: cityService,
		jwtService:  jwtService,
	}
}

func (u *cityController) CreateCity(ctx *gin.Context) {
	var CreateCity dto.CreateCity
	errDTO := ctx.ShouldBind(&CreateCity)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, errCreate := u.cityService.CreateCity(CreateCity)
	if errCreate != nil {
		res := helper.BuildErrorResponse("Failed to create city", errCreate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Create success", result)
	ctx.JSON(http.StatusOK, res)
}
func (u *cityController) UpdateCity(ctx *gin.Context) {
	var UpdateCity dto.UpdateCity
	errDTO := ctx.ShouldBind(&UpdateCity)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, errCreate := u.cityService.UpdateCity(UpdateCity)
	if errCreate != nil {
		res := helper.BuildErrorResponse("Failed to update city", errCreate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Update success", result)
	ctx.JSON(http.StatusOK, res)
}
func (u *cityController) DeleteCity(ctx *gin.Context) {
	cityID := ctx.Param("id")
	errDel := u.cityService.DeleteCity(cityID)
	if errDel != nil {
		res := helper.BuildErrorResponse("Failed to delete", errDel.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Delete success", nil)
	ctx.JSON(http.StatusOK, res)
}
func (u *cityController) ShowCity(ctx *gin.Context) {
	cityID := ctx.Param("id")
	result, errShow := u.cityService.ShowCity(cityID)
	if errShow != nil {
		res := helper.BuildErrorResponse("Failed to show data", errShow.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Show data success", result)
	ctx.JSON(http.StatusOK, res)
}

func (u *cityController) GetAllCity(ctx *gin.Context) {
	var filterPagination dto.FilterPagination

	errDTO := ctx.ShouldBind(&filterPagination)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var admins, pagination, err = u.cityService.GetAllCity(filterPagination)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get all data", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponsePage(true, "Get all data success", admins, pagination)
	ctx.JSON(http.StatusOK, res)
}
