package midtrans

import (
	"aszaychik/smartcafe-api/config"
	"fmt"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func New(config *config.MidtransConfig) snap.Client {
	var snapClient snap.Client
	snapClient.New(config.ServerKey, midtrans.Sandbox)

	return snapClient
}

func CreateSnapRequest(snapClient snap.Client, orderID string, totalPrice int64) (*snap.Response, error) {
	snapRequest := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderID,
			GrossAmt: totalPrice,
		},
		EnabledPayments: snap.AllSnapPaymentType,
	}

	snapResponse, err := snapClient.CreateTransaction(snapRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to get snap response: %w", err)
	}

	return snapResponse, nil
}