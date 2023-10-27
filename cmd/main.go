package main

import (
	"aszaychik/smartcafe-api/internal/app/admin"
	"aszaychik/smartcafe-api/internal/app/category"
	"aszaychik/smartcafe-api/internal/app/customer"
	"aszaychik/smartcafe-api/internal/app/menu"
	"aszaychik/smartcafe-api/internal/app/order"
	"aszaychik/smartcafe-api/internal/infrastructure/config"
	"aszaychik/smartcafe-api/internal/infrastructure/database"
	"aszaychik/smartcafe-api/pkg/midtrans"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		logrus.Fatal("Error loading config:", err.Error())
	}

	// Initialize database connection
	db, err := database.NewMySQLConnection(&cfg.MySQL)
	if err != nil {
		logrus.Fatal("Error connecting to MySQL:", err.Error())
	}

	// Create a validator instance
	validate := validator.New()

	snapClient := midtrans.New(&cfg.Midtrans)

	// Create an Echo instance
	e := echo.New()

	// Set up handler
	// Admin
	adminRepository := admin.NewAdminRepository(db)
	adminService := admin.NewAdminService(adminRepository, validate)
	adminHandler := admin.NewAdminHandler(adminService)
	adminRoutes := admin.NewAdminRoutes(e, adminHandler)

	// Menu
	menuRepository := menu.NewMenuRepository(db)
	menuService := menu.NewMenuService(menuRepository, validate)
	menuHandler := menu.NewMenuHandler(menuService)
	menuRoutes := menu.NewMenuRoutes(e, menuHandler)

	// Category
	categoryRepository := category.NewCategoryRepository(db)
	categoryService := category.NewCategoryService(categoryRepository, validate)
	categoryHandler := category.NewCategoryHandler(categoryService)
	categoryRoutes := category.NewCategoryRoutes(e, categoryHandler)

	// Customer
	customerRepository := customer.NewCustomerRepository(db)
	customerService := customer.NewCustomerService(customerRepository, validate)
	customerHandler := customer.NewCustomerHandler(customerService)
	customerRoutes := customer.NewCustomerRoutes(e, customerHandler)

	// Order
	orderRepository := order.NewOrderRepository(db)
	orderService := order.NewOrderService(orderRepository, menuRepository, customerRepository, validate, snapClient)
	orderHandler := order.NewOrderHandler(orderService)
	orderRoutes := order.NewOrderRoutes(e, orderHandler)


	// Set up routes
	adminRoutes.Auth()
	adminRoutes.Admin()
	menuRoutes.Menu()
	categoryRoutes.Category()
	customerRoutes.Customer()
	orderRoutes.Order()
	
	// Middleware and server configuration
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
		},
	))

	// Start the Echo server in a goroutine
	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("Error starting server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Shutdown Echo gracefully
	if err := e.Shutdown(context.Background()); err != nil {
		logrus.Fatal("Error shutting down server:", err)
	}

	logrus.Info("Server shut down gracefully")
}