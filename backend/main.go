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

func main() {
	r := gin.Default()

	// ✅ CORS middleware đúng cách
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	services.StartCronJobs()
	config.Connect()

	config.AutoMigrate(
		&models.Role{},
		&models.Permission{},
		&models.User{},
		&models.Vehicle{},
		&models.Supplier{},
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
	config.SeedInitData(config.DB)

	routes.RegisterRoutes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/ping", PingHandler)

	// ✅ đúng chuẩn gin
	r.Run(":8080")
}

func PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}
