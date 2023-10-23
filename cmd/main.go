package main

import (
	"aszaychik/smartcafe-api/domain"
	"aszaychik/smartcafe-api/internal/infrastructure/config"
	"aszaychik/smartcafe-api/internal/infrastructure/database"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		logrus.Fatal(err.Error())
	}
	
	fmt.Println(cfg)

	db, err := database.NewMySQLConnection(&cfg.MySQL)
	if err != nil {
			logrus.Fatal("Error connecting to MySQL:", err.Error())
	}

	db.AutoMigrate(&domain.Admin{}, &domain.Menu{})

	e := echo.New()

	e.Logger.Fatal(e.Start(":8080"))
}