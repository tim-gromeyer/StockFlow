package services

import (
	"errors"
	"log"

	"github.com/tim/StockFlow/database"
	"github.com/tim/StockFlow/models"
	"gorm.io/gorm"
)

// BuyStock handles the purchase of stock.
func BuyStock(userID uint, stockSymbol string, quantity int, orderType models.OrderType, limitPrice, stopPrice float64) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		var user models.User
		if err := tx.First(&user, userID).Error; err != nil {
			return errors.New("user not found")
		}

		// Get current stock price
		currentPrice, err := GetStockPrice(stockSymbol)
		if err != nil {
			return err
		}

		switch orderType {
		case models.OrderTypeMarket:
			// Market order: execute immediately
			totalCost := currentPrice * float64(quantity)
			if user.CashBalance < totalCost {
				return errors.New("insufficient cash balance")
			}

			user.CashBalance -= totalCost
			if err := tx.Save(&user).Error; err != nil {
				return err
			}

			// Update portfolio
			var portfolio models.Portfolio
			if err := tx.Where("user_id = ? AND stock_symbol = ?", userID, stockSymbol).FirstOrCreate(&portfolio, models.Portfolio{UserID: userID, StockSymbol: stockSymbol}).Error; err != nil {
				return err
			}
			portfolio.Quantity += quantity
			if err := tx.Save(&portfolio).Error; err != nil {
				return err
			}
			log.Printf("User %d bought %d of %s at market price %.2f. New balance: %.2f", userID, quantity, stockSymbol, currentPrice, user.CashBalance)

		case models.OrderTypeLimit, models.OrderTypeStop:
			// Limit or Stop order: create pending order
			order := models.Order{
				UserID:      userID,
				StockSymbol: stockSymbol,
				Quantity:    quantity,
				OrderType:   orderType,
				IsBuy:       true,
				Status:      models.OrderStatusPending,
			}
			if orderType == models.OrderTypeLimit {
				order.LimitPrice = limitPrice
			} else if orderType == models.OrderTypeStop {
				order.StopPrice = stopPrice
			}

			if err := tx.Create(&order).Error; err != nil {
				return err
			}
			log.Printf("User %d placed a %s buy order for %d of %s. Order ID: %d", userID, orderType, quantity, stockSymbol, order.ID)

		default:
			return errors.New("invalid order type")
		}

		return nil
	})
}

// SellStock handles the selling of stock.
func SellStock(userID uint, stockSymbol string, quantity int, orderType models.OrderType, limitPrice, stopPrice float64) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		var user models.User
		if err := tx.First(&user, userID).Error; err != nil {
			return errors.New("user not found")
		}

		// Get current stock price
		currentPrice, err := GetStockPrice(stockSymbol)
		if err != nil {
			return err
		}

		switch orderType {
		case models.OrderTypeMarket:
			// Market order: execute immediately
			var portfolio models.Portfolio
			if err := tx.Where("user_id = ? AND stock_symbol = ?", userID, stockSymbol).First(&portfolio).Error; err != nil {
				return errors.New("stock not found in portfolio")
			}

			if portfolio.Quantity < quantity {
				return errors.New("insufficient stock quantity")
			}

			portfolio.Quantity -= quantity
			if err := tx.Save(&portfolio).Error; err != nil {
				return err
			}

			user.CashBalance += currentPrice * float64(quantity)
			if err := tx.Save(&user).Error; err != nil {
				return err
			}
			log.Printf("User %d sold %d of %s at market price %.2f. New balance: %.2f", userID, quantity, stockSymbol, currentPrice, user.CashBalance)

		case models.OrderTypeLimit, models.OrderTypeStop:
			// Limit or Stop order: create pending order
			order := models.Order{
				UserID:      userID,
				StockSymbol: stockSymbol,
				Quantity:    quantity,
				OrderType:   orderType,
				IsBuy:       false,
				Status:      models.OrderStatusPending,
			}
			if orderType == models.OrderTypeLimit {
				order.LimitPrice = limitPrice
			} else if orderType == models.OrderTypeStop {
				order.StopPrice = stopPrice
			}

			if err := tx.Create(&order).Error; err != nil {
				return err
			}
			log.Printf("User %d placed a %s sell order for %d of %s. Order ID: %d", userID, orderType, quantity, stockSymbol, order.ID)

		default:
			return errors.New("invalid order type")
		}

		return nil
	})
}
