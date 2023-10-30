package web

import "aszaychik/smartcafe-api/domain"

type AdminCreateRequest struct {
	Username string `json:"username" validate:"required,min=1"`
	Password string `json:"password" validate:"required,min=8"`
}

type AdminLoginRequest struct {
	Username string `json:"username" validate:"required,min=1"`
	Password string `json:"password" validate:"required"`
}

type AdminUpdateRequest struct {
	Username string `json:"username"`
	Password string `json:"password" validate:"min=8"`
}

type CategoryCreateRequest struct {
	CategoryName        string `json:"category_name" validate:"required,min=1"`
	CategoryDescription string `json:"category_description" validate:"min=1"`
}

type CategoryUpdateRequest struct {
	CategoryName        string `json:"category_name"`
	CategoryDescription string `json:"category_description"`
}

type MenuCreateRequest struct {
	ItemName        string `json:"item_name" validate:"required,min=1"`
	ItemPrice       float64    `json:"item_price" validate:"required"`
	ItemDescription string `json:"item_description" validate:"min=1"`
	ItemImage       string `json:"item_image" validate:"min=1"`
	CategoryId      int    `json:"category_id" validate:"required,min=1"`
}

type MenuUpdateRequest struct {
	ItemName        string `json:"item_name"`
	ItemPrice       float64    `json:"item_price"`
	ItemDescription string `json:"item_description"`
	ItemImage       string `json:"item_image"`
	CategoryId      int    `json:"category_id"`
}

type CustomerCreateRequest struct {
	CustomerName  string `json:"customer_name" validate:"required,min=1"`
	CustomerEmail string `json:"customer_email" validate:"required,email"`
}

type CustomerUpdateRequest struct {
	CustomerName  string `json:"customer_name"`
	CustomerEmail string `json:"customer_email" validate:"email"`
}

type OrderItemRequest struct {
	ItemID   int `json:"item_id" validate:"required,min=1"`
	Quantity int `json:"quantity" validate:"required,min=1"`
}

type OrderCreateRequest struct {
	CustomerId int     `json:"customer_id" validate:"required,min=1"`
	SeatNumber int     `json:"seat_number" validate:"required,min=1"`
	Items      []domain.OrderItem `json:"items" validate:"required,min=1"`
}

type FeedbackCreateRequest struct {
	CustomerId int    `json:"customer_id" validate:"required,min=1"`
	OrderId    int    `json:"order_id" validate:"required,min=1"`
	FeedbackText string `json:"feedback_text" validate:"required,min=1"`
	FeedbackRating int    `json:"feedback_rating" validate:"required,min=1"`
}

type FeedbackUpdateRequest struct {
	FeedbackText string `json:"feedback_text"`
	FeedbackRating int    `json:"feedback_rating"`
}