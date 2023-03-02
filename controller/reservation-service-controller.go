package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/helper"
	"github.com/yusufbagussh/pet_hotel_backend/service"
	"net/http"
)

// ReservationServiceController is a contract of what reservationServiceController can do
type ReservationServiceController interface {
	GetAllReservationService(ctx *gin.Context)
	CreateReservationService(ctx *gin.Context)
	DeleteReservationService(ctx *gin.Context)
	ShowReservationService(ctx *gin.Context)
	UpdateReservationService(ctx *gin.Context)
}

type reservationServiceController struct {
	reservationServiceService service.ReservationServiceService
	jwtService                service.JWTService
}

// NewReservationServiceController is creating anew instance of ReservationServiceControlller
func NewReservationServiceController(reservationServiceService service.ReservationServiceService, jwtService service.JWTService) ReservationServiceController {
	return &reservationServiceController{
		reservationServiceService: reservationServiceService,
		jwtService:                jwtService,
	}
}

func (u *reservationServiceController) CreateReservationService(ctx *gin.Context) {
	var CreateReservationService dto.CreateReservationService
	errDTO := ctx.ShouldBind(&CreateReservationService)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, errCreate := u.reservationServiceService.CreateReservationService(CreateReservationService)
	if errCreate != nil {
		res := helper.BuildErrorResponse("Failed to create reservationService", errCreate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Create success", result)
	ctx.JSON(http.StatusOK, res)
}
func (u *reservationServiceController) UpdateReservationService(ctx *gin.Context) {
	var UpdateReservationService dto.UpdateReservationService
	errDTO := ctx.ShouldBind(&UpdateReservationService)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, errCreate := u.reservationServiceService.UpdateReservationService(UpdateReservationService)
	if errCreate != nil {
		res := helper.BuildErrorResponse("Failed to update reservationService", errCreate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Update success", result)
	ctx.JSON(http.StatusOK, res)
}
func (u *reservationServiceController) DeleteReservationService(ctx *gin.Context) {
	reservationServiceID := ctx.Param("id")
	errDel := u.reservationServiceService.DeleteReservationService(reservationServiceID)
	if errDel != nil {
		res := helper.BuildErrorResponse("Failed to delete", errDel.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Delete success", nil)
	ctx.JSON(http.StatusOK, res)
}
func (u *reservationServiceController) ShowReservationService(ctx *gin.Context) {
	reservationServiceID := ctx.Param("id")
	result, errShow := u.reservationServiceService.ShowReservationService(reservationServiceID)
	if errShow != nil {
		res := helper.BuildErrorResponse("Failed to show data", errShow.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Show data success", result)
	ctx.JSON(http.StatusOK, res)
}

func (u *reservationServiceController) GetAllReservationService(ctx *gin.Context) {
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

	var reservationServicees, pagination, err = u.reservationServiceService.GetAllReservationService(filterPagination)
	//type reservationServicePage struct {
	//	ReservationServicees    []entity.ReservationService
	//	Pagination dto.Pagination
	//}
	//reservationServiceesPage := reservationServicePage{
	//	reservationServicees,
	//	pagination,
	//}
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get all data", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponsePage(true, "Get all data success", reservationServicees, pagination)
	ctx.JSON(http.StatusOK, res)
	//ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "Success get all data reservationService", "data": reservationServicees, "pagination": pagination})
}
