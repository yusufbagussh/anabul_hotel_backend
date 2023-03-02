package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/helper"
	"github.com/yusufbagussh/pet_hotel_backend/service"
	"net/http"
	"strconv"
)

// RequestController is a contract of what requestController can do
type RequestController interface {
	GetAllRequest(ctx *gin.Context)
	CreateRequest(ctx *gin.Context)
	DeleteRequest(ctx *gin.Context)
	ShowRequest(ctx *gin.Context)
	UpdateRequest(ctx *gin.Context)
}

type requestController struct {
	requestService service.RequestService
	jwtService     service.JWTService
}

// NewRequestController is creating anew instance of RequestControlller
func NewRequestController(requestService service.RequestService, jwtService service.JWTService) RequestController {
	return &requestController{
		requestService: requestService,
		jwtService:     jwtService,
	}
}

func (u *requestController) CreateRequest(ctx *gin.Context) {
	var CreateRequest dto.CreateRequest
	errDTO := ctx.ShouldBind(&CreateRequest)
	fmt.Println("create request 1")
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to binding request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	fmt.Println("create request 2")
	result, errCreate := u.requestService.CreateRequest(CreateRequest, ctx)
	if errCreate != nil {
		res := helper.BuildErrorResponse("Failed to create request", errCreate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Create success", result)
	ctx.JSON(http.StatusOK, res)
}
func (u *requestController) UpdateRequest(ctx *gin.Context) {
	var UpdateRequest dto.UpdateRequest
	errDTO := ctx.ShouldBind(&UpdateRequest)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	//result, errUpdate, emailData := u.requestService.UpdateRequest(UpdateRequest, ctx)
	result, errUpdate, errSend := u.requestService.UpdateRequest(UpdateRequest, ctx)
	if errUpdate != nil {
		res := helper.BuildErrorResponse("Failed to update request", errUpdate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	//errSend := utils.SendEmailStatus(&result, &emailData, "emailReset.html")
	if errSend != nil {
		res := helper.BuildErrorResponse("Failed to send Email", errSend.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Update success", result)
	ctx.JSON(http.StatusOK, res)
}
func (u *requestController) DeleteRequest(ctx *gin.Context) {
	requestID := ctx.Param("id")
	errDel := u.requestService.DeleteRequest(requestID)
	if errDel != nil {
		res := helper.BuildErrorResponse("Failed to delete", errDel.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Delete success", nil)
	ctx.JSON(http.StatusOK, res)
}
func (u *requestController) ShowRequest(ctx *gin.Context) {
	requestID := ctx.Param("id")
	result, errShow := u.requestService.ShowRequest(requestID)
	if errShow != nil {
		res := helper.BuildErrorResponse("Failed to show data", errShow.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Show data success", result)
	ctx.JSON(http.StatusOK, res)
}

func (u *requestController) GetAllRequest(ctx *gin.Context) {
	search := ctx.Query("search")
	sortBy := ctx.Query("sortBy")
	orderBy := ctx.Query("orderBy")
	page, _ := strconv.Atoi(ctx.Query("page"))
	perPage, _ := strconv.Atoi(ctx.Query("perPage"))

	filterPagination := dto.FilterPagination{
		Search:  search,
		SortBy:  sortBy,
		OrderBy: orderBy,
		Page:    uint32(page),
		PerPage: uint32(perPage),
	}

	var requestes, pagination, err = u.requestService.GetAllRequest(filterPagination)
	//type requestPage struct {
	//	Requestes    []entity.Request
	//	Pagination dto.Pagination
	//}
	//requestesPage := requestPage{
	//	requestes,
	//	pagination,
	//}
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get all data", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponsePage(true, "Get all data success", requestes, pagination)
	ctx.JSON(http.StatusOK, res)
	//ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "Success get all data request", "data": requestes, "pagination": pagination})
}
