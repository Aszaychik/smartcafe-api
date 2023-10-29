package domain

import "time"

type Feedback struct {
	ID          uint `gorm:"type:int;primarykey" json:"id"`
	OrderID        uint       `gorm:"type:int" json:"order_id"`
	CustomerID     uint       `gorm:"type:int" json:"customer_id"`
	FeedbackText   string    `gorm:"type:text" json:"feedback_text"`
	FeedbackRating int       `gorm:"type:int" json:"feedback_rating"`
	FeedbackDate   time.Time `gorm:"type:datetime" json:"feedback_date"`
	Customer Customer `gorm:"foreignKey:CustomerID" json:"customer"`
}