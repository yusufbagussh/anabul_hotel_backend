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

// ServiceDetailController is a contract of what serviceDetailController can do
type ServiceDetailController interface {
	GetAllServiceDetail(ctx *gin.Context)
	CreateServiceDetail(ctx *gin.Context)
	DeleteServiceDetail(ctx *gin.Context)
	ShowServiceDetail(ctx *gin.Context)
	UpdateServiceDetail(ctx *gin.Context)
}

type serviceDetailController struct {
	serviceDetailService service.ServiceDetailService
	jwtService           service.JWTService
}

// NewServiceDetailController is creating anew instance of ServiceDetailControlller
func NewServiceDetailController(serviceDetailService service.ServiceDetailService, jwtService service.JWTService) ServiceDetailController {
	return &serviceDetailController{
		serviceDetailService: serviceDetailService,
		jwtService:           jwtService,
	}
}

func (u *serviceDetailController) CreateServiceDetail(ctx *gin.Context) {
	var CreateServiceDetail dto.CreateServiceDetail
	errDTO := ctx.ShouldBind(&CreateServiceDetail)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, errCreate := u.serviceDetailService.CreateServiceDetail(CreateServiceDetail)
	if errCreate != nil {
		res := helper.BuildErrorResponse("Failed to create serviceDetail", errCreate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Create success", result)
	ctx.JSON(http.StatusOK, res)
}
func (u *serviceDetailController) UpdateServiceDetail(ctx *gin.Context) {
	var UpdateServiceDetail dto.UpdateServiceDetail
	errDTO := ctx.ShouldBind(&UpdateServiceDetail)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	authHeader := ctx.GetHeader("Authorization")
	token, _ := u.jwtService.ValidateToken(authHeader)
	claims := token.Claims.(jwt.MapClaims)
	hotelID := fmt.Sprintf("%v", claims["hotel_id"])
	if UpdateServiceDetail.HotelID != hotelID {
		res := helper.BuildErrorResponse(
			"You don't have permission",
			"You are not the owner",
			helper.EmptyObj{},
		)
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}
	result, errCreate := u.serviceDetailService.UpdateServiceDetail(UpdateServiceDetail)
	if errCreate != nil {
		res := helper.BuildErrorResponse("Failed to update serviceDetail", errCreate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Update success", result)
	ctx.JSON(http.StatusOK, res)
}
func (u *serviceDetailController) DeleteServiceDetail(ctx *gin.Context) {
	serviceDetailID := ctx.Param("id")
	authHeader := ctx.GetHeader("Authorization")
	token, _ := u.jwtService.ValidateToken(authHeader)
	claims := token.Claims.(jwt.MapClaims)
	hotelID := fmt.Sprintf("%v", claims["hotel_id"])
	errDel, ok := u.serviceDetailService.DeleteServiceDetail(serviceDetailID, hotelID)
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
func (u *serviceDetailController) ShowServiceDetail(ctx *gin.Context) {
	serviceDetailID := ctx.Param("id")
	authHeader := ctx.GetHeader("Authorization")
	token, _ := u.jwtService.ValidateToken(authHeader)
	claims := token.Claims.(jwt.MapClaims)
	hotelID := fmt.Sprintf("%v", claims["hotel_id"])
	result, errShow, ok := u.serviceDetailService.ShowServiceDetail(serviceDetailID, hotelID)
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

func (u *serviceDetailController) GetAllServiceDetail(ctx *gin.Context) {
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

	var serviceDetails, pagination, err = u.serviceDetailService.GetAllServiceDetail(hotelID, filterPagination)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get all data", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponsePage(true, "Get all data success", serviceDetails, pagination)
	ctx.JSON(http.StatusOK, res)
}
