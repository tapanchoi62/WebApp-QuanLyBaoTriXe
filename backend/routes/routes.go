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
		api.POST("/vehicles", controllers.CreateVehicle)
		api.PUT("/vehicles/:id", controllers.UpdateVehicle)
		api.DELETE("/vehicles/:id", controllers.DeleteVehicle)
	}
}
