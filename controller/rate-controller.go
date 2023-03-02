package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/helper"
	"github.com/yusufbagussh/pet_hotel_backend/service"
	"net/http"
)

// RateController is a contract of what rateController can do
type RateController interface {
	GetAllRate(ctx *gin.Context)
	CreateRate(ctx *gin.Context)
	DeleteRate(ctx *gin.Context)
	ShowRate(ctx *gin.Context)
	UpdateRate(ctx *gin.Context)
}

type rateController struct {
	rateService service.RateService
	jwtService      service.JWTService
}

// NewRateController is creating anew instance of RateControlller
func NewRateController(rateService service.RateService, jwtService service.JWTService) RateController {
	return &rateController{
		rateService: rateService,
		jwtService:      jwtService,
	}
}

func (u *rateController) CreateRate(ctx *gin.Context) {
	var CreateRate dto.CreateRate
	errDTO := ctx.ShouldBind(&CreateRate)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, errCreate := u.rateService.CreateRate(CreateRate)
	if errCreate != nil {
		res := helper.BuildErrorResponse("Failed to create rate", errCreate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Create success", result)
	ctx.JSON(http.StatusOK, res)
}
func (u *rateController) UpdateRate(ctx *gin.Context) {
	var UpdateRate dto.UpdateRate
	errDTO := ctx.ShouldBind(&UpdateRate)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, errCreate := u.rateService.UpdateRate(UpdateRate)
	if errCreate != nil {
		res := helper.BuildErrorResponse("Failed to update rate", errCreate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Update success", result)
	ctx.JSON(http.StatusOK, res)
}
func (u *rateController) DeleteRate(ctx *gin.Context) {
	rateID := ctx.Param("id")
	errDel := u.rateService.DeleteRate(rateID)
	if errDel != nil {
		res := helper.BuildErrorResponse("Failed to delete", errDel.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Delete success", nil)
	ctx.JSON(http.StatusOK, res)
}
func (u *rateController) ShowRate(ctx *gin.Context) {
	rateID := ctx.Param("id")
	result, errShow := u.rateService.ShowRate(rateID)
	if errShow != nil {
		res := helper.BuildErrorResponse("Failed to show data", errShow.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Show data success", result)
	ctx.JSON(http.StatusOK, res)
}

func (u *rateController) GetAllRate(ctx *gin.Context) {
	var filterPagination dto.FilterPagination
	hotelID := ctx.Param("hotel_id")
	errDTO := ctx.ShouldBind(&filterPagination)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var admins, pagination, err = u.rateService.GetAllRate(hotelID, filterPagination)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get all data", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponsePage(true, "Get all data success", admins, pagination)
	ctx.JSON(http.StatusOK, res)
}
