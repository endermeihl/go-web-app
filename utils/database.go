package utils

import (
	"go-web-app/config"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := config.Config.GetString("DB_USER") + ":" +
		config.Config.GetString("DB_PASSWORD") + "@tcp(" +
		config.Config.GetString("DB_HOST") + ":" +
		config.Config.GetString("DB_PORT") + ")/" +
		config.Config.GetString("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
}
