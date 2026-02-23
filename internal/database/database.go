package database

import (
	"log"
	"nexample/internal/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	cfg := config.Load()
	var err error

	DB, err = gorm.Open(mysql.Open(cfg.DSN), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	sqlDB, _ := DB.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
}

func Migrate() {
	err := DB.AutoMigrate(
		&Account{},
	)
	if err != nil {
		log.Fatal("failed to migrate:", err)
	}
	log.Println("migration completed")
}
