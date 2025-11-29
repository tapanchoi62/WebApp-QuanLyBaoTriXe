package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/config"
	"github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/models"
)

// Get all logs
// @Summary Get all inventory logs
// @Tags InventoryLog
// @Produce json
// @Success 200 {array} models.InventoryLog
// @Router /inventory-log [get]
func GetInventoryLogs(c *gin.Context) {
	var logs []models.InventoryLog
	config.DB.Find(&logs)
	c.JSON(http.StatusOK, logs)
}

// Create log
// @Summary Create inventory log
// @Tags InventoryLog
// @Accept json
// @Produce json
// @Param data body models.InventoryLog true "InventoryLog"
// @Success 200 {object} models.InventoryLog
// @Router /inventory-log [post]
func CreateInventoryLog(c *gin.Context) {
	var input models.InventoryLog
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log := models.InventoryLog{
		InventoryID: input.InventoryID,
		Action:      input.Action,
		Quantity:    input.Quantity,
		Note:        input.Note,
	}

	config.DB.Create(&log)
	c.JSON(http.StatusOK, log)
}

// Get 1 log
// @Summary Get one log
// @Tags InventoryLog
// @Produce json
// @Param id path int true "Log ID"
// @Success 200 {object} models.InventoryLog
// @Router /inventory-log/{id} [get]
func GetInventoryLog(c *gin.Context) {
	id := c.Param("id")
	var log models.InventoryLog

	if err := config.DB.First(&log, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.JSON(http.StatusOK, log)
}

// Update log
// @Summary Update inventory log
// @Tags InventoryLog
// @Accept json
// @Produce json
// @Param id path int true "Log ID"
// @Param data body models.InventoryLog true "InventoryLog"
// @Success 200 {object} models.InventoryLog
// @Router /inventory-log/{id} [put]
func UpdateInventoryLog(c *gin.Context) {
	id := c.Param("id")
	var log models.InventoryLog

	if err := config.DB.First(&log, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	var input models.InventoryLog
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Action = input.Action
	log.Quantity = input.Quantity
	log.Note = input.Note
	log.InventoryID = input.InventoryID

	config.DB.Save(&log)
	c.JSON(http.StatusOK, log)
}

// Delete log
// @Summary Delete inventory log
// @Tags InventoryLog
// @Param id path int true "Log ID"
// @Success 200
// @Router /inventory-log/{id} [delete]
func DeleteInventoryLog(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&models.InventoryLog{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
