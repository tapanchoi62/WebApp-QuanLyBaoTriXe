package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/config"
	"github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/models"
)

func GetStocks(c *gin.Context) {
	warehouseID := c.Query("warehouse_id")

	db := config.DB.Preload("Item").Preload("Warehouse")
	if warehouseID != "" {
		db = db.Where("warehouse_id = ?", warehouseID)
	}

	var stocks []models.Stock
	if err := db.Find(&stocks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": stocks})
}

func GetStocksByWarehouse(c *gin.Context) {
	id := c.Param("id")
	var stocks []models.Stock
	if err := config.DB.Preload("Item").Where("warehouse_id = ?", id).Find(&stocks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": stocks})
}

type StockTransactionInput struct {
	StockID     *uint   `json:"stock_id"`
	ItemID      uint    `json:"item_id" binding:"required"`
	WarehouseID uint    `json:"warehouse_id" binding:"required"`
	Type        string  `json:"type" binding:"required"`
	Quantity    float64 `json:"quantity" binding:"required"`
	Note        string  `json:"note"`
	SupplierID  *uint   `json:"supplier_id"`
}

func StockTransaction(c *gin.Context) {
	var input StockTransactionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find or create stock record for (item_id, warehouse_id)
	var stock models.Stock
	result := config.DB.Where("item_id = ? AND warehouse_id = ?", input.ItemID, input.WarehouseID).First(&stock)
	if result.Error != nil {
		// Create new stock record
		stock = models.Stock{
			ItemID:      input.ItemID,
			WarehouseID: input.WarehouseID,
			Quantity:    0,
		}
		if err := config.DB.Create(&stock).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	// Apply transaction
	switch input.Type {
	case "IN":
		stock.Quantity += input.Quantity
	case "OUT":
		if stock.Quantity < input.Quantity {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Not enough stock"})
			return
		}
		stock.Quantity -= input.Quantity
	case "ADJUST":
		stock.Quantity = input.Quantity
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction type. Must be IN, OUT, or ADJUST"})
		return
	}

	if err := config.DB.Save(&stock).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create stock log
	log := models.StockLog{
		StockID:    stock.ID,
		Type:       input.Type,
		Quantity:   input.Quantity,
		Note:       input.Note,
		SupplierID: input.SupplierID,
	}
	if err := config.DB.Create(&log).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"stock": stock,
		"log":   log,
	})
}
