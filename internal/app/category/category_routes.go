package category

import (
	"aszaychik/smartcafe-api/config"
	"aszaychik/smartcafe-api/internal/interfaces"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CategoryRoutesImpl struct {
	Echo        *echo.Echo
	CategoryHandler interfaces.CategoryHandler
	AuthConfig config.AuthConfig
}

func NewCategoryRoutes(e *echo.Echo, categoryHandler interfaces.CategoryHandler) interfaces.CategoryRoutes {
	return &CategoryRoutesImpl{
		Echo:        e,
		CategoryHandler: categoryHandler,
	}
}

func (ar *CategoryRoutesImpl) Category(config *config.AuthConfig) {
	categoriesGroup := ar.Echo.Group("categories")

	categoriesGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(config.XAPIKey),
	}))
	categoriesGroup.POST("", ar.CategoryHandler.CreateCategoryHandler)
	categoriesGroup.GET("", ar.CategoryHandler.GetCategoriesHandler)
	categoriesGroup.GET("/:id", ar.CategoryHandler.GetCategoryHandler)
	categoriesGroup.PUT("/:id", ar.CategoryHandler.UpdateCategoryHandler)
	categoriesGroup.DELETE("/:id", ar.CategoryHandler.DeleteCategoryHandler)
}