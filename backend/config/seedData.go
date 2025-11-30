package config

import (
	"github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/models"
	"gorm.io/gorm"
)

func SeedRBAC(db *gorm.DB) {
	var count int64
	db.Model(&models.Role{}).Count(&count)
	if count > 0 {
		return // đã seed trước đó
	}

	roles := []models.Role{
		{Name: "Admin"},
		{Name: "Technician"},
		{Name: "Warehouse"},
	}
	db.Create(&roles)

	perms := []models.Permission{
		{Name: "CREATE_REQUEST"},
		{Name: "APPROVE_REQUEST"},
		{Name: "IN_STOCK"},
		{Name: "OUT_STOCK"},
		{Name: "ADJUST_STOCK"},
	}
	db.Create(&perms)

	// Admin full quyền
	admin := roles[0]
	db.Model(&admin).Association("Permissions").Append(&perms)

	// Technician
	technician := roles[1]
	db.Model(&technician).Association("Permissions").Append(&perms[0]) // CREATE_REQUEST

	// Warehouse
	warehouse := roles[2]
	db.Model(&warehouse).Association("Permissions").Append(perms[2], perms[3], perms[4])
}
