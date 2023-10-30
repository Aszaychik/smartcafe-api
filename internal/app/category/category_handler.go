package category

import (
	"aszaychik/smartcafe-api/domain/web"
	"aszaychik/smartcafe-api/internal/interfaces"
	"aszaychik/smartcafe-api/pkg/res"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type CategoryHandlerImpl struct {
	CategoryService interfaces.CategoryService
}

func NewCategoryHandler(categoryService interfaces.CategoryService) interfaces.CategoryHandler {
	return &CategoryHandlerImpl{CategoryService: categoryService}
}

func(handler *CategoryHandlerImpl) CreateCategoryHandler(ctx echo.Context) error {
	categoryCreateRequest := web.CategoryCreateRequest{}
	err := ctx.Bind(&categoryCreateRequest) 
	if err != nil {
		return res.StatusBadRequest(ctx, err)
	}

	response, err := handler.CategoryService.CreateCategory(ctx, categoryCreateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return res.StatusBadRequest(ctx, err)
		}

		if strings.Contains(err.Error(), "already exists") {
			return res.StatusAlreadyExist(ctx, err)
		}

		return res.StatusInternalServerError(ctx, err)
	}

	return res.StatusCreated(ctx, "Success to create category", response)
}


func (handler *CategoryHandlerImpl) UpdateCategoryHandler(ctx echo.Context) error {
	categoryId := ctx.Param("id")
	categoryIdInt, err := strconv.Atoi(categoryId)
	if err != nil {
		return res.StatusInternalServerError(ctx, err)
	}

	categoryUpdateRequest := web.CategoryUpdateRequest{}
	err = ctx.Bind(&categoryUpdateRequest)
	if err != nil {
		return res.StatusBadRequest(ctx, err)
	}

	response, err := handler.CategoryService.UpdateCategory(ctx, categoryUpdateRequest, categoryIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return res.StatusBadRequest(ctx, err)
		}

		if strings.Contains(err.Error(), "already exists") {
			return res.StatusAlreadyExist(ctx, err)
		}
		
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, err)
		}

		return res.StatusInternalServerError(ctx, err)
	}

	return res.StatusOK(ctx, "Success to update category", response)
}

func (handler *CategoryHandlerImpl) GetCategoryHandler(ctx echo.Context) error {
	categoryId := ctx.Param("id")
	categoryIdInt, err := strconv.Atoi(categoryId)
	if err != nil {
		return res.StatusInternalServerError(ctx, err)
	}
	
	response, err := handler.CategoryService.FindById(ctx, categoryIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, err)
		}

		return res.StatusInternalServerError(ctx, err)
	}

	return res.StatusOK(ctx, "Success to get category", response)
}

func (handler *CategoryHandlerImpl) GetCategoriesHandler(ctx echo.Context) error {
	response, err := handler.CategoryService.FindAll(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, err)
		}

		return res.StatusInternalServerError(ctx, err)
	}

	return res.StatusOK(ctx, "Success to get categories", response)
}

func (handler *CategoryHandlerImpl) DeleteCategoryHandler(ctx echo.Context) error {
	categoryId := ctx.Param("id")
	categoryIdInt, err := strconv.Atoi(categoryId)
	if err != nil {
		return res.StatusInternalServerError(ctx, err)
	}

	err = handler.CategoryService.DeleteCategory(ctx, categoryIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, err)
		}

		return res.StatusInternalServerError(ctx, err)
	}

	return res.StatusOK(ctx, "Success to delete category", nil)
}