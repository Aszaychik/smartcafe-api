package customer

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

type CustomerServiceImpl struct {
	CustomerRepository interfaces.CustomerRepository
	Validate       *validator.Validate
}

func NewCustomerService(customerRepository interfaces.CustomerRepository, validate *validator.Validate) interfaces.CustomerService {
	return &CustomerServiceImpl{
		CustomerRepository: customerRepository,
		Validate:        validate,
	}
}

func (service *CustomerServiceImpl) CreateCustomer(ctx echo.Context, request web.CustomerCreateRequest) (*domain.Customer, error) {
	// Check if the request is valid
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, validation.ValidationError(ctx, err)
	}

	// Check if the Customer email already exists
	existingCustomer, _ := service.CustomerRepository.FindByEmail(request.CustomerEmail)
	if existingCustomer != nil {
		return nil, fmt.Errorf("Customer Email already exists")
	}

	// Convert request to domain
	customer := conversion.CustomerCreateRequestToCustomerDomain(request)

	result, err := service.CustomerRepository.Save(customer)
	if err != nil {
		return nil, fmt.Errorf("Error when create : %s", err.Error())
	}

	return result, nil
}

func (service *CustomerServiceImpl) UpdateCustomer(ctx echo.Context, request web.CustomerUpdateRequest, id int) (*domain.Customer, error) {
	// Check if the request is valid
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, validation.ValidationError(ctx, err)
	}

	// Check if the customer exists
	existingCustomer, _ := service.CustomerRepository.FindById(id)
	if existingCustomer == nil {
		return nil, fmt.Errorf("Customer not found")
	}

	// Convert request to domain
	customer := conversion.CustomerUpdateRequestToCustomerDomain(request)

	result, err := service.CustomerRepository.Update(customer, id)
	if err != nil {
		return nil, fmt.Errorf("Error when updating : %s", err.Error())
	}

	return result, nil
}

func (service *CustomerServiceImpl) FindById(ctx echo.Context, id int) (*domain.Customer, error) {
	// Check if the customer exists
	existingCustomer, _ := service.CustomerRepository.FindById(id)
	if existingCustomer == nil {
		return nil, fmt.Errorf("Customer not found")
	}

	return existingCustomer, nil
}

func (service *CustomerServiceImpl) FindAll(ctx echo.Context) ([]domain.Customer, error) {
	customers, err := service.CustomerRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("Customers not found")
	}

	return customers, nil
}

func (service *CustomerServiceImpl) DeleteCustomer(ctx echo.Context, id int) error {
	// Check if the customer exists
	existingCustomer, _ := service.CustomerRepository.FindById(id)
	if existingCustomer == nil {
		return fmt.Errorf("Customer not found")
	}

	err := service.CustomerRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("Error when deleting : %s", err)
	}

	return nil
}