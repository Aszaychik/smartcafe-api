package web

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