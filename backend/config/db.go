package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		getEnv("DB_USER", "root"),
		getEnv("DB_PASSWORD", "123456"),
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "3306"),
		getEnv("DB_NAME", "fleet"),
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}
	log.Println("✅ Database connected!")
}

func AutoMigrate(models ...interface{}) {
	err := DB.AutoMigrate(models...)
	if err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}
	log.Println("✅ Auto migration completed!")
	log.Println("sucess 2")
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
