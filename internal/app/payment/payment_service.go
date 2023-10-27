package payment

import (
	"aszaychik/smartcafe-api/domain"
	"aszaychik/smartcafe-api/internal/interfaces"
	"aszaychik/smartcafe-api/pkg/midtrans"
	"errors"
	"time"

	"github.com/midtrans/midtrans-go/coreapi"
)

type OrderPaymentServiceImpl struct {
	OrderPaymentRepository interfaces.OrderPaymentRepository
	OrderRepository interfaces.OrderRepository
	CoreAPIClient coreapi.Client
}

func NewOrderPaymentService(menuRepository interfaces.OrderPaymentRepository, orderRepository interfaces.OrderRepository, coreAPIClient coreapi.Client) interfaces.OrderPaymentService {
	return &OrderPaymentServiceImpl{
		OrderPaymentRepository: menuRepository,
		OrderRepository: orderRepository,
		CoreAPIClient: coreAPIClient,
	}
}

func (service *OrderPaymentServiceImpl) Notifications(notificationPayload map[string]any) error {
	orderId, exist := notificationPayload["order_id"].(string)
	if !exist {
		return errors.New("invalid notification payload")
	}

	paymentType, exist := notificationPayload["payment_type"].(string)
	if !exist {
		return errors.New("invalid notification payload")
	}

	// paymentDate, exist := notificationPayload["transaction_time"].(string)
	// if !exist {
	// 	return errors.New("invalid notification payload")
	// }

	orderPaymentStatus, err := service.CheckTransaction(orderId)
	if err != nil {
		return err
	}

	orderPayment, err := service.OrderPaymentRepository.FindById(orderId)
	if err != nil {
		return errors.New("transaction data not found")
	}

	orderPaymentStatus.PaymentMethod = paymentType
	orderPaymentStatus.OrderID = orderPayment.OrderID
	orderPaymentStatus.PaymentDate = time.Now()
	// orderPaymentStatus.PaymentDate, err = time.Parse("2006-01-02 15:04:05", paymentDate)
	// if err != nil {
	// 	return fmt.Errorf("Parse time error: %w", err)
	// }

	err = service.OrderPaymentRepository.UpdateOrderPaymentStatus(orderPayment.ID, orderPaymentStatus)
	if err != nil {
		return err
	}

	err = service.OrderRepository.UpdateOrderStatus(orderPaymentStatus)
	if err != nil {
		return err
	}

	return nil
}

func (service *OrderPaymentServiceImpl) CheckTransaction(orderID string) (domain.PaymentStatus, error) {
	PaymentStatus := domain.PaymentStatus{}

	transactionStatusResp, err := service.CoreAPIClient.CheckTransaction(orderID)
	if err != nil {
		return PaymentStatus, err
	} else {
		if transactionStatusResp != nil {
			PaymentStatus.PaymentStatus = midtrans.TransactionStatus(transactionStatusResp)
			return PaymentStatus, nil
		}
	}

	return PaymentStatus, err
}