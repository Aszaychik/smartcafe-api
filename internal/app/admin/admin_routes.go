package admin

import (
	"aszaychik/smartcafe-api/config"
	"aszaychik/smartcafe-api/internal/interfaces"
	"aszaychik/smartcafe-api/internal/middleware"

	"github.com/labstack/echo/v4"
)

type AdminRoutesImpl struct {
	Echo           *echo.Echo
	AdminHandler interfaces.AdminHandler
}

func NewAdminRoutes(e *echo.Echo, adminHandler interfaces.AdminHandler) interfaces.AdminRoutes {
	return &AdminRoutesImpl{
		Echo:           e,
		AdminHandler: adminHandler,
	}
}

func (ar *AdminRoutesImpl) Auth(config *config.AuthConfig) {
	authGroup := ar.Echo.Group("auth")

	authGroup.POST("/login", ar.AdminHandler.LoginAdminHandler)
	authGroup.POST("/register", ar.AdminHandler.RegisterAdminHandler, middleware.RegisterAuth(config))
}

func (ar *AdminRoutesImpl) Admin() {
	adminsGroup := ar.Echo.Group("admins")

	adminsGroup.GET("", ar.AdminHandler.GetAdminsHandler)
	adminsGroup.GET("/:id", ar.AdminHandler.GetAdminHandler)
	adminsGroup.PUT("/:id", ar.AdminHandler.UpdateAdminHandler)
	adminsGroup.DELETE("/:id", ar.AdminHandler.DeleteAdminHandler)
}