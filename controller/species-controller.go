package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/helper"
	"github.com/yusufbagussh/pet_hotel_backend/service"
	"net/http"
)

// SpeciesController is a contract of what speciesController can do
type SpeciesController interface {
	GetSpecies(ctx *gin.Context)
	CreateSpecies(ctx *gin.Context)
	DeleteSpecies(ctx *gin.Context)
	ShowSpecies(ctx *gin.Context)
	UpdateSpecies(ctx *gin.Context)
}

type speciesController struct {
	speciesService service.SpeciesService
	jwtService     service.JWTService
}

// NewSpeciesController is creating anew instance of SpeciesControlller
func NewSpeciesController(speciesService service.SpeciesService, jwtService service.JWTService) SpeciesController {
	return &speciesController{
		speciesService: speciesService,
		jwtService:     jwtService,
	}
}

func (u *speciesController) CreateSpecies(ctx *gin.Context) {
	var CreateSpecies dto.CreateSpecies
	errDTO := ctx.ShouldBind(&CreateSpecies)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, errCreate := u.speciesService.CreateSpecies(CreateSpecies)
	if errCreate != nil {
		res := helper.BuildErrorResponse("Failed to create species", errCreate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Create success", result)
	ctx.JSON(http.StatusOK, res)
}

func (u *speciesController) GetSpecies(ctx *gin.Context) {
	var filterPagination dto.FilterPagination

	errDTO := ctx.ShouldBind(&filterPagination)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	var staffs, pagination, errGet = u.speciesService.GetAllSpecies(filterPagination)
	if errGet != nil {
		res := helper.BuildErrorResponse("Failed to get all data", errGet.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponsePage(true, "OK", staffs, pagination)
	ctx.JSON(http.StatusOK, res)
}

func (u *speciesController) UpdateSpecies(ctx *gin.Context) {
	var UpdateSpecies dto.UpdateSpecies
	errDTO := ctx.ShouldBind(&UpdateSpecies)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, errCreate := u.speciesService.UpdateSpecies(UpdateSpecies)
	if errCreate != nil {
		res := helper.BuildErrorResponse("Failed to update species", errCreate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Update success", result)
	ctx.JSON(http.StatusOK, res)
}
func (u *speciesController) DeleteSpecies(ctx *gin.Context) {
	speciesID := ctx.Param("id")
	errDel := u.speciesService.DeleteSpecies(speciesID)
	if errDel != nil {
		res := helper.BuildErrorResponse("Failed to delete", errDel.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Delete success", nil)
	ctx.JSON(http.StatusOK, res)
}
func (u *speciesController) ShowSpecies(ctx *gin.Context) {
	speciesID := ctx.Param("id")
	result, errShow := u.speciesService.ShowSpecies(speciesID)
	if errShow != nil {
		res := helper.BuildErrorResponse("Failed to show data", errShow.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Show data success", result)
	ctx.JSON(http.StatusOK, res)
}
