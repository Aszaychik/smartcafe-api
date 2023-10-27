package database

import (
	"aszaychik/smartcafe-api/config"
	"aszaychik/smartcafe-api/domain"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQLConnection(cfg *config.MySQLConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.Username,
			cfg.Password,
			cfg.Host,
			cfg.Port,
			cfg.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	migration(db)

	return db, nil
}

func migration(db *gorm.DB) {
	db.AutoMigrate(&domain.Admin{}, &domain.Category{}, &domain.Menu{}, &domain.Customer{}, &domain.Order{}, &domain.OrderItem{}, &domain.OrderPayment{})
}