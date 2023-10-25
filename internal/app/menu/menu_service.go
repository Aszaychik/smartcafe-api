package menu

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

type MenuServiceImpl struct {
	MenuRepository interfaces.MenuRepository
	Validate       *validator.Validate
}

func NewMenuService(menuRepository interfaces.MenuRepository, validate *validator.Validate) interfaces.MenuService {
	return &MenuServiceImpl{
		MenuRepository: menuRepository,
		Validate:        validate,
	}
}

func (service *MenuServiceImpl) CreateMenu(ctx echo.Context, request web.MenuCreateRequest) (*domain.Menu, error) {
	if request.ItemImage == "" {
		request.ItemImage = "https://placewaifu.com/image/200"
	}
	
	// Check if the request is valid
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, validation.ValidationError(ctx, err)
	}

	// Check if the Menu name already exists
	existingMenu, _ := service.MenuRepository.FindByName(request.ItemName)
	if existingMenu != nil {
		return nil, fmt.Errorf("Menu Name already exists")
	}

	// Convert request to domain
	menu := conversion.MenuCreateRequestToMenuDomain(request)

	result, err := service.MenuRepository.Save(menu)
	if err != nil {
		return nil, fmt.Errorf("Error when create : %s", err.Error())
	}

	return result, nil
}

func (service *MenuServiceImpl) UpdateMenu(ctx echo.Context, request web.MenuUpdateRequest, id int) (*domain.Menu, error) {
	// Check if the request is valid
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, validation.ValidationError(ctx, err)
	}

	// Check if the menu exists
	existingMenu, _ := service.MenuRepository.FindById(id)
	if existingMenu == nil {
		return nil, fmt.Errorf("Menu not found")
	}

	// Convert request to domain
	menu := conversion.MenuUpdateRequestToMenuDomain(request)

	result, err := service.MenuRepository.Update(menu, id)
	if err != nil {
		return nil, fmt.Errorf("Error when updating : %s", err.Error())
	}

	return result, nil
}

func (service *MenuServiceImpl) FindById(ctx echo.Context, id int) (*domain.Menu, error) {
	// Check if the menu exists
	existingMenu, _ := service.MenuRepository.FindById(id)
	if existingMenu == nil {
		return nil, fmt.Errorf("Menu not found")
	}

	return existingMenu, nil
}

func (service *MenuServiceImpl) FindAll(ctx echo.Context) ([]domain.Menu, error) {
	menus, err := service.MenuRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("Menus not found")
	}

	return menus, nil
}

func (service *MenuServiceImpl) DeleteMenu(ctx echo.Context, id int) error {
	// Check if the menu exists
	existingMenu, _ := service.MenuRepository.FindById(id)
	if existingMenu == nil {
		return fmt.Errorf("Menu not found")
	}

	err := service.MenuRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("Error when deleting : %s", err)
	}

	return nil
}