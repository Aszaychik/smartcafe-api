package menu

import (
	"aszaychik/smartcafe-api/domain"
	"aszaychik/smartcafe-api/domain/web"
	"aszaychik/smartcafe-api/internal/interfaces"
	"aszaychik/smartcafe-api/pkg/conversion"
	"aszaychik/smartcafe-api/pkg/uploader"
	"aszaychik/smartcafe-api/pkg/validation"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strings"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type MenuServiceImpl struct {
	MenuRepository interfaces.MenuRepository
	Validate       *validator.Validate
	AWSUploader uploader.AWSUploader
}

func NewMenuService(menuRepository interfaces.MenuRepository, validate *validator.Validate, awsUploader uploader.AWSUploader) interfaces.MenuService {
	return &MenuServiceImpl{
		MenuRepository: menuRepository,
		Validate:        validate,
		AWSUploader: awsUploader,
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

	result, _ = service.MenuRepository.FindByName(request.ItemName)

	return result, nil
}

func (service *MenuServiceImpl) UpdateMenu(ctx echo.Context, request web.MenuUpdateRequest, id int) (*domain.Menu, error) {
	// Check if the request is valid
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, validation.ValidationError(ctx, err)
	}

	// Check if the menu exists
	existingMenuId, _ := service.MenuRepository.FindById(id)
	if existingMenuId == nil {
		return nil, fmt.Errorf("Menu not found")
	}

	existingMenuName, _ := service.MenuRepository.FindByName(request.ItemName)
	if existingMenuName != nil && int(existingMenuName.ID) != id  {
		return nil, fmt.Errorf("Menu Name already exists")
	}

	// Convert request to domain
	menu := conversion.MenuUpdateRequestToMenuDomain(request)

	result, err := service.MenuRepository.Update(menu, id)
	if err != nil {
		return nil, fmt.Errorf("Error when updating : %s", err.Error())
	}

	result, _ = service.MenuRepository.FindByName(request.ItemName)

	return result, nil
}

func (service *MenuServiceImpl) UpdateImageMenu(ctx echo.Context, fileHeader *multipart.FileHeader, id int) (*domain.Menu, error) {
	// Check if the menu exists
	existingMenu, err := service.MenuRepository.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("Error finding menu: %s", err.Error())
	}
	if existingMenu == nil {
		return nil, fmt.Errorf("Menu not found")
	}

	// Open the uploaded file
	file, err := fileHeader.Open()
	if err != nil {
		return nil, fmt.Errorf("Error opening file: %s", err.Error())
	}
	defer file.Close()

	// Generate a unique filename using the current timestamp
	fileName := fmt.Sprintf("%s-%s", time.Now(), strings.ReplaceAll(fileHeader.Filename, " ", "-"))

	// Create a new file on the server
	localFilePath := fmt.Sprintf("uploads/%s", fileName)
	localFile, err := os.Create(localFilePath)
	if err != nil {
		return nil, fmt.Errorf("Error creating file: %s", err.Error())
	}
	defer localFile.Close()

	// Copy the contents of the uploaded file to the new file on the server
	_, err = io.Copy(localFile, file)
	if err != nil {
		return nil, fmt.Errorf("Error copying file contents: %s", err.Error())
	}

	// Specify the S3 folder name
	folderName := "menu/images"

	// Upload the file to AWS S3
	uploadedImage, err := service.AWSUploader.UploadFile(fileName, folderName)
	if err != nil {
		return nil, fmt.Errorf("Error when uploading to AWS S3: %s", err.Error())
	}

	// Update the menu record with the S3 image location
	err = service.MenuRepository.UpdateImage(uploadedImage.Location, id)
	if err != nil {
		return nil, fmt.Errorf("Error updating menu image: %s", err.Error())
	}

	// Fetch and return the updated menu
	result, err := service.MenuRepository.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("Error finding updated menu: %s", err.Error())
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