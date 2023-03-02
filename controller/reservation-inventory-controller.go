package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/helper"
	"github.com/yusufbagussh/pet_hotel_backend/service"
	"net/http"
)

// ReservationInventoryController is a contract of what reservationInventoryController can do
type ReservationInventoryController interface {
	GetAllReservationInventory(ctx *gin.Context)
	CreateReservationInventory(ctx *gin.Context)
	DeleteReservationInventory(ctx *gin.Context)
	ShowReservationInventory(ctx *gin.Context)
	UpdateReservationInventory(ctx *gin.Context)
}

type reservationInventoryController struct {
	reservationInventoryService service.ReservationInventoryService
	jwtService                  service.JWTService
}

// NewReservationInventoryController is creating anew instance of ReservationInventoryControlller
func NewReservationInventoryController(reservationInventoryService service.ReservationInventoryService, jwtService service.JWTService) ReservationInventoryController {
	return &reservationInventoryController{
		reservationInventoryService: reservationInventoryService,
		jwtService:                  jwtService,
	}
}

func (u *reservationInventoryController) CreateReservationInventory(ctx *gin.Context) {
	var CreateReservationInventory dto.CreateReservationInventory
	errDTO := ctx.ShouldBind(&CreateReservationInventory)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, errCreate := u.reservationInventoryService.CreateReservationInventory(CreateReservationInventory)
	if errCreate != nil {
		res := helper.BuildErrorResponse("Failed to create reservationInventory", errCreate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Create success", result)
	ctx.JSON(http.StatusOK, res)
}
func (u *reservationInventoryController) UpdateReservationInventory(ctx *gin.Context) {
	var UpdateReservationInventory dto.UpdateReservationInventory
	errDTO := ctx.ShouldBind(&UpdateReservationInventory)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, errCreate := u.reservationInventoryService.UpdateReservationInventory(UpdateReservationInventory)
	if errCreate != nil {
		res := helper.BuildErrorResponse("Failed to update reservationInventory", errCreate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Update success", result)
	ctx.JSON(http.StatusOK, res)
}
func (u *reservationInventoryController) DeleteReservationInventory(ctx *gin.Context) {
	reservationInventoryID := ctx.Param("id")
	errDel := u.reservationInventoryService.DeleteReservationInventory(reservationInventoryID)
	if errDel != nil {
		res := helper.BuildErrorResponse("Failed to delete", errDel.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Delete success", nil)
	ctx.JSON(http.StatusOK, res)
}
func (u *reservationInventoryController) ShowReservationInventory(ctx *gin.Context) {
	reservationInventoryID := ctx.Param("id")
	result, errShow := u.reservationInventoryService.ShowReservationInventory(reservationInventoryID)
	if errShow != nil {
		res := helper.BuildErrorResponse("Failed to show data", errShow.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Show data success", result)
	ctx.JSON(http.StatusOK, res)
}

func (u *reservationInventoryController) GetAllReservationInventory(ctx *gin.Context) {
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

	var reservationInventoryes, pagination, err = u.reservationInventoryService.GetAllReservationInventory(filterPagination)
	//type reservationInventoryPage struct {
	//	ReservationInventoryes    []entity.ReservationInventory
	//	Pagination dto.Pagination
	//}
	//reservationInventoryesPage := reservationInventoryPage{
	//	reservationInventoryes,
	//	pagination,
	//}
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get all data", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponsePage(true, "Get all data success", reservationInventoryes, pagination)
	ctx.JSON(http.StatusOK, res)
	//ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "Success get all data reservationInventory", "data": reservationInventoryes, "pagination": pagination})
}
