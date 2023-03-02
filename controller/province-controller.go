package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/helper"
	"github.com/yusufbagussh/pet_hotel_backend/service"
	"net/http"
	"strconv"
)

// ProvinceController is a contract of what provinceController can do
type ProvinceController interface {
	GetAllProvince(ctx *gin.Context)
	CreateProvince(ctx *gin.Context)
	DeleteProvince(ctx *gin.Context)
	ShowProvince(ctx *gin.Context)
	UpdateProvince(ctx *gin.Context)
}

type provinceController struct {
	provinceService service.ProvinceService
	jwtService      service.JWTService
}

// NewProvinceController is creating anew instance of ProvinceControlller
func NewProvinceController(provinceService service.ProvinceService, jwtService service.JWTService) ProvinceController {
	return &provinceController{
		provinceService: provinceService,
		jwtService:      jwtService,
	}
}

func (u *provinceController) CreateProvince(ctx *gin.Context) {
	var CreateProvince dto.CreateProvince
	errDTO := ctx.ShouldBind(&CreateProvince)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, errCreate := u.provinceService.CreateProvince(CreateProvince)
	if errCreate != nil {
		res := helper.BuildErrorResponse("Failed to create province", errCreate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Create success", result)
	ctx.JSON(http.StatusOK, res)
}
func (u *provinceController) UpdateProvince(ctx *gin.Context) {
	var UpdateProvince dto.UpdateProvince
	errDTO := ctx.ShouldBind(&UpdateProvince)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, errCreate := u.provinceService.UpdateProvince(UpdateProvince)
	if errCreate != nil {
		res := helper.BuildErrorResponse("Failed to update province", errCreate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Update success", result)
	ctx.JSON(http.StatusOK, res)
}
func (u *provinceController) DeleteProvince(ctx *gin.Context) {
	provinceID := ctx.Param("id")
	errDel := u.provinceService.DeleteProvince(provinceID)
	if errDel != nil {
		res := helper.BuildErrorResponse("Failed to delete", errDel.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Delete success", nil)
	ctx.JSON(http.StatusOK, res)
}
func (u *provinceController) ShowProvince(ctx *gin.Context) {
	provinceID := ctx.Param("id")
	result, errShow := u.provinceService.ShowProvince(provinceID)
	if errShow != nil {
		res := helper.BuildErrorResponse("Failed to show data", errShow.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Show data success", result)
	ctx.JSON(http.StatusOK, res)
}

func (u *provinceController) GetAllProvince(ctx *gin.Context) {
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

	var provincees, pagination, err = u.provinceService.GetAllProvince(filterPagination)
	//type provincePage struct {
	//	Provincees    []entity.Province
	//	Pagination dto.Pagination
	//}
	//provinceesPage := provincePage{
	//	provincees,
	//	pagination,
	//}
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get all data", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponsePage(true, "Get all data success", provincees, pagination)
	ctx.JSON(http.StatusOK, res)
	//ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "Success get all data province", "data": provincees, "pagination": pagination})
}
