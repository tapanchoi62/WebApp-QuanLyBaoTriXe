package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/config"
	"github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/models"
)

// Get all items
// @Summary Get all maintenance request items
// @Tags MaintenanceRequestItem
// @Produce json
// @Success 200 {array} models.MaintenanceRequestItem
// @Router /maintenance-request-items [get]
func GetMaintenanceRequestItems(c *gin.Context) {
	var data []models.MaintenanceRequestItem
	config.DB.Find(&data)
	c.JSON(http.StatusOK, data)
}

// Get by ID
// @Summary Get maintenance request item by ID
// @Tags MaintenanceRequestItem
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.MaintenanceRequestItem
// @Router /maintenance-request-items/{id} [get]
func GetMaintenanceRequestItem(c *gin.Context) {
	id := c.Param("id")
	var item models.MaintenanceRequestItem

	if err := config.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, item)
}

// Create
// @Summary Create maintenance request item
// @Tags MaintenanceRequestItem
// @Accept json
// @Produce json
// @Param item body models.MaintenanceRequestItem true "Item"
// @Success 201 {object} models.MaintenanceRequestItem
// @Router /maintenance-request-items [post]
func CreateMaintenanceRequestItem(c *gin.Context) {
	var input models.MaintenanceRequestItem

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.Status = "pending"

	config.DB.Create(&input)
	c.JSON(http.StatusCreated, input)
}

// Update
// @Summary Update maintenance request item
// @Tags MaintenanceRequestItem
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param item body models.MaintenanceRequestItem true "Item"
// @Success 200 {object} models.MaintenanceRequestItem
// @Router /maintenance-request-items/{id} [put]
func UpdateMaintenanceRequestItem(c *gin.Context) {
	id := c.Param("id")
	var item models.MaintenanceRequestItem

	if err := config.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	var input models.MaintenanceRequestItem
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item.ItemName = input.ItemName
	item.Cost = input.Cost
	item.Status = input.Status
	item.Note = input.Note

	config.DB.Save(&item)
	c.JSON(http.StatusOK, item)
}

// Delete
// @Summary Delete maintenance request item
// @Tags MaintenanceRequestItem
// @Param id path int true "ID"
// @Success 200 {string} string "deleted"
// @Router /maintenance-request-items/{id} [delete]
func DeleteMaintenanceRequestItem(c *gin.Context) {
	id := c.Param("id")
	var item models.MaintenanceRequestItem

	if err := config.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	config.DB.Delete(&item)
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
