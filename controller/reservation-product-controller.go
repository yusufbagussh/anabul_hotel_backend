package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/helper"
	"github.com/yusufbagussh/pet_hotel_backend/service"
	"net/http"
)

// ReservationProductController is a contract of what reservationProductController can do
type ReservationProductController interface {
	GetAllReservationProduct(ctx *gin.Context)
	CreateReservationProduct(ctx *gin.Context)
	DeleteReservationProduct(ctx *gin.Context)
	ShowReservationProduct(ctx *gin.Context)
	UpdateReservationProduct(ctx *gin.Context)
}

type reservationProductController struct {
	reservationProduct service.ReservationProductService
	jwtProduct         service.JWTService
}

// NewReservationProductController is creating anew instance of ReservationProductControlller
func NewReservationProductController(reservationProduct service.ReservationProductService, jwtProduct service.JWTService) ReservationProductController {
	return &reservationProductController{
		reservationProduct: reservationProduct,
		jwtProduct:         jwtProduct,
	}
}

func (u *reservationProductController) CreateReservationProduct(ctx *gin.Context) {
	var CreateReservationProduct dto.CreateReservationProduct
	errDTO := ctx.ShouldBind(&CreateReservationProduct)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, errCreate := u.reservationProduct.CreateReservationProduct(CreateReservationProduct)
	if errCreate != nil {
		res := helper.BuildErrorResponse("Failed to create reservationProduct", errCreate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Create success", result)
	ctx.JSON(http.StatusOK, res)
}
func (u *reservationProductController) UpdateReservationProduct(ctx *gin.Context) {
	var UpdateReservationProduct dto.UpdateReservationProduct
	errDTO := ctx.ShouldBind(&UpdateReservationProduct)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, errCreate := u.reservationProduct.UpdateReservationProduct(UpdateReservationProduct)
	if errCreate != nil {
		res := helper.BuildErrorResponse("Failed to update reservationProduct", errCreate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Update success", result)
	ctx.JSON(http.StatusOK, res)
}
func (u *reservationProductController) DeleteReservationProduct(ctx *gin.Context) {
	reservationProductID := ctx.Param("id")
	errDel := u.reservationProduct.DeleteReservationProduct(reservationProductID)
	if errDel != nil {
		res := helper.BuildErrorResponse("Failed to delete", errDel.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Delete success", nil)
	ctx.JSON(http.StatusOK, res)
}
func (u *reservationProductController) ShowReservationProduct(ctx *gin.Context) {
	reservationProductID := ctx.Param("id")
	result, errShow := u.reservationProduct.ShowReservationProduct(reservationProductID)
	if errShow != nil {
		res := helper.BuildErrorResponse("Failed to show data", errShow.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Show data success", result)
	ctx.JSON(http.StatusOK, res)
}

func (u *reservationProductController) GetAllReservationProduct(ctx *gin.Context) {
	//search := ctx.Query("search")
	//sortBy := ctx.Query("sortBy")
	//orderBy := ctx.Query("orderBy")
	//page, _ := strconv.Atoi(ctx.Query("page"))
	//perPage, _ := strconv.Atoi(ctx.Query("perPage"))
	//
	//filterPagination := dto.FilterPagination{
	//	Search:  search,
	//	SortBy:  sortBy,
	//	OrderBy: orderBy,
	//	Page:    uint32(page),
	//	PerPage: uint32(perPage),
	//}
	var filterPagination dto.FilterPagination

	errDTO := ctx.ShouldBind(&filterPagination)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var reservationProductes, pagination, err = u.reservationProduct.GetAllReservationProduct(filterPagination)
	//type reservationProductPage struct {
	//	ReservationProductes    []entity.ReservationProduct
	//	Pagination dto.Pagination
	//}
	//reservationProductesPage := reservationProductPage{
	//	reservationProductes,
	//	pagination,
	//}
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get all data", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponsePage(true, "Get all data success", reservationProductes, pagination)
	ctx.JSON(http.StatusOK, res)
	//ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "Success get all data reservationProduct", "data": reservationProductes, "pagination": pagination})
}
