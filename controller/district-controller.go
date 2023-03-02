package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/helper"
	"github.com/yusufbagussh/pet_hotel_backend/service"
	"net/http"
)

// DistrictController is a contract of what districtController can do
type DistrictController interface {
	GetAllDistrict(ctx *gin.Context)
	CreateDistrict(ctx *gin.Context)
	DeleteDistrict(ctx *gin.Context)
	ShowDistrict(ctx *gin.Context)
	UpdateDistrict(ctx *gin.Context)
}

type districtController struct {
	districtService service.DistrictService
	jwtService      service.JWTService
}

// NewDistrictController is creating anew instance of DistrictControlller
func NewDistrictController(districtService service.DistrictService, jwtService service.JWTService) DistrictController {
	return &districtController{
		districtService: districtService,
		jwtService:      jwtService,
	}
}

func (u *districtController) CreateDistrict(ctx *gin.Context) {
	var CreateDistrict dto.CreateDistrict
	errDTO := ctx.ShouldBind(&CreateDistrict)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, errCreate := u.districtService.CreateDistrict(CreateDistrict)
	if errCreate != nil {
		res := helper.BuildErrorResponse("Failed to create district", errCreate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Create success", result)
	ctx.JSON(http.StatusOK, res)
}
func (u *districtController) UpdateDistrict(ctx *gin.Context) {
	var UpdateDistrict dto.UpdateDistrict
	errDTO := ctx.ShouldBind(&UpdateDistrict)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, errCreate := u.districtService.UpdateDistrict(UpdateDistrict)
	if errCreate != nil {
		res := helper.BuildErrorResponse("Failed to update district", errCreate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Update success", result)
	ctx.JSON(http.StatusOK, res)
}
func (u *districtController) DeleteDistrict(ctx *gin.Context) {
	districtID := ctx.Param("id")
	errDel := u.districtService.DeleteDistrict(districtID)
	if errDel != nil {
		res := helper.BuildErrorResponse("Failed to delete", errDel.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Delete success", nil)
	ctx.JSON(http.StatusOK, res)
}
func (u *districtController) ShowDistrict(ctx *gin.Context) {
	districtID := ctx.Param("id")
	result, errShow := u.districtService.ShowDistrict(districtID)
	if errShow != nil {
		res := helper.BuildErrorResponse("Failed to show data", errShow.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Show data success", result)
	ctx.JSON(http.StatusOK, res)
}

func (u *districtController) GetAllDistrict(ctx *gin.Context) {
	var filterPagination dto.FilterPagination

	errDTO := ctx.ShouldBind(&filterPagination)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var admins, pagination, err = u.districtService.GetAllDistrict(filterPagination)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get all data", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponsePage(true, "Get all data success", admins, pagination)
	ctx.JSON(http.StatusOK, res)
}
