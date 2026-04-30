package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/config"
	"github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/models"
)

func GetWarehouses(c *gin.Context) {
	var warehouses []models.Warehouse
	if err := config.DB.Find(&warehouses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": warehouses})
}

func CreateWarehouse(c *gin.Context) {
	var input struct {
		Name     string `json:"name" binding:"required"`
		Location string `json:"location"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	warehouse := models.Warehouse{
		Name:     input.Name,
		Location: input.Location,
	}
	if err := config.DB.Create(&warehouse).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": warehouse})
}

func UpdateWarehouse(c *gin.Context) {
	id := c.Param("id")
	var warehouse models.Warehouse
	if err := config.DB.First(&warehouse, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Warehouse not found"})
		return
	}

	var input struct {
		Name     string `json:"name"`
		Location string `json:"location"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&warehouse).Updates(map[string]interface{}{
		"name":     input.Name,
		"location": input.Location,
	})
	c.JSON(http.StatusOK, gin.H{"data": warehouse})
}

func DeleteWarehouse(c *gin.Context) {
	id := c.Param("id")
	var warehouse models.Warehouse
	if err := config.DB.First(&warehouse, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Warehouse not found"})
		return
	}

	config.DB.Delete(&warehouse)
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}
