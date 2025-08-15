package services

import (
	"fmt"

	"github.com/tim/StockFlow/database"
	"github.com/tim/StockFlow/models"
)

func GetPortfolio(userID uint) ([]models.Portfolio, float64, error) {
	var portfolio []models.Portfolio
	if err := database.DB.Where("user_id = ?", userID).Find(&portfolio).Error; err != nil {
		return nil, 0, fmt.Errorf("user not found")
	}

	totalValue := 0.0
	for _, item := range portfolio {
		price, err := GetStockPrice(item.StockSymbol)
		if err != nil {
			return nil, 0, err
		}
		totalValue += float64(item.Quantity) * price
	}

	return portfolio, totalValue, nil
}

func GetUserBalance(userID uint) (float64, error) {
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return 0, fmt.Errorf("user not found")
	}
	return user.CashBalance, nil
}
