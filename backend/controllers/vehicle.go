package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/config"
	"github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/models"
)

// GetVehicles godoc
// @Summary Get all vehicles
// @Description Get list of all vehicles
// @Tags Vehicles
// @Accept json
// @Produce json
// @Success 200 {array} map[string]interface{}
// @Security BearerAuth
// @Router /api/vehicles [get]
func GetVehicles(c *gin.Context) {
	var vehicles []models.Vehicle
	config.DB.Find(&vehicles)
	c.JSON(http.StatusOK, vehicles)
}

// GetVehicle godoc
// @Summary Get vehicle by ID
// @Description Get a vehicle by its ID
// @Tags Vehicles
// @Accept json
// @Produce json
// @Param id path int true "Vehicle ID"
// @Success 200 {object} models.Vehicle
// @Failure 404 {object} map[string]string
// @Security BearerAuth
// @Router /api/vehicles/{id} [get]
func GetVehicle(c *gin.Context) {
	id := c.Param("id")
	var vehicle models.Vehicle
	if err := config.DB.First(&vehicle, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vehicle not found"})
		return
	}
	c.JSON(http.StatusOK, vehicle)
}

// CreateVehicle godoc
// @Summary Create a new vehicle
// @Description Create vehicle by JSON body
// @Tags Vehicles
// @Accept json
// @Produce json
// @Param vehicle body models.Vehicle true "Vehicle info"
// @Success 201 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/vehicles [post]
func CreateVehicle(c *gin.Context) {
	var input models.Vehicle
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&input)
	c.JSON(http.StatusOK, input)
}

// UpdateVehicle godoc
// @Summary Update a vehicle
// @Description Update vehicle by ID
// @Tags Vehicles
// @Accept json
// @Produce json
// @Param id path int true "Vehicle ID"
// @Param vehicle body models.Vehicle true "Vehicle info"
// @Success 200 {object} models.Vehicle
// @Failure 404 {object} map[string]string
// @Security BearerAuth
// @Router /api/vehicles/{id} [put]
func UpdateVehicle(c *gin.Context) {
	id := c.Param("id")
	var vehicle models.Vehicle
	if err := config.DB.First(&vehicle, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vehicle not found"})
		return
	}

	var input models.Vehicle
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&vehicle).Updates(input)
	c.JSON(http.StatusOK, vehicle)
}

// DeleteVehicle godoc
// @Summary Delete a vehicle
// @Description Delete vehicle by ID
// @Tags Vehicles
// @Accept json
// @Produce json
// @Param id path int true "Vehicle ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Security BearerAuth
// @Router /api/vehicles/{id} [delete]
func DeleteVehicle(c *gin.Context) {
	id := c.Param("id")
	var vehicle models.Vehicle
	if err := config.DB.First(&vehicle, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vehicle not found"})
		return
	}

	config.DB.Delete(&vehicle)
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}
