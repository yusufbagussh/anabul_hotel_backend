package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yusufbagussh/pet_hotel_backend/dto"
	"github.com/yusufbagussh/pet_hotel_backend/helper"
	"github.com/yusufbagussh/pet_hotel_backend/service"
	"net/http"
)

// CategoryController is a contract of what categoryController can do
type CategoryController interface {
	GetAllCategory(ctx *gin.Context)
	CreateCategory(ctx *gin.Context)
	DeleteCategory(ctx *gin.Context)
	ShowCategory(ctx *gin.Context)
	UpdateCategory(ctx *gin.Context)
}

type categoryController struct {
	categoryService service.CategoryService
	jwtService      service.JWTService
}

// NewCategoryController is creating anew instance of CategoryControlller
func NewCategoryController(categoryService service.CategoryService, jwtService service.JWTService) CategoryController {
	return &categoryController{
		categoryService: categoryService,
		jwtService:      jwtService,
	}
}

func (u *categoryController) CreateCategory(ctx *gin.Context) {
	var CreateCategory dto.CreateCategory
	errDTO := ctx.ShouldBind(&CreateCategory)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, errCreate := u.categoryService.CreateCategory(CreateCategory)
	if errCreate != nil {
		res := helper.BuildErrorResponse("Failed to create category", errCreate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Create success", result)
	ctx.JSON(http.StatusOK, res)
}
func (u *categoryController) UpdateCategory(ctx *gin.Context) {
	var UpdateCategory dto.UpdateCategory
	errDTO := ctx.ShouldBind(&UpdateCategory)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, errCreate := u.categoryService.UpdateCategory(UpdateCategory)
	if errCreate != nil {
		res := helper.BuildErrorResponse("Failed to update category", errCreate.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Update success", result)
	ctx.JSON(http.StatusOK, res)
}
func (u *categoryController) DeleteCategory(ctx *gin.Context) {
	categoryID := ctx.Param("id")
	errDel := u.categoryService.DeleteCategory(categoryID)
	if errDel != nil {
		res := helper.BuildErrorResponse("Failed to delete", errDel.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Delete success", nil)
	ctx.JSON(http.StatusOK, res)
}
func (u *categoryController) ShowCategory(ctx *gin.Context) {
	categoryID := ctx.Param("id")
	result, errShow := u.categoryService.ShowCategory(categoryID)
	if errShow != nil {
		res := helper.BuildErrorResponse("Failed to show data", errShow.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "Show data success", result)
	ctx.JSON(http.StatusOK, res)
}

func (u *categoryController) GetAllCategory(ctx *gin.Context) {
	var filterPagination dto.FilterPagination

	errDTO := ctx.ShouldBind(&filterPagination)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var admins, pagination, err = u.categoryService.GetAllCategory(filterPagination)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get all data", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponsePage(true, "Get all data success", admins, pagination)
	ctx.JSON(http.StatusOK, res)
}
