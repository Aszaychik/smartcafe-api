package customer

import (
	"aszaychik/smartcafe-api/config"
	"aszaychik/smartcafe-api/internal/interfaces"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomerRoutesImpl struct {
	Echo        *echo.Echo
	CustomerHandler interfaces.CustomerHandler
}

func NewCustomerRoutes(e *echo.Echo, customerHandler interfaces.CustomerHandler) interfaces.CustomerRoutes {
	return &CustomerRoutesImpl{
		Echo:        e,
		CustomerHandler: customerHandler,
	}
}

func (ar *CustomerRoutesImpl) Customer(config *config.AuthConfig) {
	customersGroup := ar.Echo.Group("customers")

	customersGroup.POST("", ar.CustomerHandler.CreateCustomerHandler)
	
	customersGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte([]byte(config.JWTSecret)),
	}))
	customersGroup.GET("", ar.CustomerHandler.GetCustomersHandler)
	customersGroup.GET("/:id", ar.CustomerHandler.GetCustomerHandler)
	customersGroup.PUT("/:id", ar.CustomerHandler.UpdateCustomerHandler)
	customersGroup.DELETE("/:id", ar.CustomerHandler.DeleteCustomerHandler)
}