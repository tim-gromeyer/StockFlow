package services

import (
	"fmt"

	"github.com/tim/StockFlow/database"
	"github.com/tim/StockFlow/models"
	"github.com/tim/StockFlow/types"
)

func GetPortfolio(userID uint) ([]types.PortfolioItem, float64, float64, float64, error) {
	var portfolio []models.Portfolio
	if err := database.DB.Where("user_id = ?", userID).Find(&portfolio).Error; err != nil {
		return nil, 0, 0, 0, fmt.Errorf("user not found")
	}

	var portfolioItems []types.PortfolioItem
	totalValue := 0.0
	totalCost := 0.0
	for _, item := range portfolio {
		currentPrice := GetCurrentPrice(item.StockSymbol)
		averagePrice, err := GetAveragePrice(userID, item.StockSymbol)
		if err != nil {
			return nil, 0, 0, 0, err
		}
		previousDayClose, err := GetPreviousDayClose(item.StockSymbol)
		if err != nil {
			// If there is no previous day close, we can't calculate the percentage change.
			previousDayClose = 0
		}

		var percentageChange float64
		if previousDayClose > 0 {
			percentageChange = ((currentPrice - previousDayClose) / previousDayClose) * 100
		}

		portfolioItems = append(portfolioItems, types.PortfolioItem{
			StockSymbol:      item.StockSymbol,
			Quantity:         item.Quantity,
			AveragePrice:     averagePrice,
			CurrentPrice:     currentPrice,
			PercentageChange: percentageChange,
		})

		totalValue += float64(item.Quantity) * currentPrice
		totalCost += float64(item.Quantity) * averagePrice
	}

	overallGainLoss := totalValue - totalCost
	var overallGainLossPercentage float64
	if totalCost > 0 {
		overallGainLossPercentage = (overallGainLoss / totalCost) * 100
	}

	return portfolioItems, totalValue, overallGainLoss, overallGainLossPercentage, nil
}

func GetAveragePrice(userID uint, symbol string) (float64, error) {
	var orders []models.Order
	if err := database.DB.Where("user_id = ? AND stock_symbol = ? AND is_buy = ?", userID, symbol, true).Find(&orders).Error; err != nil {
		return 0, err
	}

	totalCost := 0.0
	totalQuantity := 0
	for _, order := range orders {
		totalCost += order.ExecutedPrice * float64(order.Quantity)
		totalQuantity += order.Quantity
	}

	if totalQuantity == 0 {
		return 0, nil
	}

	return totalCost / float64(totalQuantity), nil
}

func GetUserBalance(userID uint) (float64, error) {
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return 0, fmt.Errorf("user not found")
	}
	return user.CashBalance, nil
}
