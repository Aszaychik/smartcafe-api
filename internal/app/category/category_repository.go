package category

import (
	"aszaychik/smartcafe-api/domain"
	"aszaychik/smartcafe-api/internal/interfaces"

	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) interfaces.CategoryRepository {
	return &CategoryRepositoryImpl{DB: db}
}

func (repository *CategoryRepositoryImpl) Save(category *domain.Category) (*domain.Category, error) {
	result := repository.DB.Create(&category)
	if result.Error != nil {
		return nil, result.Error
	}

	return category, nil
}

func (repository *CategoryRepositoryImpl) Update(category *domain.Category, id int) (*domain.Category, error) {
	result := repository.DB.Save(&category)
	if result.Error != nil {
		return nil, result.Error
	}

	return category, nil
}

func (repository *CategoryRepositoryImpl) FindById(id int) (*domain.Category, error) {
	category := domain.Category{}

	result := repository.DB.First(&category, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &category, nil
}

func (repository *CategoryRepositoryImpl) FindByName(categoryName string) (*domain.Category, error) {
	category := domain.Category{}

	result := repository.DB.Where("category_name = ?", categoryName).First(&category)
	if result.Error != nil {
		return nil, result.Error
	}

	return &category, nil
}

func (repository *CategoryRepositoryImpl) FindByCategoryId(categoryId int) ([]domain.Category, error) {
	category := []domain.Category{}

	result := repository.DB.Where("category_id = ?", categoryId).First(&category)
	if result.Error != nil {
		return nil, result.Error
	}

	return category, nil
}

func (repository *CategoryRepositoryImpl) FindAll() ([]domain.Category, error) {
	category := []domain.Category{}

	result := repository.DB.Find(&category)
	if result.Error != nil {
		return nil, result.Error
	}

	return category, nil
}

func (repository *CategoryRepositoryImpl) Delete(id int) error {
	result := repository.DB.Table("categories").Where("id = ?", id).Unscoped().Delete(id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}