package payment

import (
	"aszaychik/smartcafe-api/domain"
	"aszaychik/smartcafe-api/internal/interfaces"

	"gorm.io/gorm"
)

type OrderPaymentRepositoryImpl struct {
	DB *gorm.DB
}

func NewOrderPaymentRepository(db *gorm.DB) interfaces.OrderPaymentRepository {
	return &OrderPaymentRepositoryImpl{DB: db}
}

func (repository *OrderPaymentRepositoryImpl) FindById(id string) (*domain.OrderPayment, error) {
	OrderPayment := domain.OrderPayment{}

	result := repository.DB.Where("id = ?", id).Find(&OrderPayment)
	if result.Error != nil {
		return nil, result.Error
	}

	return &OrderPayment, nil
}

func (repository *OrderPaymentRepositoryImpl) UpdateOrderPaymentStatus(paymentId string, orderPaymentStatus domain.PaymentStatus) error {
	result := repository.DB.Table("order_payments").Where("id = ?", paymentId).Updates(domain.OrderPayment{
		PaymentStatus: orderPaymentStatus.PaymentStatus,
		PaymentMethod: orderPaymentStatus.PaymentMethod,
		PaymentDate: orderPaymentStatus.PaymentDate,
	})
	
	if result.Error != nil {
		return result.Error
	}

	return nil
}
