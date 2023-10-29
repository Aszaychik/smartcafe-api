package customer

import (
	"aszaychik/smartcafe-api/domain/web"
	"aszaychik/smartcafe-api/internal/interfaces"
	"aszaychik/smartcafe-api/pkg/res"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type CustomerHandlerImpl struct {
	CustomerService interfaces.CustomerService
}

func NewCustomerHandler(customerService interfaces.CustomerService) interfaces.CustomerHandler {
	return &CustomerHandlerImpl{CustomerService: customerService}
}

func(handler *CustomerHandlerImpl) CreateCustomerHandler(ctx echo.Context) error {
	customerCreateRequest := web.CustomerCreateRequest{}
	err := ctx.Bind(&customerCreateRequest) 
	if err != nil {
		return res.StatusBadRequest(ctx, err)
	}

	response, err := handler.CustomerService.CreateCustomer(ctx, customerCreateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return res.StatusBadRequest(ctx, err)
		}

		if strings.Contains(err.Error(), "already exists") {
			return res.StatusAlreadyExist(ctx, err)
		}

		return res.StatusInternalServerError(ctx, err)
	}

	return res.StatusCreated(ctx, "Success to create customer", response)
}


func (handler *CustomerHandlerImpl) UpdateCustomerHandler(ctx echo.Context) error {
	customerId := ctx.Param("id")
	customerIdInt, err := strconv.Atoi(customerId)
	if err != nil {
		return res.StatusInternalServerError(ctx, err)
	}

	customerUpdateRequest := web.CustomerUpdateRequest{}
	err = ctx.Bind(&customerUpdateRequest)
	if err != nil {
		return res.StatusBadRequest(ctx, err)
	}

	response, err := handler.CustomerService.UpdateCustomer(ctx, customerUpdateRequest, customerIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return res.StatusBadRequest(ctx, err)
		}

		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, err)
		}

		return res.StatusInternalServerError(ctx, err)
	}

	return res.StatusOK(ctx, "Success to update customer", response)
}

func (handler *CustomerHandlerImpl) GetCustomerHandler(ctx echo.Context) error {
	customerId := ctx.Param("id")
	customerIdInt, err := strconv.Atoi(customerId)
	if err != nil {
		return res.StatusInternalServerError(ctx, err)
	}
	
	response, err := handler.CustomerService.FindById(ctx, customerIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, err)
		}

		return res.StatusInternalServerError(ctx, err)
	}

	return res.StatusOK(ctx, "Success to get customer", response)
}

func (handler *CustomerHandlerImpl) GetCustomersHandler(ctx echo.Context) error {
	response, err := handler.CustomerService.FindAll(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, err)
		}

		return res.StatusInternalServerError(ctx, err)
	}

	return res.StatusOK(ctx, "Success to get customers", response)
}

func (handler *CustomerHandlerImpl) DeleteCustomerHandler(ctx echo.Context) error {
	customerId := ctx.Param("id")
	customerIdInt, err := strconv.Atoi(customerId)
	if err != nil {
		return res.StatusInternalServerError(ctx, err)
	}

	err = handler.CustomerService.DeleteCustomer(ctx, customerIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, err)
		}

		return res.StatusInternalServerError(ctx, err)
	}

	return res.StatusOK(ctx, "Success to delete customer", nil)
}