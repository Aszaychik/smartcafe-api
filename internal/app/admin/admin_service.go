package admin

import (
	"aszaychik/smartcafe-api/domain"
	"aszaychik/smartcafe-api/domain/web"
	"aszaychik/smartcafe-api/internal/interfaces"
	"aszaychik/smartcafe-api/pkg/conversion"
	"aszaychik/smartcafe-api/pkg/password"
	"aszaychik/smartcafe-api/pkg/validation"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type AdminServiceImpl struct {
	AdminRepository interfaces.AdminRepository
	Validate       *validator.Validate
}

func NewAdminService(adminRepository interfaces.AdminRepository, validate *validator.Validate) interfaces.AdminService {
	return &AdminServiceImpl{
		AdminRepository: adminRepository,
		Validate:        validate,
	}
}

func (service *AdminServiceImpl) RegisterAdmin(ctx echo.Context, request web.AdminCreateRequest) (*domain.Admin, error) {
	// Check if the request is valid
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, validation.ValidationError(ctx, err)
	}

	// Check if the username already exists
	existingAdmin, _ := service.AdminRepository.FindByUsername(request.Username)
	if existingAdmin != nil {
		return nil, fmt.Errorf("Username already exists")
	}

	// Convert request to domain
	admin := conversion.AdminCreateRequestToAdminDomain(request)
	// Convert password to hash
	admin.Password = password.HashPassword(admin.Password)

	result, err := service.AdminRepository.Save(admin)
	if err != nil {
		return nil, fmt.Errorf("Error when register : %s", err.Error())
	}

	return result, nil
}

func (service *AdminServiceImpl) LoginAdmin(ctx echo.Context, request web.AdminLoginRequest) (*domain.Admin, error) {
	// Check if the request is valid
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, validation.ValidationError(ctx, err)
	}

	// Check if the username exists
	existingAdmin, err := service.AdminRepository.FindByUsername(request.Username)
	if err != nil {
		return nil, fmt.Errorf("Invalid username or password")
	}

	// Convert request to domain
	admin := conversion.AdminLoginRequestToAdminDomain(request)

	// Compare password
	err = password.ComparePassword(existingAdmin.Password, admin.Password)
	if err != nil {
		return nil, fmt.Errorf("Invalid username or password")
	}

	return existingAdmin, nil
}

func (service *AdminServiceImpl) UpdateAdmin(ctx echo.Context, request web.AdminUpdateRequest, id int) (*domain.Admin, error) {
	// Check if the request is valid
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, validation.ValidationError(ctx, err)
	}

	// Check if the admin exists
	existingAdmin, _ := service.AdminRepository.FindById(id)
	if existingAdmin == nil {
		return nil, fmt.Errorf("Admin not found")
	}

	// Convert request to domain
	admin := conversion.AdminUpdateRequestToAdminDomain(request)
	admin.Password = password.HashPassword(admin.Password)

	result, err := service.AdminRepository.Update(admin, id)
	if err != nil {
		return nil, fmt.Errorf("Error when updating : %s", err.Error())
	}

	return result, nil
}

func (service *AdminServiceImpl) FindById(ctx echo.Context, id int) (*domain.Admin, error) {
	// Check if the admin exists
	existingAdmin, _ := service.AdminRepository.FindById(id)
	if existingAdmin == nil {
		return nil, fmt.Errorf("Admin not found")
	}

	return existingAdmin, nil
}

func (service *AdminServiceImpl) FindAll(ctx echo.Context) ([]domain.Admin, error) {
	admins, err := service.AdminRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("Admins not found")
	}

	return admins, nil
}

func (service *AdminServiceImpl) DeleteAdmin(ctx echo.Context, id int) error {
	// Check if the admin exists
	existingAdmin, _ := service.AdminRepository.FindById(id)
	if existingAdmin == nil {
		return fmt.Errorf("Admin not found")
	}

	err := service.AdminRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("Error when deleting : %s", err)
	}

	return nil
}