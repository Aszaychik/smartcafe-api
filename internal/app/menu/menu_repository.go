package menu

import (
	"aszaychik/smartcafe-api/domain"
	"aszaychik/smartcafe-api/internal/interfaces"

	"gorm.io/gorm"
)

type MenuRepositoryImpl struct {
	DB *gorm.DB
}

func NewMenuRepository(db *gorm.DB) interfaces.MenuRepository {
	return &MenuRepositoryImpl{DB: db}
}

func (repository *MenuRepositoryImpl) Save(menu *domain.Menu) (*domain.Menu, error) {
	result := repository.DB.Create(&menu)
	if result.Error != nil {
		return nil, result.Error
	}

	return menu, nil
}

func (repository *MenuRepositoryImpl) Update(menu *domain.Menu, id int) (*domain.Menu, error) {
	result := repository.DB.Save(&menu)
	if result.Error != nil {
		return nil, result.Error
	}

	return menu, nil
}

func (repository *MenuRepositoryImpl) FindById(id int) (*domain.Menu, error) {
	menu := domain.Menu{}

	result := repository.DB.First(&menu, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &menu, nil
}

func (repository *MenuRepositoryImpl) FindByName(itemName string) (*domain.Menu, error) {
	menu := domain.Menu{}

	result := repository.DB.Where("item_name = ?", itemName).First(&menu)
	if result.Error != nil {
		return nil, result.Error
	}

	return &menu, nil
}

func (repository *MenuRepositoryImpl) FindByCategoryId(categoryId int) ([]domain.Menu, error) {
	menu := []domain.Menu{}

	result := repository.DB.Where("category_id = ?", categoryId).First(&menu)
	if result.Error != nil {
		return nil, result.Error
	}

	return menu, nil
}

func (repository *MenuRepositoryImpl) FindAll() ([]domain.Menu, error) {
	menu := []domain.Menu{}

	result := repository.DB.Find(&menu)
	if result.Error != nil {
		return nil, result.Error
	}

	return menu, nil
}

func (repository *MenuRepositoryImpl) Delete(id int) error {
	result := repository.DB.Table("menus").Where("id = ?", id).Unscoped().Delete(id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}