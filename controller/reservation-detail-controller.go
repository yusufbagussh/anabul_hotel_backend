package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/helper"
	"github.com/yusufbagussh/pet_hotel_backend/service"
	"net/http"
)

// ReservationDetailController is a contract of what reservationDetailController can do
type ReservationDetailController interface {
	GetAllReservationDetail(ctx *gin.Context)
	CreateReservationDetail(ctx *gin.Context)
	DeleteReservationDetail(ctx *gin.Context)
	ShowReservationDetail(ctx *gin.Context)
	UpdateReservationDetail(ctx *gin.Context)
}

type reservationDetailController struct {
	reservationDetailService service.ReservationDetailService
	jwtService               service.JWTService
}

// NewReservationDetailController is creating anew instance of ReservationDetailControlller
func NewReservationDetailController(reservationDetailService service.ReservationDetailService, jwtService service.JWTService) ReservationDetailController {
	return &reservationDetailController{
		reservationDetailService: reservationDetailService,
		jwtService:               jwtService,
	}
}

func (u *reservationDetailController) CreateReservationDetail(ctx *gin.Context) {
	var CreateReservationDetail dto.CreateReservationDetail
	errDTO := ctx.ShouldBind(&CreateReservationDetail)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, errCreate := u.reservationDetailService.CreateReservationDetail(CreateReservationDetail)
	if errCreate != nil {
		res := helper.BuildErrorResponse("Failed to create reservationDetail", errCreate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Create success", result)
	ctx.JSON(http.StatusOK, res)
}
func (u *reservationDetailController) UpdateReservationDetail(ctx *gin.Context) {
	var UpdateReservationDetail dto.UpdateReservationDetail
	errDTO := ctx.ShouldBind(&UpdateReservationDetail)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, errCreate := u.reservationDetailService.UpdateReservationDetail(UpdateReservationDetail)
	if errCreate != nil {
		res := helper.BuildErrorResponse("Failed to update reservationDetail", errCreate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Update success", result)
	ctx.JSON(http.StatusOK, res)
}
func (u *reservationDetailController) DeleteReservationDetail(ctx *gin.Context) {
	reservationDetailID := ctx.Param("id")
	errDel := u.reservationDetailService.DeleteReservationDetail(reservationDetailID)
	if errDel != nil {
		res := helper.BuildErrorResponse("Failed to delete", errDel.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Delete success", nil)
	ctx.JSON(http.StatusOK, res)
}
func (u *reservationDetailController) ShowReservationDetail(ctx *gin.Context) {
	reservationDetailID := ctx.Param("id")
	result, errShow := u.reservationDetailService.ShowReservationDetail(reservationDetailID)
	if errShow != nil {
		res := helper.BuildErrorResponse("Failed to show data", errShow.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Show data success", result)
	ctx.JSON(http.StatusOK, res)
}

func (u *reservationDetailController) GetAllReservationDetail(ctx *gin.Context) {
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

	var reservationDetailes, pagination, err = u.reservationDetailService.GetAllReservationDetail(filterPagination)
	//type reservationDetailPage struct {
	//	ReservationDetailes    []entity.ReservationDetail
	//	Pagination dto.Pagination
	//}
	//reservationDetailesPage := reservationDetailPage{
	//	reservationDetailes,
	//	pagination,
	//}
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get all data", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponsePage(true, "Get all data success", reservationDetailes, pagination)
	ctx.JSON(http.StatusOK, res)
	//ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "Success get all data reservationDetail", "data": reservationDetailes, "pagination": pagination})
}
