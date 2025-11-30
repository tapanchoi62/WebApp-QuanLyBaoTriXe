package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/config"
	"github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/models"
)

type CreateRoleReq struct {
	Name          string `json:"name" binding:"required"`
	PermissionIDs []uint `json:"permissionIds"`
}

func GetRoles(c *gin.Context) {
	var roles []models.Role
	if err := config.DB.Preload("Permissions").Find(&roles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể lấy danh sách role"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": roles,
	})
}

func CreateRole(c *gin.Context) {
	var req CreateRoleReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role := models.Role{Name: req.Name}

	if err := config.DB.Create(&role).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể tạo role"})
		return
	}

	if len(req.PermissionIDs) > 0 {
		var perms []models.Permission
		config.DB.Where("id IN ?", req.PermissionIDs).Find(&perms)

		config.DB.Model(&role).Association("Permissions").Replace(&perms)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": role,
	})
}

func UpdateRole(c *gin.Context) {
	id := c.Param("id")

	var req CreateRoleReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var role models.Role
	if err := config.DB.Preload("Permissions").First(&role, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role không tồn tại"})
		return
	}

	role.Name = req.Name
	config.DB.Save(&role)

	// Update permission list
	var perms []models.Permission
	config.DB.Where("id IN ?", req.PermissionIDs).Find(&perms)
	config.DB.Model(&role).Association("Permissions").Replace(&perms)

	c.JSON(http.StatusOK, gin.H{
		"data": role,
	})
}

func DeleteRole(c *gin.Context) {
	id := c.Param("id")

	// NOTE: nếu User đang dùng Role này, nên chặn
	if err := config.DB.Delete(&models.Role{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể xóa role"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Đã xóa role"})
}
