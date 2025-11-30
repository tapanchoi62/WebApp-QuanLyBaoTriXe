package services

import (
	"log"
	"os"
	"path/filepath"

	"github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/config"
	"github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/models"
)

func CleanupUnusedFiles() {
	uploadDir := "uploads"

	files, err := os.ReadDir(uploadDir)
	if err != nil {
		log.Println("Không thể đọc thư mục uploads:", err)
		return
	}

	for _, f := range files {
		filePath := filepath.Join(uploadDir, f.Name())
		// Kiểm tra file có nằm trong DB không
		var count int64

		config.DB.Model(&models.File{}).
			Where("file_url = ?", "/uploads/"+f.Name()).
			Count(&count)

			// Nếu không tìm thấy trong DB → xóa file
		if count == 0 {
			if err := os.Remove(filePath); err == nil {
				log.Println("Đã xoá file rác:", f.Name())
			} else {
				log.Println("Không thể xóa file:", f.Name(), err)
			}
		}
	}
}
