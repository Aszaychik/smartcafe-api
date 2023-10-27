package conversion

import (
	"aszaychik/smartcafe-api/domain"
	"aszaychik/smartcafe-api/domain/web"
	"time"
)

func AdminCreateRequestToAdminDomain(request web.AdminCreateRequest) *domain.Admin {
	return &domain.Admin{
		Username:    request.Username,
		Password: request.Password,
	}
}

func AdminLoginRequestToAdminDomain(request web.AdminLoginRequest) *domain.Admin {
	return &domain.Admin{
		Username:    request.Username,
		Password: request.Password,
	}
}

func AdminUpdateRequestToAdminDomain(request web.AdminUpdateRequest) *domain.Admin {
	return &domain.Admin{
		Username:    request.Username,
		Password: request.Password,
	}
}

func MenuCreateRequestToMenuDomain(request web.MenuCreateRequest) *domain.Menu {
	return &domain.Menu{
		ItemName: request.ItemName,
		ItemPrice: request.ItemPrice,
		ItemDescription: request.ItemDescription,
		ItemImage: request.ItemImage,
		CategoryID: uint(request.CategoryId),
	}
}

func MenuUpdateRequestToMenuDomain(request web.MenuUpdateRequest) *domain.Menu {
	return &domain.Menu{
		ItemName: request.ItemName,
		ItemPrice: request.ItemPrice,
		ItemDescription: request.ItemDescription,
		ItemImage: request.ItemImage,
		CategoryID: uint(request.CategoryId),
	}
}

func CategoryCreateRequestToCategoryDomain(request web.CategoryCreateRequest) *domain.Category {
	return &domain.Category{
		CategoryName: request.CategoryName,
		CategoryDescription: request.CategoryDescription,
	}
}

func CategoryUpdateRequestToCategoryDomain(request web.CategoryUpdateRequest) *domain.Category {
	return &domain.Category{
		CategoryName: request.CategoryName,
		CategoryDescription: request.CategoryDescription,
	}
}

func CustomerCreateRequestToCustomerDomain(request web.CustomerCreateRequest) *domain.Customer {
	return &domain.Customer{
		CustomerName: request.CustomerName,
		CustomerEmail: request.CustomerEmail,
	}
}

func CustomerUpdateRequestToCustomerDomain(request web.CustomerUpdateRequest) *domain.Customer {
	return &domain.Customer{
		CustomerName: request.CustomerName,
		CustomerEmail: request.CustomerEmail,
	}
}

func OrderCreateRequestToOrderDomain(request web.OrderCreateRequest, totalPrice float64) *domain.Order {
	order := &domain.Order{
		CustomerID:  uint(request.CustomerId),
		OrderDate:   time.Now(),
		OrderStatus: domain.Pending,
		TotalPrice:  totalPrice,
	}

	for _, itemRequest := range request.Items {
		orderItem := &domain.OrderItem{
			ItemID:    uint(itemRequest.ItemID),
			Quantity:  itemRequest.Quantity,
		}

		order.Items = append(order.Items, *orderItem)
	}

	return order
}

func OrderCreateDomainToOrderPaymentDomain(order *domain.Order, id string, paymentUrl string) *domain.Order {
	order.OrderPayment = domain.OrderPayment{
		ID: id,
		PaymentUrl: paymentUrl,
		PaymentDate: time.Now(),
	}

	return order
}