package admin

import (
	"aszaychik/smartcafe-api/domain/web"
	"aszaychik/smartcafe-api/internal/interfaces"
	"aszaychik/smartcafe-api/pkg/conversion"
	"aszaychik/smartcafe-api/pkg/jwt"
	"aszaychik/smartcafe-api/pkg/res"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type AdminHandlerImpl struct {
	AdminService interfaces.AdminService
}

func NewAdminHandler(adminService interfaces.AdminService) interfaces.AdminHandler {
	return &AdminHandlerImpl{AdminService: adminService}
}

func(handler *AdminHandlerImpl) RegisterAdminHandler(ctx echo.Context) error {
	adminCreateRequest := web.AdminCreateRequest{}
	err := ctx.Bind(&adminCreateRequest) 
	if err != nil {
		return res.StatusBadRequest(ctx, err)
	}

	response, err := handler.AdminService.RegisterAdmin(ctx, adminCreateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return res.StatusBadRequest(ctx, err)
		}

		if strings.Contains(err.Error(), "Username already exists") {
			return res.StatusAlreadyExist(ctx, err)
		}

		return res.StatusInternalServerError(ctx, err)
	}

	return res.StatusCreated(ctx, "Success to create admin", response)
}

func (handler *AdminHandlerImpl) LoginAdminHandler(ctx echo.Context) error {
	adminLoginRequest := web.AdminLoginRequest{}
	err := ctx.Bind(&adminLoginRequest)
	if err != nil {
		return res.StatusBadRequest(ctx, err)
	}

	response, err := handler.AdminService.LoginAdmin(ctx, adminLoginRequest)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return res.StatusBadRequest(ctx, err)
		}
		
		if strings.Contains(err.Error(), "Invalid username or password") {
			return res.StatusBadRequest(ctx, err)
		}
		
		return res.StatusInternalServerError(ctx, err)
	}

	adminLoginResponse := conversion.AdminDomainToAdminLoginResponse(response)

	token, err := jwt.GenerateToken(&adminLoginResponse, response.ID)
	if err != nil {
		return res.StatusInternalServerError(ctx, err)
	}

	adminLoginResponse.Token = token

	return res.StatusOK(ctx, "Success to login admin", adminLoginResponse)
}


func (handler *AdminHandlerImpl) UpdateAdminHandler(ctx echo.Context) error {
	adminId := ctx.Param("id")
	adminIdInt, err := strconv.Atoi(adminId)
	if err != nil {
		return res.StatusInternalServerError(ctx, err)
	}

	adminUpdateRequest := web.AdminUpdateRequest{}
	err = ctx.Bind(&adminUpdateRequest)
	if err != nil {
		return res.StatusBadRequest(ctx, err)
	}

	response, err := handler.AdminService.UpdateAdmin(ctx, adminUpdateRequest, adminIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return res.StatusBadRequest(ctx, err)
		}

		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, err)
		}

		return res.StatusInternalServerError(ctx, err)
	}

	return res.StatusOK(ctx, "Success to update admin", response)
}

func (handler *AdminHandlerImpl) GetAdminHandler(ctx echo.Context) error {
	adminId := ctx.Param("id")
	adminIdInt, err := strconv.Atoi(adminId)
	if err != nil {
		return res.StatusInternalServerError(ctx, err)
	}
	
	response, err := handler.AdminService.FindById(ctx, adminIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, err)
		}

		return res.StatusInternalServerError(ctx, err)
	}

	return res.StatusOK(ctx, "Success to get admin", response)
}

func (handler *AdminHandlerImpl) GetAdminsHandler(ctx echo.Context) error {
	response, err := handler.AdminService.FindAll(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, err)
		}

		return res.StatusInternalServerError(ctx, err)
	}

	return res.StatusOK(ctx, "Success to get admins", response)
}

func (handler *AdminHandlerImpl) DeleteAdminHandler(ctx echo.Context) error {
	adminId := ctx.Param("id")
	adminIdInt, err := strconv.Atoi(adminId)
	if err != nil {
		return res.StatusInternalServerError(ctx, err)
	}

	err = handler.AdminService.DeleteAdmin(ctx, adminIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, err)
		}

		return res.StatusInternalServerError(ctx, err)
	}

	return res.StatusOK(ctx, "Success to delete admin", nil)
}