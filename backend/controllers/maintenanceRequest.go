package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/config"
	"github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/models"
)

// Get all
// @Summary Get all maintenance requests
// @Tags MaintenanceRequest
// @Produce json
// @Success 200 {array} models.MaintenanceRequest
// @Router /maintenance-requests [get]
func GetMaintenanceRequests(c *gin.Context) {
	var data []models.MaintenanceRequest
	config.DB.Find(&data)
	c.JSON(http.StatusOK, data)
}

// Get by ID
// @Summary Get maintenance request by ID
// @Tags MaintenanceRequest
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.MaintenanceRequest
// @Router /maintenance-requests/{id} [get]
func GetMaintenanceRequest(c *gin.Context) {
	id := c.Param("id")
	var item models.MaintenanceRequest

	if err := config.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}

	c.JSON(http.StatusOK, item)
}

// Create
// @Summary Create maintenance request
// @Tags MaintenanceRequest
// @Accept json
// @Produce json
// @Param request body models.MaintenanceRequest true "Request"
// @Success 201 {object} models.MaintenanceRequest
// @Router /maintenance-requests [post]
func CreateMaintenanceRequest(c *gin.Context) {
	var input models.MaintenanceRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.Status = "pending"

	config.DB.Create(&input)
	c.JSON(http.StatusCreated, input)
}

// Update
// @Summary Update maintenance request
// @Tags MaintenanceRequest
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param request body models.MaintenanceRequest true "Request"
// @Success 200 {object} models.MaintenanceRequest
// @Router /maintenance-requests/{id} [put]
func UpdateMaintenanceRequest(c *gin.Context) {
	id := c.Param("id")
	var item models.MaintenanceRequest

	if err := config.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}

	var input models.MaintenanceRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item.VehicleID = input.VehicleID
	item.Description = input.Description
	item.Status = input.Status

	config.DB.Save(&item)
	c.JSON(http.StatusOK, item)
}

// Delete
// @Summary Delete maintenance request
// @Tags MaintenanceRequest
// @Param id path int true "ID"
// @Success 200 {string} string "deleted"
// @Router /maintenance-requests/{id} [delete]
func DeleteMaintenanceRequest(c *gin.Context) {
	id := c.Param("id")
	var item models.MaintenanceRequest

	if err := config.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}

	config.DB.Delete(&item)
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
