package web

type AdminCreateRequest struct {
	Username string `json:"username" validate:"required,min=1,max=255"`
	Password string `json:"password" validate:"required,min=8,max=255"`
}

type AdminLoginRequest struct {
	Username string `json:"username" validate:"required,min=1,max=255"`
	Password string `json:"password" validate:"required,max=255"`
}

type AdminUpdateRequest struct {
	Username string `json:"username" validate:"min=1,max=255"`
	Password string `json:"password" validate:"min=8,max=255"`
}

type CategoryCreateRequest struct {
	CategoryName        string `json:"category_name" validate:"required,min=1,max=255"`
	CategoryDescription string `json:"category_description" validate:"min=1,max=255"`
}

type CategoryUpdateRequest struct {
	CategoryName        string `json:"category_name" validate:"min=1,max=255"`
	CategoryDescription string `json:"category_description" validate:"min=1,max=255"`
}

type MenuCreateRequest struct {
	ItemName        string `json:"item_name" validate:"required,min=1,max=255"`
	ItemPrice       int    `json:"item_price" validate:"required"`
	ItemDescription string `json:"item_description" validate:"min=1,max=255"`
	ItemImage       string `json:"item_image" validate:"min=1,max=255"`
	CategoryId      int    `json:"category_id" validate:"required,min=1,max=255"`
}

type MenuUpdateRequest struct {
	ItemName        string `json:"item_name" validate:"min=1,max=255"`
	ItemPrice       int    `json:"item_price"`
	ItemDescription string `json:"item_description" validate:"min=1,max=255"`
	ItemImage       string `json:"item_image" validate:"min=1,max=255"`
	CategoryId      int    `json:"category_id" validate:"min=1,max=255"`
}

type CustomerCreateRequest struct {
	CustomerName  string `json:"customer_name" validate:"required,min=1,max=255"`
	CustomerEmail string `json:"customer_email" validate:"required,email"`
}

type CustomerUpdateRequest struct {
	CustomerName  string `json:"customer_name" validate:"min=1,max=255"`
	CustomerEmail string `json:"customer_email" validate:"email"`
}