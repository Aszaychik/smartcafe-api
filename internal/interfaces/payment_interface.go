package interfaces

import (
	"aszaychik/smartcafe-api/domain"

	"github.com/labstack/echo/v4"
)

type OrderPaymentRepository interface {
	UpdateOrderPaymentStatus(paymentId string, orderPaymentStatus domain.PaymentStatus) error
	FindById(orderId string) (*domain.OrderPayment, error)
}

type OrderPaymentService interface {
	Notifications(notificationPayload map[string]any) error
}

type OrderPaymentHandler interface {
	FetchNotifications(ctx echo.Context) error
}

type OrderPaymentRoutes interface {
	OrderPayment()
}