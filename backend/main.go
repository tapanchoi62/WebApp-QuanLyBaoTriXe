package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/config"
	_ "github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/docs"
	"github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/models"
	"github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/routes"
	"github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/services"
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
	services.StartCronJobs()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	config.Connect()
	config.AutoMigrate(
		&models.Role{},
		&models.Permission{},
		&models.User{},
		&models.Vehicle{},
		&models.Item{},
		&models.Warehouse{},
		&models.Stock{},
		&models.StockLog{},
		&models.MaintenanceRequest{},
		&models.MaintenanceRequestItem{},
		&models.MaintenanceRecord{},
		&models.MaintenanceRecordItem{},
		&models.RolePermission{},
	)
	config.SeedRBAC(config.DB)

	routes.RegisterRoutes(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", PingHandler)
	r.Run(":8080")
}

func PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}
