package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/helper"
	"github.com/yusufbagussh/pet_hotel_backend/service"
	"net/http"
)

// ResponseController is a contract of what responseController can do
type ResponseController interface {
	GetAllResponse(ctx *gin.Context)
	CreateResponse(ctx *gin.Context)
	DeleteResponse(ctx *gin.Context)
	ShowResponse(ctx *gin.Context)
	UpdateResponse(ctx *gin.Context)
}

type responseController struct {
	responseService service.ResponseService
	jwtService      service.JWTService
}

// NewResponseController is creating anew instance of ResponseControlller
func NewResponseController(responseService service.ResponseService, jwtService service.JWTService) ResponseController {
	return &responseController{
		responseService: responseService,
		jwtService:      jwtService,
	}
}

func (u *responseController) CreateResponse(ctx *gin.Context) {
	var CreateResponse dto.CreateResponse
	errDTO := ctx.ShouldBind(&CreateResponse)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, errCreate := u.responseService.CreateResponse(CreateResponse)
	if errCreate != nil {
		res := helper.BuildErrorResponse("Failed to create response", errCreate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Create success", result)
	ctx.JSON(http.StatusOK, res)
}
func (u *responseController) UpdateResponse(ctx *gin.Context) {
	var UpdateResponse dto.UpdateResponse
	errDTO := ctx.ShouldBind(&UpdateResponse)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, errCreate := u.responseService.UpdateResponse(UpdateResponse)
	if errCreate != nil {
		res := helper.BuildErrorResponse("Failed to update response", errCreate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Update success", result)
	ctx.JSON(http.StatusOK, res)
}
func (u *responseController) DeleteResponse(ctx *gin.Context) {
	responseID := ctx.Param("id")
	errDel := u.responseService.DeleteResponse(responseID)
	if errDel != nil {
		res := helper.BuildErrorResponse("Failed to delete", errDel.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Delete success", nil)
	ctx.JSON(http.StatusOK, res)
}
func (u *responseController) ShowResponse(ctx *gin.Context) {
	responseID := ctx.Param("id")
	result, errShow := u.responseService.ShowResponse(responseID)
	if errShow != nil {
		res := helper.BuildErrorResponse("Failed to show data", errShow.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Show data success", result)
	ctx.JSON(http.StatusOK, res)
}

func (u *responseController) GetAllResponse(ctx *gin.Context) {
	var filterPagination dto.FilterPagination
	hotelID := ctx.Param("hotel_id")
	errDTO := ctx.ShouldBind(&filterPagination)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var admins, pagination, err = u.responseService.GetAllResponse(hotelID, filterPagination)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get all data", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponsePage(true, "Get all data success", admins, pagination)
	ctx.JSON(http.StatusOK, res)
}
