package interfaces

import (
	"aszaychik/smartcafe-api/domain"
	"aszaychik/smartcafe-api/domain/web"

	"github.com/labstack/echo/v4"
)

type CustomerRepository interface {
	Save(customer *domain.Customer) (*domain.Customer, error)
	Update(customer *domain.Customer, id int) (*domain.Customer, error)
	FindById(id int) (*domain.Customer, error)
	FindByName(itemName string) (*domain.Customer, error)
	FindByEmail(email string) (*domain.Customer, error)
	FindAll() ([]domain.Customer, error)
	Delete(id int) error
}

type CustomerService interface {
	CreateCustomer(ctx echo.Context, request web.CustomerCreateRequest) (*domain.Customer, error)
	UpdateCustomer(ctx echo.Context, request web.CustomerUpdateRequest, id int) (*domain.Customer, error)
	FindById(ctx echo.Context, id int) (*domain.Customer, error)
	FindAll(ctx echo.Context) ([]domain.Customer, error)
	DeleteCustomer(ctx echo.Context, id int) error
}

type CustomerHandler interface {
	CreateCustomerHandler(ctx echo.Context) error
	UpdateCustomerHandler(ctx echo.Context) error
	GetCustomerHandler(ctx echo.Context) error
	GetCustomersHandler(ctx echo.Context) error
	DeleteCustomerHandler(ctx echo.Context) error
}

type CustomerRoutes interface {
	Customer()
}