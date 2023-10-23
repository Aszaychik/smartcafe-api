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