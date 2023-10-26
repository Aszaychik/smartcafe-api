package menu

import (
	"aszaychik/smartcafe-api/internal/interfaces"

	"github.com/labstack/echo/v4"
)

type MenuRoutesImpl struct {
	Echo        *echo.Echo
	MenuHandler interfaces.MenuHandler
}

func NewMenuRoutes(e *echo.Echo, menuHandler interfaces.MenuHandler) interfaces.MenuRoutes {
	return &MenuRoutesImpl{
		Echo:        e,
		MenuHandler: menuHandler,
	}
}

func (ar *MenuRoutesImpl) Menu() {
	menusGroup := ar.Echo.Group("menus")

	menusGroup.POST("", ar.MenuHandler.CreateMenuHandler)
	menusGroup.GET("", ar.MenuHandler.GetMenusHandler)
	menusGroup.GET("/:id", ar.MenuHandler.GetMenuHandler)
	menusGroup.PUT("/:id", ar.MenuHandler.UpdateMenuHandler)
	menusGroup.DELETE("/:id", ar.MenuHandler.DeleteMenuHandler)
}