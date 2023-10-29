package order

import (
	"aszaychik/smartcafe-api/internal/interfaces"

	"github.com/labstack/echo/v4"
)

type OrderRoutesImpl struct {
	Echo        *echo.Echo
	OrderHandler interfaces.OrderHandler
}

func NewOrderRoutes(e *echo.Echo, orderHandler interfaces.OrderHandler) interfaces.OrderRoutes {
	return &OrderRoutesImpl{
		Echo:        e,
		OrderHandler: orderHandler,
	}
}

func (ar *OrderRoutesImpl) Order() {
	ordersGroup := ar.Echo.Group("orders")

	ordersGroup.POST("", ar.OrderHandler.CreateOrderHandler)
	ordersGroup.GET("/:id", ar.OrderHandler.GetOrderHandler)
}