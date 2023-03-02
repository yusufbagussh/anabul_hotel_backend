package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/helper"
	"github.com/yusufbagussh/pet_hotel_backend/service"
	"net/http"
)

// ReservationConditionController is a contract of what reservationConditionController can do
type ReservationConditionController interface {
	GetAllReservationCondition(ctx *gin.Context)
	CreateReservationCondition(ctx *gin.Context)
	DeleteReservationCondition(ctx *gin.Context)
	ShowReservationCondition(ctx *gin.Context)
	UpdateReservationCondition(ctx *gin.Context)
}

type reservationConditionController struct {
	reservationConditionService service.ReservationConditionService
	jwtService                  service.JWTService
}

// NewReservationConditionController is creating anew instance of ReservationConditionControlller
func NewReservationConditionController(reservationConditionService service.ReservationConditionService, jwtService service.JWTService) ReservationConditionController {
	return &reservationConditionController{
		reservationConditionService: reservationConditionService,
		jwtService:                  jwtService,
	}
}

func (u *reservationConditionController) CreateReservationCondition(ctx *gin.Context) {
	var CreateReservationCondition dto.CreateReservationCondition
	errDTO := ctx.ShouldBind(&CreateReservationCondition)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, errCreate := u.reservationConditionService.CreateReservationCondition(CreateReservationCondition, ctx)
	if errCreate != nil {
		res := helper.BuildErrorResponse("Failed to create reservationCondition", errCreate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Create success", result)
	ctx.JSON(http.StatusOK, res)
}
func (u *reservationConditionController) UpdateReservationCondition(ctx *gin.Context) {
	var UpdateReservationCondition dto.UpdateReservationCondition
	errDTO := ctx.ShouldBind(&UpdateReservationCondition)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, errCreate := u.reservationConditionService.UpdateReservationCondition(UpdateReservationCondition)
	if errCreate != nil {
		res := helper.BuildErrorResponse("Failed to update reservationCondition", errCreate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Update success", result)
	ctx.JSON(http.StatusOK, res)
}
func (u *reservationConditionController) DeleteReservationCondition(ctx *gin.Context) {
	reservationConditionID := ctx.Param("id")
	errDel := u.reservationConditionService.DeleteReservationCondition(reservationConditionID)
	if errDel != nil {
		res := helper.BuildErrorResponse("Failed to delete", errDel.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Delete success", nil)
	ctx.JSON(http.StatusOK, res)
}
func (u *reservationConditionController) ShowReservationCondition(ctx *gin.Context) {
	reservationConditionID := ctx.Param("id")
	result, errShow := u.reservationConditionService.ShowReservationCondition(reservationConditionID)
	if errShow != nil {
		res := helper.BuildErrorResponse("Failed to show data", errShow.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Show data success", result)
	ctx.JSON(http.StatusOK, res)
}

func (u *reservationConditionController) GetAllReservationCondition(ctx *gin.Context) {
	var filterPagination dto.FilterPagination
	hotelID := ctx.Param("hotel_id")
	errDTO := ctx.ShouldBind(&filterPagination)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var admins, pagination, err = u.reservationConditionService.GetAllReservationCondition(hotelID, filterPagination)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get all data", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponsePage(true, "Get all data success", admins, pagination)
	ctx.JSON(http.StatusOK, res)
}
