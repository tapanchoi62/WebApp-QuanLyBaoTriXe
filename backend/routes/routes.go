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

		api.GET("/roles", controllers.GetRoles)
		api.POST("/roles", controllers.CreateRole)
		api.PUT("/roles/:id", controllers.UpdateRole)
		api.DELETE("/roles/:id", controllers.DeleteRole)

		api.GET("/permissions", controllers.GetPermissions)
		api.POST("/permissions", controllers.CreatePermission)
		api.PUT("/permissions/:id", controllers.UpdatePermission)
		api.DELETE("/permissions/:id", controllers.DeletePermission)

		api.GET("/users", controllers.GetUsers)
		api.GET("/users/:id", controllers.GetUser)
		api.POST("/users", controllers.RegisterUser)
		api.PUT("/users/:id", controllers.UpdateUser)
		api.DELETE("/users/:id", controllers.DeleteUser)

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
	}
}
