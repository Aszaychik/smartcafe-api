package midtrans

import (
	"aszaychik/smartcafe-api/config"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

func NewMidtransCoreAPIClient(config *config.MidtransConfig) coreapi.Client {
	var coreAPIClient coreapi.Client
	coreAPIClient.New(config.ServerKey, midtrans.Sandbox)

	return coreAPIClient
}

func TransactionStatus(transactionStatusResp *coreapi.TransactionStatusResponse) string {
	var paymentStatus string
	
	if transactionStatusResp.TransactionStatus == "capture" {
		if transactionStatusResp.FraudStatus == "challenge" {
			paymentStatus = "challenge"
		} else if transactionStatusResp.FraudStatus == "accept" {
			paymentStatus = "success"
		}
	} else if transactionStatusResp.TransactionStatus == "settlement" {
		paymentStatus = "success"
	} else if transactionStatusResp.TransactionStatus == "deny" {
		paymentStatus = "denied"
	} else if transactionStatusResp.TransactionStatus == "cancel" || transactionStatusResp.TransactionStatus == "expire" {
		paymentStatus = "failed"
	} else if transactionStatusResp.TransactionStatus == "pending" {
		paymentStatus = "pending"
	}

	return paymentStatus
}

func TransactionStatusMethod(transactionStatusResp *coreapi.TransactionStatusResponse) string {
	return transactionStatusResp.PaymentType
}