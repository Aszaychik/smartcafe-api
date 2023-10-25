package interfaces

import (
	"aszaychik/smartcafe-api/domain"
	"aszaychik/smartcafe-api/domain/web"

	"github.com/labstack/echo/v4"
)

type AdminRepository interface {
	Save(admin *domain.Admin) (*domain.Admin, error)
	Update(admin *domain.Admin, id int) (*domain.Admin, error)
	FindById(id int) (*domain.Admin, error)
	FindByUsername(username string) (*domain.Admin, error)
	FindAll() ([]domain.Admin, error)
	Delete(id int) error
}

type AdminService interface {
	RegisterAdmin(ctx echo.Context, request web.AdminCreateRequest) (*domain.Admin, error)
	LoginAdmin(ctx echo.Context, request web.AdminLoginRequest) (*domain.Admin, error)
	UpdateAdmin(ctx echo.Context, request web.AdminUpdateRequest, id int) (*domain.Admin, error)
	FindById(ctx echo.Context, id int) (*domain.Admin, error)
	FindAll(ctx echo.Context) ([]domain.Admin, error)
	DeleteAdmin(ctx echo.Context, id int) error
}

type AdminHandler interface {
	RegisterAdminHandler(ctx echo.Context) error
	LoginAdminHandler(ctx echo.Context) error
	UpdateAdminHandler(ctx echo.Context) error
	GetAdminHandler(ctx echo.Context) error
	GetAdminsHandler(ctx echo.Context) error
	DeleteAdminHandler(ctx echo.Context) error
}

type AdminRoutes interface {
	Auth()
	Admin()
}