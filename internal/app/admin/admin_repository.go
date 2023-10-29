package admin

import (
	"aszaychik/smartcafe-api/domain"
	"aszaychik/smartcafe-api/internal/interfaces"

	"gorm.io/gorm"
)

type AdminRepositoryImpl struct {
	DB *gorm.DB
}

func NewAdminRepository(db *gorm.DB) interfaces.AdminRepository {
	return &AdminRepositoryImpl{DB: db}
}

func (repository *AdminRepositoryImpl) Save(admin *domain.Admin) (*domain.Admin, error) {
	result := repository.DB.Create(&admin)
	if result.Error != nil {
		return nil, result.Error
	}

	return admin, nil
}

func (repository *AdminRepositoryImpl) Update(admin *domain.Admin, id int) (*domain.Admin, error) {
	result := repository.DB.Table("admins").Where("id = ?", id).Updates(domain.Admin{Username: admin.Username, Password: admin.Password})
	if result.Error != nil {
		return nil, result.Error
	}

	return admin, nil
}

func (repository *AdminRepositoryImpl) FindById(id int) (*domain.Admin, error) {
	admin := domain.Admin{}

	result := repository.DB.First(&admin, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &admin, nil
}

func (repository *AdminRepositoryImpl) FindByUsername(username string) (*domain.Admin, error) {
	admin := domain.Admin{}

	result := repository.DB.Where("username = ?", username).First(&admin)
	if result.Error != nil {
		return nil, result.Error
	}

	return &admin, nil
}

func (repository *AdminRepositoryImpl) FindAll() ([]domain.Admin, error) {
	admin := []domain.Admin{}

	result := repository.DB.Find(&admin)
	if result.Error != nil {
		return nil, result.Error
	}

	return admin, nil
}

func (repository *AdminRepositoryImpl) Delete(id int) error {
	result := repository.DB.Table("admins").Where("id = ?", id).Unscoped().Delete(id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}