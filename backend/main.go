package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/config"
	_ "github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/docs"
	"github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/models"
	"github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/routes"
)

// @title Fleet Management API
// @version 1.0
// @description API for fleet management
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	r := gin.Default()
	config.Connect()
	config.AutoMigrate(
		&models.User{},
		&models.Vehicle{},
		&models.Inventory{},
		&models.MaintenanceLog{},
		&models.MaintenanceRequest{},
		&models.MaintenanceRequestItem{},
		&models.Item{},
		&models.InventoryLog{},
	)
	routes.RegisterRoutes(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", PingHandler)
	r.Run(":8080")
}

func PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}
