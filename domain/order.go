package domain

import "time"

type OrderStatus string

const (
	Pending    OrderStatus = "pending"
	Preparing  OrderStatus = "preparing"
	Completed  OrderStatus = "completed"
)

type Order struct {
	ID          uint `gorm:"type:int;primarykey" json:"id"`
	CustomerID  uint `gorm:"type:int" json:"customer_id"`
	TotalPrice  int  `gorm:"type:decimal(10,2)" json:"total_price"`
	OrderStatus OrderStatus `gorm:"type:varchar(255)" json:"order_status"`
	OrderDate   time.Time `gorm:"type:datetime" json:"order_date"`
	Customer Customer `gorm:"foreignKey:CustomerID" json:"customer"`
}