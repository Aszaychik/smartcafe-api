package domain

import (
	"time"
)

type OrderPayment struct {
	ID            string   `gorm:"type:varchar(255);primarykey" json:"id"`
	OrderID       uint   `gorm:"type:int" json:"order_id"`
	PaymentMethod string `gorm:"type:varchar(255)" json:"payment_method"`
	PaymentStatus string `gorm:"type:varchar(255);default:'pending'" json:"payment_status"`
	PaymentDate time.Time `gorm:"type:datetime" json:"payment_date"`
	PaymentUrl string `gorm:"type:varchar(255)" json:"payment_url"`
}

type PaymentStatus struct {
	PaymentId string `json:"payment_id"`
	PaymentStatus string `json:"payment_status"`
	PaymentMethod string `json:"payment_method"`
	PaymentDate time.Time `json:"payment_date"`
	OrderID uint `json:"order_id"`
}