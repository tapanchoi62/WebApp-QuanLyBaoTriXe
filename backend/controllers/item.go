package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/config"
	"github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/models"
	"github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/utils"
)

// GetItem godoc
// @Summary Get all Item
// @Description Get list of all Item
// @Tags Item
// @Accept json
// @Produce json
// @Success 200 {array} map[string]interface{}
// @Security BearerAuth
// @Router /api/items [get]
func GetItems(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "4"))
	search := c.Query("search")

	data, paging, err := utils.Paginate[models.Item](config.DB.Model(&models.Item{}), page, pageSize, search, []string{"name", "category"})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":       data,
		"pagination": paging,
	})
}

// GetItem godoc
// @Summary Get Item by ID
// @Description Get a Item by its ID
// @Tags Item
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} models.Item
// @Failure 404 {object} map[string]string
// @Security BearerAuth
// @Router /api/items/{id} [get]
func GetItem(c *gin.Context) {
	id := c.Param("id")
	var Item models.Item
	if err := config.DB.First(&Item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": Item,
	})
}

// CreateItem godoc
// @Summary Create a new Item
// @Description Create Item by JSON body
// @Tags Item
// @Accept json
// @Produce json
// @Param Item body models.Item true "Item info"
// @Success 201 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/items [post]
func CreateItem(c *gin.Context) {
	var input models.Item
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&input)
	c.JSON(http.StatusOK, input)
}

// UpdateItem godoc
// @Summary Update a Item
// @Description Update Item by ID
// @Tags Item
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Param Item body models.Item true "Item info"
// @Success 200 {object} models.Item
// @Failure 404 {object} map[string]string
// @Security BearerAuth
// @Router /api/items/{id} [put]
func UpdateItem(c *gin.Context) {
	id := c.Param("id")
	var item models.Item
	if err := config.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	var input models.Item
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&item).Updates(input)
	c.JSON(http.StatusOK, item)

}

// DeleteItem godoc
// @Summary Delete a Item
// @Description Delete Item by ID
// @Tags Item
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Security BearerAuth
// @Router /api/items/{id} [delete]
func DeleteItem(c *gin.Context) {
	id := c.Param("id")
	var item models.Item
	if err := config.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	config.DB.Delete(&item)
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}
