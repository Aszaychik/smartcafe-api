package customer

import (
	"aszaychik/smartcafe-api/domain"
	"aszaychik/smartcafe-api/internal/interfaces"

	"gorm.io/gorm"
)

type CustomerRepositoryImpl struct {
	DB *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) interfaces.CustomerRepository {
	return &CustomerRepositoryImpl{DB: db}
}

func (repository *CustomerRepositoryImpl) Save(customer *domain.Customer) (*domain.Customer, error) {
	result := repository.DB.Create(&customer)
	if result.Error != nil {
		return nil, result.Error
	}

	return customer, nil
}

func (repository *CustomerRepositoryImpl) Update(customer *domain.Customer, id int) (*domain.Customer, error) {
	result := repository.DB.Save(&customer)
	if result.Error != nil {
		return nil, result.Error
	}

	return customer, nil
}

func (repository *CustomerRepositoryImpl) FindById(id int) (*domain.Customer, error) {
	customer := domain.Customer{}

	result := repository.DB.First(&customer, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &customer, nil
}

func (repository *CustomerRepositoryImpl) FindByName(itemName string) (*domain.Customer, error) {
	customer := domain.Customer{}

	result := repository.DB.Where("item_name = ?", itemName).First(&customer)
	if result.Error != nil {
		return nil, result.Error
	}

	return &customer, nil
}

func (repository *CustomerRepositoryImpl) FindByEmail(email string) (*domain.Customer, error) {
	customer := domain.Customer{}

	result := repository.DB.Where("email = ?", email).First(&customer)
	if result.Error != nil {
		return nil, result.Error
	}

	return &customer, nil
}

func (repository *CustomerRepositoryImpl) FindAll() ([]domain.Customer, error) {
	customer := []domain.Customer{}

	result := repository.DB.Find(&customer)
	if result.Error != nil {
		return nil, result.Error
	}

	return customer, nil
}

func (repository *CustomerRepositoryImpl) Delete(id int) error {
	result := repository.DB.Table("customers").Where("id = ?", id).Unscoped().Delete(id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}