package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/config"
	"github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/routes"
)

func main() {
	r := gin.Default()
	config.ConnectDB()
	routes.RegisterRoutes(r)
	r.Run(":8080")
}
