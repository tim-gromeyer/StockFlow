package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"github.com/tim/StockFlow/models"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("stockflow.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB.AutoMigrate(&models.User{}, &models.Portfolio{})

	// Create a dummy user for testing
	var user models.User
	if err := DB.First(&user, 1).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			DB.Create(&models.User{ID: 1, CashBalance: 10000})
		}
	}
}
