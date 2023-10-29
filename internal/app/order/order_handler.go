package order

import (
	"aszaychik/smartcafe-api/domain/web"
	"aszaychik/smartcafe-api/internal/interfaces"
	"aszaychik/smartcafe-api/pkg/res"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type OrderHandlerImpl struct {
	OrderService interfaces.OrderService
}

func NewOrderHandler(orderService interfaces.OrderService) interfaces.OrderHandler {
	return &OrderHandlerImpl{OrderService: orderService}
}

func (handler *OrderHandlerImpl) CreateOrderHandler(ctx echo.Context) error {
	orderCreateRequest := web.OrderCreateRequest{}
	err := ctx.Bind(&orderCreateRequest)
	if err != nil {
		return res.StatusBadRequest(ctx, err)
	}

	response, err := handler.OrderService.CreateOrder(ctx, orderCreateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return res.StatusBadRequest(ctx, err)
		}

		if strings.Contains(err.Error(), "already exists") {
			return res.StatusAlreadyExist(ctx, err)
		}

		return res.StatusInternalServerError(ctx, err)
	}

	return res.StatusCreated(ctx, "Success to create order", response)
}

func (handler *OrderHandlerImpl) GetOrderHandler(ctx echo.Context) error {
	orderId := ctx.Param("id")
	orderIdInt, err := strconv.Atoi(orderId)
	if err != nil {
		return res.StatusInternalServerError(ctx, err)
	}
	
	response, err := handler.OrderService.FindById(ctx, orderIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, err)
		}

		return res.StatusInternalServerError(ctx, err)
	}

	return res.StatusOK(ctx, "Success to get order", response)
}
