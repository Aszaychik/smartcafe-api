package domain

import "gorm.io/gorm"

type Admin struct {
	*gorm.Model
	Username string `gorm:"type:varchar(255)" json:"username"`
	Password string `gorm:"type:varchar(255)" json:"password"`
}