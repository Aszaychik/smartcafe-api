package menu

import (
	"aszaychik/smartcafe-api/domain/web"
	"aszaychik/smartcafe-api/internal/interfaces"
	"aszaychik/smartcafe-api/pkg/res"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type MenuHandlerImpl struct {
	MenuService interfaces.MenuService
}

func NewMenuHandler(menuService interfaces.MenuService) interfaces.MenuHandler {
	return &MenuHandlerImpl{MenuService: menuService}
}

func(handler *MenuHandlerImpl) CreateMenuHandler(ctx echo.Context) error {
	menuCreateRequest := web.MenuCreateRequest{}
	err := ctx.Bind(&menuCreateRequest) 
	if err != nil {
		return res.StatusBadRequest(ctx, err)
	}

	response, err := handler.MenuService.CreateMenu(ctx, menuCreateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return res.StatusBadRequest(ctx, err)
		}

		if strings.Contains(err.Error(), "already exists") {
			return res.StatusAlreadyExist(ctx, err)
		}

		return res.StatusInternalServerError(ctx, err)
	}

	return res.StatusCreated(ctx, "Success to create menu", response)
}


func (handler *MenuHandlerImpl) UpdateMenuHandler(ctx echo.Context) error {
	menuId := ctx.Param("id")
	menuIdInt, err := strconv.Atoi(menuId)
	if err != nil {
		return res.StatusInternalServerError(ctx, err)
	}

	menuUpdateRequest := web.MenuUpdateRequest{}
	err = ctx.Bind(&menuUpdateRequest)
	if err != nil {
		return res.StatusBadRequest(ctx, err)
	}

	response, err := handler.MenuService.UpdateMenu(ctx, menuUpdateRequest, menuIdInt)
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

	return res.StatusOK(ctx, "Success to update menu", response)
}

func (handler *MenuHandlerImpl) UploadImageMenuHandler(ctx echo.Context) error {
	menuId := ctx.Param("id")
	menuIdInt, err := strconv.Atoi(menuId)
	if err != nil {
		return res.StatusInternalServerError(ctx, err)
	}

	itemImage, err := ctx.FormFile("item_image")
	if err != nil {
		return res.StatusInternalServerError(ctx, err)
	}

	response, err := handler.MenuService.UpdateImageMenu(ctx, itemImage, menuIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return res.StatusBadRequest(ctx, err)
		}

		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, err)
		}

		return res.StatusInternalServerError(ctx, err)
	}

	return res.StatusOK(ctx, "Success to update menu", response)
}

func (handler *MenuHandlerImpl) GetMenuHandler(ctx echo.Context) error {
	menuId := ctx.Param("id")
	menuIdInt, err := strconv.Atoi(menuId)
	if err != nil {
		return res.StatusInternalServerError(ctx, err)
	}
	
	response, err := handler.MenuService.FindById(ctx, menuIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, err)
		}

		return res.StatusInternalServerError(ctx, err)
	}

	return res.StatusOK(ctx, "Success to get menu", response)
}

func (handler *MenuHandlerImpl) GetMenusHandler(ctx echo.Context) error {
	response, err := handler.MenuService.FindAll(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, err)
		}

		return res.StatusInternalServerError(ctx, err)
	}

	return res.StatusOK(ctx, "Success to get menus", response)
}

func (handler *MenuHandlerImpl) DeleteMenuHandler(ctx echo.Context) error {
	menuId := ctx.Param("id")
	menuIdInt, err := strconv.Atoi(menuId)
	if err != nil {
		return res.StatusInternalServerError(ctx, err)
	}

	err = handler.MenuService.DeleteMenu(ctx, menuIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, err)
		}

		return res.StatusInternalServerError(ctx, err)
	}

	return res.StatusOK(ctx, "Success to delete menu", nil)
}