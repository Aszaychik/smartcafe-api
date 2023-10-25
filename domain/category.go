package domain

import "gorm.io/gorm"

type Category struct {
	*gorm.Model
	CategoryName string `gorm:"type:varchar(255)" json:"category_name"`
	CategoryDescription string `gorm:"type:varchar(255)" json:"category_description"`
	Menus               []Menu `gorm:"foreignKey:CategoryID" json:"menus"`
}