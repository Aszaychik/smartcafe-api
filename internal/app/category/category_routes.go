package category

import (
	"aszaychik/smartcafe-api/internal/interfaces"

	"github.com/labstack/echo/v4"
)

type CategoryRoutesImpl struct {
	Echo        *echo.Echo
	CategoryHandler interfaces.CategoryHandler
}

func NewCategoryRoutes(e *echo.Echo, categoryHandler interfaces.CategoryHandler) interfaces.CategoryRoutes {
	return &CategoryRoutesImpl{
		Echo:        e,
		CategoryHandler: categoryHandler,
	}
}

func (ar *CategoryRoutesImpl) Category() {
	categoriesGroup := ar.Echo.Group("categories")

	categoriesGroup.POST("", ar.CategoryHandler.CreateCategoryHandler)
	categoriesGroup.GET("", ar.CategoryHandler.GetCategoriesHandler)
	categoriesGroup.GET("/:id", ar.CategoryHandler.GetCategoryHandler)
	categoriesGroup.PUT("/:id", ar.CategoryHandler.UpdateCategoryHandler)
	categoriesGroup.DELETE("/:id", ar.CategoryHandler.DeleteCategoryHandler)
}