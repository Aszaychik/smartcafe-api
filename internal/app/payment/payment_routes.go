package payment

import (
	"aszaychik/smartcafe-api/internal/interfaces"

	"github.com/labstack/echo/v4"
)

type OrderPaymentRoutesImpl struct {
	Echo        *echo.Echo
	OrderPaymentHandler interfaces.OrderPaymentHandler
}

func NewOrderPaymentRoutes(e *echo.Echo, orderPaymentHandler interfaces.OrderPaymentHandler) interfaces.OrderPaymentRoutes {
	return &OrderPaymentRoutesImpl{
		Echo:        e,
		OrderPaymentHandler: orderPaymentHandler,
	}
}

func (ar *OrderPaymentRoutesImpl) OrderPayment() {
	orderPaymentsGroup := ar.Echo.Group("orders")

	orderPaymentsGroup.POST("/notifications", ar.OrderPaymentHandler.FetchNotifications)
}