package interfaces

import (
	"aszaychik/smartcafe-api/domain"
	"aszaychik/smartcafe-api/domain/web"

	"github.com/labstack/echo/v4"
)

type CategoryRepository interface {
	Save(category *domain.Category) (*domain.Category, error)
	Update(category *domain.Category, id int) (*domain.Category, error)
	FindById(id int) (*domain.Category, error)
	FindByName(categoryName string) (*domain.Category, error)
	FindByCategoryId(categoryId int) ([]domain.Category, error)
	FindAll() ([]domain.Category, error)
	Delete(id int) error
}

type CategoryService interface {
	CreateCategory(ctx echo.Context, request web.CategoryCreateRequest) (*domain.Category, error)
	UpdateCategory(ctx echo.Context, request web.CategoryUpdateRequest, id int) (*domain.Category, error)
	FindById(ctx echo.Context, id int) (*domain.Category, error)
	FindAll(ctx echo.Context) ([]domain.Category, error)
	DeleteCategory(ctx echo.Context, id int) error
}

type CategoryHandler interface {
	CreateCategoryHandler(ctx echo.Context) error
	UpdateCategoryHandler(ctx echo.Context) error
	GetCategoryHandler(ctx echo.Context) error
	GetCategoriesHandler(ctx echo.Context) error
	DeleteCategoryHandler(ctx echo.Context) error
}

type CategoryRoutes interface {
	Category()
}