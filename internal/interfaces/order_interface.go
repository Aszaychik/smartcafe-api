package interfaces

import (
	"aszaychik/smartcafe-api/domain"
	"aszaychik/smartcafe-api/domain/web"

	"github.com/labstack/echo/v4"
)

type OrderRepository interface {
	Save(order *domain.Order) (*domain.Order, error)
	FindById(orderId int) (*domain.Order, error)
}

type OrderService interface {
	CreateOrder(ctx echo.Context, request web.OrderCreateRequest) (*domain.Order, error)
}

type OrderHandler interface {
	CreateOrderHandler(ctx echo.Context) error
}

type OrderRoutes interface {
	Order()
}