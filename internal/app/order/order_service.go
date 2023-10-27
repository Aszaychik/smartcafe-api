package order

import (
	"aszaychik/smartcafe-api/domain"
	"aszaychik/smartcafe-api/domain/web"
	"aszaychik/smartcafe-api/internal/interfaces"
	"aszaychik/smartcafe-api/pkg/conversion"
	"aszaychik/smartcafe-api/pkg/validation"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type OrderServiceImpl struct {
	OrderRepository interfaces.OrderRepository
	MenuRepository  interfaces.MenuRepository
	CustomerRepository interfaces.CustomerRepository
	Validate       *validator.Validate
}

func NewOrderService(orderRepository interfaces.OrderRepository, menuRepository interfaces.MenuRepository, customerRepository interfaces.CustomerRepository, validate *validator.Validate) interfaces.OrderService {
	return &OrderServiceImpl{
		OrderRepository: orderRepository,
		MenuRepository:  menuRepository,
		CustomerRepository: customerRepository,
		Validate: validate,
	}
}

func (service *OrderServiceImpl) CreateOrder(ctx echo.Context, request web.OrderCreateRequest) (*domain.Order, error) {
	// Check if the request is valid
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, validation.ValidationError(ctx, err)
	}

	customer, err := service.CustomerRepository.FindById(request.CustomerId)
	if err != nil {
		return nil, fmt.Errorf("failed to find customer: %w", err)
	}

	if customer == nil {
		return nil, fmt.Errorf("customer not found")
	}

	// Calculate total price and create the order domain object
	totalPrice, err := service.CalculateTotalPrice(request.Items)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate total price: %w", err)
	}

	order := conversion.OrderCreateRequestToOrderDomain(request, totalPrice)

	// Save the order to the database
	result, err := service.OrderRepository.Save(order)
	if err != nil {
		return nil, fmt.Errorf("failed to save order: %w", err)
	}

	// Fetch the saved order from the database
	result, err = service.OrderRepository.FindById(int(result.ID))
	if err != nil {
		return nil, fmt.Errorf("failed to find order: %w", err)
	}

	return result, nil
}

func (service *OrderServiceImpl) CalculateTotalPrice(items []domain.OrderItem) (float64, error) {
	var totalPrice float64

	for _, item := range items {
		// Fetch the item from the database to get the latest price
		menu, err := service.MenuRepository.FindById(int(item.ItemID))
		if err != nil {
			return 0, fmt.Errorf("failed to find menu: %w", err)
		}

		totalPrice += float64(item.Quantity) * menu.ItemPrice
	}

	return totalPrice, nil
}