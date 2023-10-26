package customer

import (
	"aszaychik/smartcafe-api/internal/interfaces"

	"github.com/labstack/echo/v4"
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

func (ar *CustomerRoutesImpl) Customer() {
	customersGroup := ar.Echo.Group("customers")

	customersGroup.POST("", ar.CustomerHandler.CreateCustomerHandler)
	customersGroup.GET("", ar.CustomerHandler.GetCustomersHandler)
	customersGroup.GET("/:id", ar.CustomerHandler.GetCustomerHandler)
	customersGroup.PUT("/:id", ar.CustomerHandler.UpdateCustomerHandler)
	customersGroup.DELETE("/:id", ar.CustomerHandler.DeleteCustomerHandler)
}