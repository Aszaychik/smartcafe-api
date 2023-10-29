package category

import (
	"aszaychik/smartcafe-api/domain"
	"aszaychik/smartcafe-api/domain/web"
	"aszaychik/smartcafe-api/internal/interfaces"
	"aszaychik/smartcafe-api/pkg/conversion"
	"aszaychik/smartcafe-api/pkg/validation"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CategoryServiceImpl struct {
	CategoryRepository interfaces.CategoryRepository
	Validate       *validator.Validate
}

func NewCategoryService(categoryRepository interfaces.CategoryRepository, validate *validator.Validate) interfaces.CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		Validate:        validate,
	}
}

func (service *CategoryServiceImpl) CreateCategory(ctx echo.Context, request web.CategoryCreateRequest) (*domain.Category, error) {
	// Check if the request is valid
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, validation.ValidationError(ctx, err)
	}

	// Check if the Category name already exists
	existingCategory, _ := service.CategoryRepository.FindByName(request.CategoryName)
	if existingCategory != nil {
		return nil, fmt.Errorf("Category Name already exists")
	}

	// Convert request to domain
	category := conversion.CategoryCreateRequestToCategoryDomain(request)

	result, err := service.CategoryRepository.Save(category)
	if err != nil {
		return nil, fmt.Errorf("Error when create : %s", err.Error())
	}

	result, err = service.CategoryRepository.FindById(int(result.ID))
	if err != nil {
		return nil, fmt.Errorf("Error when retrieve  : %s", err.Error())
	}

	return result, nil
}

func (service *CategoryServiceImpl) UpdateCategory(ctx echo.Context, request web.CategoryUpdateRequest, id int) (*domain.Category, error) {
	// Check if the request is valid
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, validation.ValidationError(ctx, err)
	}

	// Check if the category exists
	existingCategoryId, _ := service.CategoryRepository.FindById(id)
	if existingCategoryId == nil {
		return nil, fmt.Errorf("Category not found")
	}

	existingCategoryName, _ := service.CategoryRepository.FindByName(request.CategoryName)
	if existingCategoryName != nil {
		return nil, fmt.Errorf("Category Name already exists")
	}

	// Convert request to domain
	category := conversion.CategoryUpdateRequestToCategoryDomain(request)

	result, err := service.CategoryRepository.Update(category, id)
	if err != nil {
		return nil, fmt.Errorf("Error when updating : %s", err.Error())
	}

	result, _ = service.CategoryRepository.FindByName(category.CategoryName)

	return result, nil
}

func (service *CategoryServiceImpl) FindById(ctx echo.Context, id int) (*domain.Category, error) {
	// Check if the category exists
	existingCategory, _ := service.CategoryRepository.FindById(id)
	if existingCategory == nil {
		return nil, fmt.Errorf("Category not found")
	}

	return existingCategory, nil
}

func (service *CategoryServiceImpl) FindAll(ctx echo.Context) ([]domain.Category, error) {
	categories, err := service.CategoryRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("Categories not found")
	}

	return categories, nil
}

func (service *CategoryServiceImpl) DeleteCategory(ctx echo.Context, id int) error {
	// Check if the category exists
	existingCategory, _ := service.CategoryRepository.FindById(id)
	if existingCategory == nil {
		return fmt.Errorf("Category not found")
	}

	err := service.CategoryRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("Error when deleting : %s", err)
	}

	return nil
}