package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/helper"
	"github.com/yusufbagussh/pet_hotel_backend/service"
	"net/http"
)

// ClassController is a contract of what classController can do
type ClassController interface {
	GetAllClass(ctx *gin.Context)
	CreateClass(ctx *gin.Context)
	DeleteClass(ctx *gin.Context)
	ShowClass(ctx *gin.Context)
	UpdateClass(ctx *gin.Context)
}

type classController struct {
	classService service.ClassService
	jwtService   service.JWTService
}

// NewClassController is creating anew instance of ClassControlller
func NewClassController(classService service.ClassService, jwtService service.JWTService) ClassController {
	return &classController{
		classService: classService,
		jwtService:   jwtService,
	}
}

func (u *classController) CreateClass(ctx *gin.Context) {
	var CreateClass dto.CreateClass
	errDTO := ctx.ShouldBind(&CreateClass)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, errCreate := u.classService.CreateClass(CreateClass)
	if errCreate != nil {
		res := helper.BuildErrorResponse("Failed to create class", errCreate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Create success", result)
	ctx.JSON(http.StatusOK, res)
}
func (u *classController) UpdateClass(ctx *gin.Context) {
	var UpdateClass dto.UpdateClass
	errDTO := ctx.ShouldBind(&UpdateClass)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, errCreate := u.classService.UpdateClass(UpdateClass)
	if errCreate != nil {
		res := helper.BuildErrorResponse("Failed to update class", errCreate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Update success", result)
	ctx.JSON(http.StatusOK, res)
}
func (u *classController) DeleteClass(ctx *gin.Context) {
	classID := ctx.Param("id")
	errDel := u.classService.DeleteClass(classID)
	if errDel != nil {
		res := helper.BuildErrorResponse("Failed to delete", errDel.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Delete success", nil)
	ctx.JSON(http.StatusOK, res)
}
func (u *classController) ShowClass(ctx *gin.Context) {
	classID := ctx.Param("id")
	result, errShow := u.classService.ShowClass(classID)
	if errShow != nil {
		res := helper.BuildErrorResponse("Failed to show data", errShow.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Show data success", result)
	ctx.JSON(http.StatusOK, res)
}

func (u *classController) GetAllClass(ctx *gin.Context) {
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

	var classes, pagination, err = u.classService.GetAllClass(filterPagination)
	//type classPage struct {
	//	Classes    []entity.Class
	//	Pagination dto.Pagination
	//}
	//classesPage := classPage{
	//	classes,
	//	pagination,
	//}
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get all data", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponsePage(true, "Get all data success", classes, pagination)
	ctx.JSON(http.StatusOK, res)
	//ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "Success get all data class", "data": classes, "pagination": pagination})
}
