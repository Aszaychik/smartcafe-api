package conversion

import (
	"aszaychik/smartcafe-api/domain"
	"aszaychik/smartcafe-api/domain/web"
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