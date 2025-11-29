package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/config"
	"github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/models"
)

// Get all logs
// @Summary Get all maintenance logs
// @Tags MaintenanceLog
// @Produce json
// @Success 200 {array} models.MaintenanceLog
// @Router /maintenance-logs [get]
func GetMaintenanceLogs(c *gin.Context) {
	var logs []models.MaintenanceLog
	config.DB.Order("id desc").Find(&logs)
	c.JSON(http.StatusOK, logs)
}

// Get log by ID
// @Summary Get maintenance log by ID
// @Tags MaintenanceLog
// @Param id path int true "ID"
// @Produce json
// @Success 200 {object} models.MaintenanceLog
// @Router /maintenance-logs/{id} [get]
func GetMaintenanceLog(c *gin.Context) {
	id := c.Param("id")
	var log models.MaintenanceLog

	if err := config.DB.First(&log, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Log not found"})
		return
	}

	c.JSON(http.StatusOK, log)
}

// Create log
// @Summary Create maintenance log
// @Tags MaintenanceLog
// @Accept json
// @Produce json
// @Param log body models.MaintenanceLog true "Maintenance Log"
// @Success 201 {object} models.MaintenanceLog
// @Router /maintenance-logs [post]
func CreateMaintenanceLog(c *gin.Context) {
	var input models.MaintenanceLog

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&input)
	c.JSON(http.StatusCreated, input)
}

// Update log
// @Summary Update maintenance log
// @Tags MaintenanceLog
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param log body models.MaintenanceLog true "Maintenance Log"
// @Success 200 {object} models.MaintenanceLog
// @Router /maintenance-logs/{id} [put]
func UpdateMaintenanceLog(c *gin.Context) {
	id := c.Param("id")
	var log models.MaintenanceLog

	if err := config.DB.First(&log, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Log not found"})
		return
	}

	var input models.MaintenanceLog
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Action = input.Action
	log.Description = input.Description
	log.Cost = input.Cost
	log.CreatedBy = input.CreatedBy

	config.DB.Save(&log)
	c.JSON(http.StatusOK, log)
}

// Delete log
// @Summary Delete maintenance log
// @Tags MaintenanceLog
// @Param id path int true "ID"
// @Success 200 {string} string "deleted"
// @Router /maintenance-logs/{id} [delete]
func DeleteMaintenanceLog(c *gin.Context) {
	id := c.Param("id")
	var log models.MaintenanceLog

	if err := config.DB.First(&log, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Log not found"})
		return
	}

	config.DB.Delete(&log)
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
