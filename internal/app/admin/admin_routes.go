package admin

import (
	"aszaychik/smartcafe-api/internal/interfaces"

	"github.com/labstack/echo/v4"
)

type AdminRoutes struct {
	Echo           *echo.Echo
	AdminHandler interfaces.AdminHandler
}

func NewAdminRoutes(e *echo.Echo, adminHandler interfaces.AdminHandler) *AdminRoutes {
	return &AdminRoutes{
		Echo:           e,
		AdminHandler: adminHandler,
	}
}

func (ar *AdminRoutes) SetupAdminRoutes() {
	ar.auth()
	ar.admin()
}

func (ar *AdminRoutes) auth() {
	authGroup := ar.Echo.Group("auth")

	authGroup.POST("/register", ar.AdminHandler.RegisterAdminHandler)
	authGroup.POST("/login", ar.AdminHandler.LoginAdminHandler)
}

func (ar *AdminRoutes) admin() {
	adminsGroup := ar.Echo.Group("admins")

	adminsGroup.GET("", ar.AdminHandler.GetAdminsHandler)
	adminsGroup.GET("/:id", ar.AdminHandler.GetAdminHandler)
	adminsGroup.PUT("/:id", ar.AdminHandler.UpdateAdminHandler)
	adminsGroup.DELETE("/:id", ar.AdminHandler.DeleteAdminHandler)
}