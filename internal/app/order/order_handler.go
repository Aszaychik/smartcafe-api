package order

import (
	"aszaychik/smartcafe-api/domain/web"
	"aszaychik/smartcafe-api/internal/interfaces"
	"aszaychik/smartcafe-api/pkg/res"
	"strings"

	"github.com/labstack/echo/v4"
)

type OrderHandlerImpl struct {
	OrderService interfaces.OrderService
}

func NewOrderHandler(menuService interfaces.OrderService) interfaces.OrderHandler {
	return &OrderHandlerImpl{OrderService: menuService}
}

func (handler *OrderHandlerImpl) CreateOrderHandler(ctx echo.Context) error {
	menuCreateRequest := web.OrderCreateRequest{}
	err := ctx.Bind(&menuCreateRequest)
	if err != nil {
		return res.StatusBadRequest(ctx, err)
	}

	response, err := handler.OrderService.CreateOrder(ctx, menuCreateRequest)
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