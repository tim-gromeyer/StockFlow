package services

import (
	"fmt"

	"github.com/tim/StockFlow/database"
	"github.com/tim/StockFlow/models"
	"gorm.io/gorm"
)

func BuyStock(userID uint, symbol string, quantity int) error {
	price, err := GetStockPrice(symbol)
	if err != nil {
		return err
	}

	totalCost := float64(quantity) * price

	return database.DB.Transaction(func(tx *gorm.DB) error {
		var user models.User
		if err := tx.First(&user, userID).Error; err != nil {
			return fmt.Errorf("user not found")
		}

		if user.CashBalance < totalCost {
			return fmt.Errorf("insufficient funds")
		}

		user.CashBalance -= totalCost
		if err := tx.Save(&user).Error; err != nil {
			return err
		}

		var portfolio models.Portfolio
		if err := tx.First(&portfolio, "user_id = ? AND stock_symbol = ?", userID, symbol).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				portfolio = models.Portfolio{
					UserID:      userID,
					StockSymbol: symbol,
					Quantity:    quantity,
				}
				return tx.Create(&portfolio).Error
			}
			return err
		}

		portfolio.Quantity += quantity
		return tx.Save(&portfolio).Error
	})
}

func SellStock(userID uint, symbol string, quantity int) error {
	price, err := GetStockPrice(symbol)
	if err != nil {
		return err
	}

	totalValue := float64(quantity) * price

	return database.DB.Transaction(func(tx *gorm.DB) error {
		var portfolio models.Portfolio
		if err := tx.First(&portfolio, "user_id = ? AND stock_symbol = ?", userID, symbol).Error; err != nil {
			return fmt.Errorf("stock not found in portfolio")
		}

		if portfolio.Quantity < quantity {
			return fmt.Errorf("insufficient stock quantity")
		}

		portfolio.Quantity -= quantity
		if err := tx.Save(&portfolio).Error; err != nil {
			return err
		}

		var user models.User
		if err := tx.First(&user, userID).Error; err != nil {
			return err
		}

		user.CashBalance += totalValue
		return tx.Save(&user).Error
	})
}
