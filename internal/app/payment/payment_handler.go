package payment

import (
	"aszaychik/smartcafe-api/internal/interfaces"
	"aszaychik/smartcafe-api/pkg/res"
	"encoding/json"

	"github.com/labstack/echo/v4"
)

type OrderPaymentHandlerImpl struct {
	service interfaces.OrderPaymentService
}

func NewOrderPaymentHandler(service interfaces.OrderPaymentService) interfaces.OrderPaymentHandler {
	return &OrderPaymentHandlerImpl{
		service: service,
	}
}

func (handler *OrderPaymentHandlerImpl) FetchNotifications(ctx echo.Context) error {
	var notificationPayload map[string]any

	if err := json.NewDecoder(ctx.Request().Body).Decode(&notificationPayload); err != nil {
		return res.StatusBadRequest(ctx, err)
	}

	err := handler.service.Notifications(notificationPayload)
	if err != nil {
		return res.StatusInternalServerError(ctx, err)
	}

	return res.StatusOK(ctx, "OK", nil)
}