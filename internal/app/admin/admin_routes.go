package admin

import (
	"aszaychik/smartcafe-api/internal/interfaces"

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

func (ar *AdminRoutesImpl) Auth() {
	authGroup := ar.Echo.Group("auth")

	authGroup.POST("/register", ar.AdminHandler.RegisterAdminHandler)
	authGroup.POST("/login", ar.AdminHandler.LoginAdminHandler)
}

func (ar *AdminRoutesImpl) Admin() {
	adminsGroup := ar.Echo.Group("admins")

	adminsGroup.GET("", ar.AdminHandler.GetAdminsHandler)
	adminsGroup.GET("/:id", ar.AdminHandler.GetAdminHandler)
	adminsGroup.PUT("/:id", ar.AdminHandler.UpdateAdminHandler)
	adminsGroup.DELETE("/:id", ar.AdminHandler.DeleteAdminHandler)
}