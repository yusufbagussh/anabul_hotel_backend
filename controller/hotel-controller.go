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

// HotelController is a contract of what hotelController can do
type HotelController interface {
	GetAllHotel(ctx *gin.Context)
	CreateHotel(ctx *gin.Context)
	DeleteHotel(ctx *gin.Context)
	ShowHotel(ctx *gin.Context)
	UpdateHotel(ctx *gin.Context)
	GetProfileHotel(ctx *gin.Context)
	UpdateProfileHotel(ctx *gin.Context)
	CreateHotelAdmin(ctx *gin.Context)
}

type hotelController struct {
	hotelService service.HotelService
	jwtService   service.JWTService
}

// NewHotelController is creating anew instance of HotelControlller
func NewHotelController(hotelService service.HotelService, jwtService service.JWTService) HotelController {
	return &hotelController{
		hotelService: hotelService,
		jwtService:   jwtService,
	}
}

func (u *hotelController) CreateHotelAdmin(ctx *gin.Context) {
	var CreateHotel dto.CreateHotel
	errDTO := ctx.ShouldBind(&CreateHotel)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, errCreate := u.hotelService.CreateHotelAdmin(CreateHotel, ctx)
	if errCreate != nil {
		res := helper.BuildErrorResponse("Failed to create hotel", errCreate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Create success", result)
	ctx.JSON(http.StatusOK, res)
}

func (u *hotelController) GetProfileHotel(ctx *gin.Context) {
	//hotelID := ctx.Param("id")
	authHeader := ctx.GetHeader("Authorization")
	token, _ := u.jwtService.ValidateToken(authHeader)
	claims := token.Claims.(jwt.MapClaims)
	hotelID := fmt.Sprintf("%v", claims["hotel_id"])
	//checkAuthor := u.hotelService.IsAllowedHotel(hotelID, userID)
	//if checkAuthor == false {
	//	res := helper.BuildErrorResponse(
	//		"You don't have permission",
	//		"You are not the owner",
	//		helper.EmptyObj{},
	//	)
	//	ctx.JSON(http.StatusUnauthorized, res)
	//	return
	//}
	result, errShow := u.hotelService.ShowHotel(hotelID)
	if errShow != nil {
		res := helper.BuildErrorResponse("Failed to show data", errShow.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Show data success", result)
	ctx.JSON(http.StatusOK, res)
}

func (u *hotelController) UpdateProfileHotel(ctx *gin.Context) {
	var UpdateHotel dto.UpdateHotel
	errDTO := ctx.ShouldBind(&UpdateHotel)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	authHeader := ctx.GetHeader("Authorization")
	token, _ := u.jwtService.ValidateToken(authHeader)
	claims := token.Claims.(jwt.MapClaims)
	//userID := fmt.Sprintf("%v", claims["user_id"])
	hotelID := fmt.Sprintf("%v", claims["hotel_id"])
	//checkAuthor := u.hotelService.IsAllowedHotel(UpdateHotel.IDHotel, userID)
	//if checkAuthor == false {
	if UpdateHotel.IDHotel != hotelID {
		res := helper.BuildErrorResponse(
			"You dont have permission",
			"You are not the owner",
			helper.EmptyObj{},
		)
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}
	result, errCreate := u.hotelService.UpdateHotel(UpdateHotel, ctx)
	if errCreate != nil {
		res := helper.BuildErrorResponse("Failed to update hotel", errCreate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Update success", result)
	ctx.JSON(http.StatusOK, res)

}

func (u *hotelController) CreateHotel(ctx *gin.Context) {
	var CreateHotel dto.CreateHotel
	errDTO := ctx.ShouldBind(&CreateHotel)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, errCreate := u.hotelService.CreateHotel(CreateHotel, ctx)
	if errCreate != nil {
		res := helper.BuildErrorResponse("Failed to create hotel", errCreate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Create success", result)
	ctx.JSON(http.StatusOK, res)
}
func (u *hotelController) UpdateHotel(ctx *gin.Context) {
	var UpdateHotel dto.UpdateHotel
	errDTO := ctx.ShouldBind(&UpdateHotel)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, errCreate := u.hotelService.UpdateHotel(UpdateHotel, ctx)
	if errCreate != nil {
		res := helper.BuildErrorResponse("Failed to update hotel", errCreate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Update success", result)
	ctx.JSON(http.StatusOK, res)
}
func (u *hotelController) DeleteHotel(ctx *gin.Context) {
	hotelID := ctx.Param("id")
	errDel := u.hotelService.DeleteHotel(hotelID)
	if errDel != nil {
		res := helper.BuildErrorResponse("Failed to delete", errDel.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Delete success", nil)
	ctx.JSON(http.StatusOK, res)
}
func (u *hotelController) ShowHotel(ctx *gin.Context) {
	hotelID := ctx.Param("id")
	result, errShow := u.hotelService.ShowHotel(hotelID)
	if errShow != nil {
		res := helper.BuildErrorResponse("Failed to show data", errShow.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Show data success", result)
	ctx.JSON(http.StatusOK, res)
}

func (u *hotelController) GetAllHotel(ctx *gin.Context) {
	var filterPagination dto.FilterPagination

	errDTO := ctx.ShouldBind(&filterPagination)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var admins, pagination, err = u.hotelService.GetAllHotel(filterPagination)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get all data", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponsePage(true, "OK", admins, pagination)
	ctx.JSON(http.StatusOK, res)
}
