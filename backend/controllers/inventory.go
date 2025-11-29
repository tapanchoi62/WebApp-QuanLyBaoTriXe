package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/config"
	"github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/models"
)

// GetInventories godoc
// @Summary Lấy danh sách tồn kho
// @Description Trả về toàn bộ vật tư trong kho
// @Tags Inventory
// @Produce json
// @Success 200 {array} models.Inventory
// @Router /inventory [get]
func GetInventories(c *gin.Context) {
	var inventories []models.Inventory
	config.DB.Find(&inventories)
	c.JSON(http.StatusOK, inventories)
}

// GetInventory godoc
// @Summary Lấy chi tiết vật tư
// @Tags Inventory
// @Produce json
// @Param id path int true "Inventory ID"
// @Success 200 {object} models.Inventory
// @Router /inventory/{id} [get]
func GetInventory(c *gin.Context) {
	id := c.Param("id")
	var inventory models.Inventory

	if err := config.DB.First(&inventory, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Inventory not found"})
		return
	}
	c.JSON(http.StatusOK, inventory)
}

// CreateInventory godoc
// @Summary Tạo mới vật tư
// @Tags Inventory
// @Accept json
// @Produce json
// @Param request body models.Inventory true "Inventory body"
// @Success 201 {object} models.Inventory
// @Router /inventory [post]
func CreateInventory(c *gin.Context) {
	var input models.Inventory

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	inventory := models.Inventory{
		Warehouse: input.Warehouse,
		ItemCode:  input.ItemCode,
		ItemName:  input.ItemName,
		Quantity:  input.Quantity,
		Unit:      input.Unit,
		MinStock:  input.MinStock,
	}

	config.DB.Create(&inventory)
	c.JSON(http.StatusCreated, inventory)
}

// UpdateInventory godoc
// @Summary Cập nhật vật tư
// @Tags Inventory
// @Accept json
// @Produce json
// @Param id path int true "Inventory ID"
// @Param request body models.Inventory true "Inventory body"
// @Success 200 {object} models.Inventory
// @Router /inventory/{id} [put]
func UpdateInventory(c *gin.Context) {
	id := c.Param("id")
	var inventory models.Inventory

	if err := config.DB.First(&inventory, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Inventory not found"})
		return
	}

	var input models.Inventory
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&inventory).Updates(input)
	c.JSON(http.StatusOK, inventory)
}

// DeleteInventory godoc
// @Summary Xóa vật tư
// @Tags Inventory
// @Param id path int true "Inventory ID"
// @Success 200 {object} map[string]string
// @Router /inventory/{id} [delete]
func DeleteInventory(c *gin.Context) {
	id := c.Param("id")
	var inventory models.Inventory

	if err := config.DB.First(&inventory, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Inventory not found"})
		return
	}

	config.DB.Delete(&inventory)
	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}
