package domain

type Customer struct {
	ID            uint   `gorm:"type:int;primarykey" json:"id"`
	CustomerName  string `gorm:"type:varchar(255)" json:"customer_name"`
	CustomerEmail string `gorm:"type:varchar(255);uniqueIndex" json:"customer_email"`
}