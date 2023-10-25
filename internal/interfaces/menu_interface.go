package interfaces

import (
	"aszaychik/smartcafe-api/domain"
	"aszaychik/smartcafe-api/domain/web"

	"github.com/labstack/echo/v4"
)

type MenuRepository interface {
	Save(menu *domain.Menu) (*domain.Menu, error)
	Update(menu *domain.Menu, id int) (*domain.Menu, error)
	FindById(id int) (*domain.Menu, error)
	FindByName(itemName string) (*domain.Menu, error)
	FindByCategoryId(categoryId int) ([]domain.Menu, error)
	FindAll() ([]domain.Menu, error)
	Delete(id int) error
}

type MenuService interface {
	CreateMenu(ctx echo.Context, request web.MenuCreateRequest) (*domain.Menu, error)
	UpdateMenu(ctx echo.Context, request web.MenuUpdateRequest, id int) (*domain.Menu, error)
	FindById(ctx echo.Context, id int) (*domain.Menu, error)
	FindAll(ctx echo.Context) ([]domain.Menu, error)
	DeleteMenu(ctx echo.Context, id int) error
}

type MenuHandler interface {
	CreateMenuHandler(ctx echo.Context) error
	UpdateMenuHandler(ctx echo.Context) error
	GetMenuHandler(ctx echo.Context) error
	GetMenusHandler(ctx echo.Context) error
	DeleteMenuHandler(ctx echo.Context) error
}

type MenuRoutes interface {
	Menu()
}