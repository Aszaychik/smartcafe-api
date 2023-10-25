package conversion

import (
	"aszaychik/smartcafe-api/domain"
	"aszaychik/smartcafe-api/domain/web"
)

func AdminDomainToAdminLoginResponse(admin *domain.Admin) web.AdminLoginResponse {
	return web.AdminLoginResponse {
		Username:admin.Username,
	}
}