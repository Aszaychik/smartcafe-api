package domain

import "gorm.io/gorm"

type Menu struct {
	*gorm.Model
	ItemName        string `gorm:"type:varchar(255)" json:"item_name"`
	ItemPrice       int    `gorm:"type:int" json:"item_price"`
	ItemDescription string `gorm:"type:varchar(255)" json:"item_description"`
	ItemImage       string `gorm:"type:varchar(255)" json:"item_image"`
}