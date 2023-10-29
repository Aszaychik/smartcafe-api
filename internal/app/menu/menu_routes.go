package menu

import (
	"aszaychik/smartcafe-api/config"
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

func (mr *MenuRoutesImpl) Menu(config *config.AuthConfig) {
	menusGroup := mr.Echo.Group("menus")

	menusGroup.POST("", mr.MenuHandler.CreateMenuHandler)
	menusGroup.GET("", mr.MenuHandler.GetMenusHandler)
	menusGroup.GET("/:id", mr.MenuHandler.GetMenuHandler)
	menusGroup.PUT("/:id", mr.MenuHandler.UpdateMenuHandler)
	menusGroup.PATCH("/:id", mr.MenuHandler.UploadImageMenuHandler)
	menusGroup.DELETE("/:id", mr.MenuHandler.DeleteMenuHandler)
}