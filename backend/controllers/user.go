package controllers

import (
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/config"
	"github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/models"
)

func GetUsers(c *gin.Context) {
	var users []models.User

	// Phân trang
	page := 1
	pageSize := 10
	if p, ok := c.GetQuery("page"); ok {
		page, _ = strconv.Atoi(p)
	}
	if ps, ok := c.GetQuery("pageSize"); ok {
		pageSize, _ = strconv.Atoi(ps)
	}
	offset := (page - 1) * pageSize

	var total int64
	config.DB.Model(&models.User{}).Count(&total)

	// Preload Role + Permissions
	err := config.DB.Preload("Role").Preload("Role.Permissions").
		Limit(pageSize).Offset(offset).
		Find(&users).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
		"pagination": gin.H{
			"page":       page,
			"pageSize":   pageSize,
			"total":      total,
			"totalPages": int(math.Ceil(float64(total) / float64(pageSize))),
		},
	})
}

func GetUser(c *gin.Context) {
	id := c.Param("id")

	var u models.User
	// Preload Role + Permissions
	if err := config.DB.Preload("Role").Preload("Role.Permissions").First(&u, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy user"})
		return
	}

	// Build permissions slice
	perms := []string{}
	for _, p := range u.Role.Permissions {
		perms = append(perms, p.Name)
	}

	resp := struct {
		ID          uint        `json:"id"`
		Username    string      `json:"username"`
		Role        models.Role `json:"role"`
		Permissions []string    `json:"permissions"`
	}{
		ID:          u.ID,
		Username:    u.Username,
		Role:        u.Role,
		Permissions: perms,
	}

	c.JSON(http.StatusOK, gin.H{
		"data": resp,
	})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	if err := config.DB.Preload("Role").Preload("Role.Permissions").First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User không tồn tại"})
		return
	}

	var input struct {
		RoleID uint `json:"role_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Model(&models.User{}).Where("id = ?", id).Update("role_id", input.RoleID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cập nhật role thất bại"})
		return
	}

	// Reload role + permission
	if err := config.DB.Preload("Role").Preload("Role.Permissions").First(&user, user.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không load lại role"})
		return
	}

	perms := []string{}
	for _, p := range user.Role.Permissions {
		perms = append(perms, p.Name)
	}

	c.JSON(http.StatusOK, gin.H{
		"id":          user.ID,
		"username":    user.Username,
		"role":        user.Role.Name,
		"permissions": perms,
	})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	if err := config.DB.Delete(&models.User{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể xóa user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Xóa thành công"})
}
