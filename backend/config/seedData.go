package config

import (
	"log"

	"github.com/tapanchoi62/WebApp-QuanLyBaoTriXe/backend/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedRBAC(db *gorm.DB) {
	var count int64
	db.Model(&models.Role{}).Count(&count)
	if count > 0 {
		return
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

func SeedInitData(db *gorm.DB) {
	seedUsers(db)
	seedVehicles(db)
	seedInventory(db)
}

func hashPass(password string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("❌ Failed to hash password: %v", err)
	}
	return string(hashed)
}

func seedUsers(db *gorm.DB) {
	var count int64
	db.Model(&models.User{}).Count(&count)
	if count > 0 {
		return
	}

	var adminRole, techRole, warehouseRole models.Role
	db.Where("name = ?", "Admin").First(&adminRole)
	db.Where("name = ?", "Technician").First(&techRole)
	db.Where("name = ?", "Warehouse").First(&warehouseRole)

	users := []models.User{
		{Username: "admin", Password: hashPass("Admin@123"), RoleID: adminRole.ID},
		{Username: "kyThuat1", Password: hashPass("KyThuat@123"), RoleID: techRole.ID},
		{Username: "thuKho1", Password: hashPass("ThuKho@123"), RoleID: warehouseRole.ID},
	}
	db.Create(&users)
	log.Println("✅ Seeded users")
}

func seedVehicles(db *gorm.DB) {
	var count int64
	db.Model(&models.Vehicle{}).Count(&count)
	if count > 0 {
		return
	}

	vehicles := []models.Vehicle{
		{PlateNumber: "51A-12345", Model: "Toyota Hiace", Year: 2020, Note: "Xe tải nhẹ tuyến nội thành"},
		{PlateNumber: "51B-67890", Model: "Ford Transit", Year: 2019, Note: "Xe chở hàng liên tỉnh"},
		{PlateNumber: "51C-11111", Model: "Isuzu NPR", Year: 2021, Note: "Xe tải trung"},
		{PlateNumber: "51D-22222", Model: "Hyundai Porter", Year: 2018, Note: "Xe giao hàng khu công nghiệp"},
		{PlateNumber: "51F-33333", Model: "Mitsubishi Fuso", Year: 2022, Note: "Xe tải nặng tuyến dài"},
	}
	db.Create(&vehicles)
	log.Println("✅ Seeded vehicles")
}

func seedInventory(db *gorm.DB) {
	var itemCount int64
	db.Model(&models.Item{}).Count(&itemCount)
	if itemCount > 0 {
		return
	}

	// Suppliers
	suppliers := []models.Supplier{
		{Name: "Công ty TNHH Phụ Tùng Minh Thành", Contact: "0901234567"},
		{Name: "Đại lý Lốp Xe An Phát", Contact: "0912345678"},
		{Name: "Kho Ắc Quy Hùng Thịnh", Contact: "0923456789"},
	}
	db.Create(&suppliers)

	// Warehouses
	warehouses := []models.Warehouse{
		{Name: "Kho Trung Tâm", Location: "TP. Hồ Chí Minh"},
		{Name: "Kho Chi Nhánh", Location: "Bình Dương"},
	}
	db.Create(&warehouses)

	// Items (vật tư phổ biến)
	items := []models.Item{
		{Name: "Lốp xe 195/65R15", Category: "Tire", Unit: "cái"},
		{Name: "Lốp xe 225/75R16", Category: "Tire", Unit: "cái"},
		{Name: "Bình ắc quy 12V 70Ah", Category: "Battery", Unit: "bình"},
		{Name: "Bình ắc quy 12V 100Ah", Category: "Battery", Unit: "bình"},
		{Name: "Dầu động cơ 5W-30", Category: "Oil", Unit: "lít"},
		{Name: "Dầu hộp số", Category: "Oil", Unit: "lít"},
		{Name: "Lọc dầu động cơ", Category: "Filter", Unit: "cái"},
		{Name: "Lọc khí", Category: "Filter", Unit: "cái"},
		{Name: "Lọc nhiên liệu", Category: "Filter", Unit: "cái"},
		{Name: "Má phanh trước", Category: "Brake", Unit: "bộ"},
		{Name: "Má phanh sau", Category: "Brake", Unit: "bộ"},
		{Name: "Bóng đèn pha H4", Category: "Lighting", Unit: "cái"},
		{Name: "Dây curoa", Category: "Belt", Unit: "sợi"},
		{Name: "Bơm nước làm mát", Category: "Cooling", Unit: "cái"},
		{Name: "Nhớt phanh DOT4", Category: "Fluid", Unit: "lít"},
	}
	db.Create(&items)

	// Stock tại Kho Trung Tâm
	khoTrungTam := warehouses[0]
	stockKhoTrungTam := []models.Stock{
		{ItemID: items[0].ID, WarehouseID: khoTrungTam.ID, Quantity: 20},  // Lốp 195/65R15
		{ItemID: items[1].ID, WarehouseID: khoTrungTam.ID, Quantity: 12},  // Lốp 225/75R16
		{ItemID: items[2].ID, WarehouseID: khoTrungTam.ID, Quantity: 8},   // Ắc quy 70Ah
		{ItemID: items[3].ID, WarehouseID: khoTrungTam.ID, Quantity: 5},   // Ắc quy 100Ah
		{ItemID: items[4].ID, WarehouseID: khoTrungTam.ID, Quantity: 100}, // Dầu 5W-30
		{ItemID: items[5].ID, WarehouseID: khoTrungTam.ID, Quantity: 40},  // Dầu hộp số
		{ItemID: items[6].ID, WarehouseID: khoTrungTam.ID, Quantity: 30},  // Lọc dầu
		{ItemID: items[7].ID, WarehouseID: khoTrungTam.ID, Quantity: 25},  // Lọc khí
		{ItemID: items[8].ID, WarehouseID: khoTrungTam.ID, Quantity: 20},  // Lọc nhiên liệu
		{ItemID: items[9].ID, WarehouseID: khoTrungTam.ID, Quantity: 10},  // Má phanh trước
		{ItemID: items[10].ID, WarehouseID: khoTrungTam.ID, Quantity: 10}, // Má phanh sau
		{ItemID: items[11].ID, WarehouseID: khoTrungTam.ID, Quantity: 50}, // Bóng đèn H4
		{ItemID: items[12].ID, WarehouseID: khoTrungTam.ID, Quantity: 15}, // Dây curoa
		{ItemID: items[13].ID, WarehouseID: khoTrungTam.ID, Quantity: 6},  // Bơm nước
		{ItemID: items[14].ID, WarehouseID: khoTrungTam.ID, Quantity: 20}, // Nhớt phanh
	}
	db.Create(&stockKhoTrungTam)

	// Stock tại Kho Chi Nhánh (chỉ vật tư thiết yếu)
	khoChiNhanh := warehouses[1]
	stockKhoChiNhanh := []models.Stock{
		{ItemID: items[0].ID, WarehouseID: khoChiNhanh.ID, Quantity: 8},  // Lốp 195/65R15
		{ItemID: items[2].ID, WarehouseID: khoChiNhanh.ID, Quantity: 4},  // Ắc quy 70Ah
		{ItemID: items[4].ID, WarehouseID: khoChiNhanh.ID, Quantity: 50}, // Dầu 5W-30
		{ItemID: items[6].ID, WarehouseID: khoChiNhanh.ID, Quantity: 15}, // Lọc dầu
		{ItemID: items[9].ID, WarehouseID: khoChiNhanh.ID, Quantity: 5},  // Má phanh trước
	}
	db.Create(&stockKhoChiNhanh)

	log.Println("✅ Seeded inventory (suppliers, warehouses, items, stock)")
}
