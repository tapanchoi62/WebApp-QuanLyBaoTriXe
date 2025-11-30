package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/config"
	"github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/models"
)

func UploadFile(c *gin.Context) {
	entityId := c.Query("entityId")     // Có thể để trống lúc upload file độc lập
	entityType := c.Query("entityType") // Có thể để trống

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File không hợp lệ"})
		return
	}

	// Tạo tên file unique
	ext := filepath.Ext(file.Filename)
	newName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	savePath := filepath.Join("uploads", newName)

	// Lưu file vào ổ đĩa
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể lưu file"})
		return
	}

	// Tạo object lưu DB
	fileRecord := models.File{
		EntityID:   entityId,
		EntityType: entityType,
		FileURL:    "/uploads/" + newName,
		FileName:   file.Filename,
		MimeType:   file.Header.Get("Content-Type"),
		Size:       file.Size,
		CreatedAt:  time.Now(),
	}

	// Lưu vào database
	if err := config.DB.Create(&fileRecord).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể lưu dữ liệu file vào DB"})
		return
	}

	// Trả fileId và URL để frontend lưu hoặc gửi sang bảng khác
	c.JSON(http.StatusOK, gin.H{
		"message": "Upload thành công",
		"file": gin.H{
			"id":       fileRecord.ID,
			"url":      fileRecord.FileURL,
			"fileName": fileRecord.FileName,
		},
	})
}
