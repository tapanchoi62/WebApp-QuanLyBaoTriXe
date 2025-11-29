package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/controllers"
	"github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/middleware"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	r.POST("/api/login", controllers.Login)
	r.POST("/api/register", controllers.RegisterUser)
	api.Use(middleware.AuthRequired())
	{
		api.GET("/vehicles", controllers.GetVehicles)
		api.GET("/vehicles/:id", controllers.GetVehicle)
		api.POST("/vehicles", controllers.CreateVehicle)
		api.PUT("/vehicles/:id", controllers.UpdateVehicle)
		api.DELETE("/vehicles/:id", controllers.DeleteVehicle)

		api.GET("/items", controllers.GetItems)
		api.GET("/items/:id", controllers.GetItem)
		api.POST("/items", controllers.CreateItem)
		api.PUT("/items/:id", controllers.UpdateItem)
		api.DELETE("/items/:id", controllers.DeleteItem)

		api.GET("/inventory", controllers.GetInventories)
		api.POST("/inventory", controllers.CreateInventory)
		api.GET("/inventory/:id", controllers.GetInventory)
		api.PUT("/inventory/:id", controllers.UpdateInventory)
		api.DELETE("/inventory/:id", controllers.DeleteInventory)

		api.GET("/inventory-log", controllers.GetInventoryLogs)
		api.POST("/inventory-log", controllers.CreateInventoryLog)
		api.GET("/inventory-log/:id", controllers.GetInventoryLog)
		api.PUT("/inventory-log/:id", controllers.UpdateInventoryLog)
		api.DELETE("/inventory-log/:id", controllers.DeleteInventoryLog)

		api.GET("/maintenance-requests", controllers.GetMaintenanceRequests)
		api.GET("/maintenance-requests/:id", controllers.GetMaintenanceRequest)
		api.POST("/maintenance-requests", controllers.CreateMaintenanceRequest)
		api.PUT("/maintenance-requests/:id", controllers.UpdateMaintenanceRequest)
		api.DELETE("/maintenance-requests/:id", controllers.DeleteMaintenanceRequest)

		api.GET("/maintenance-request-items", controllers.GetMaintenanceRequestItems)
		api.GET("/maintenance-request-items/:id", controllers.GetMaintenanceRequestItem)
		api.POST("/maintenance-request-items", controllers.CreateMaintenanceRequestItem)
		api.PUT("/maintenance-request-items/:id", controllers.UpdateMaintenanceRequestItem)
		api.DELETE("/maintenance-request-items/:id", controllers.DeleteMaintenanceRequestItem)

		api.GET("/maintenance-logs", controllers.GetMaintenanceLogs)
		api.GET("/maintenance-logs/:id", controllers.GetMaintenanceLog)
		api.POST("/maintenance-logs", controllers.CreateMaintenanceLog)
		api.PUT("/maintenance-logs/:id", controllers.UpdateMaintenanceLog)
		api.DELETE("/maintenance-logs/:id", controllers.DeleteMaintenanceLog)
	}
}
