package web

type CategoryCreateRequest struct {
	CategoryName        string `json:"category_name" validate:"required,min=1,max=255"`
	CategoryDescription string `json:"category_description" validate:"min=1,max=255"`
}

type CategoryUpdateRequest struct {
	CategoryName        string `json:"category_name" validate:"min=1,max=255"`
	CategoryDescription string `json:"category_description" validate:"min=1,max=255"`
}