package order

import (
	"aszaychik/smartcafe-api/domain"
	"aszaychik/smartcafe-api/internal/interfaces"

	"gorm.io/gorm"
)

type OrderRepositoryImpl struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) interfaces.OrderRepository {
	return &OrderRepositoryImpl{DB: db}
}

func (repository *OrderRepositoryImpl) Save(order *domain.Order) (*domain.Order, error) {
	result := repository.DB.Create(&order)
	if result.Error != nil {
		return nil, result.Error
	}

	return order, nil
}

func (repository *OrderRepositoryImpl) FindById(orderId int) (*domain.Order, error) {
	order := domain.Order{}

	result := repository.DB.Preload("Items").Preload("Customer").Preload("OrderPayment").First(&order, orderId)
	if result.Error != nil {
		return nil, result.Error
	}

	// Preload the Item and its Category fields
	for i := range order.Items {
		if err := repository.DB.Preload("Category").Model(&order.Items[i]).Association("Item").Find(&order.Items[i].Item); err != nil {
			return nil, err
		}
	}

	return &order, nil
}

func (repository *OrderRepositoryImpl) UpdateOrderStatus(orderPayment domain.PaymentStatus) error {
	result := repository.DB.Table("Orders").Where("ID = ?", orderPayment.OrderID).Update("order_status", orderPayment.PaymentStatus)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
