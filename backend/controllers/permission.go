package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/config"
	"github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/models"
)

func GetPermissions(c *gin.Context) {
	var permissions []models.Permission
	if err := config.DB.Find(&permissions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể lấy danh sách permission"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": permissions,
	})
}

type CreatePermissionReq struct {
	Name string `json:"name" binding:"required"`
}

func CreatePermission(c *gin.Context) {
	var req CreatePermissionReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	permission := models.Permission{Name: req.Name}

	if err := config.DB.Create(&permission).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể tạo permission"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": permission,
	})
}

func UpdatePermission(c *gin.Context) {
	id := c.Param("id")

	var req CreatePermissionReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var permission models.Permission
	if err := config.DB.First(&permission, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Permission không tồn tại"})
		return
	}

	permission.Name = req.Name
	config.DB.Save(&permission)

	c.JSON(http.StatusOK, gin.H{
		"data": permission,
	})
}

func DeletePermission(c *gin.Context) {
	id := c.Param("id")

	if err := config.DB.Delete(&models.Permission{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể xóa permission"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Đã xóa"})
}
