package services

import (
	"time"

	"github.com/tim/StockFlow/database"
	"github.com/tim/StockFlow/models"
)

// GetPreviousDayClose returns the close price of the previous day for a given stock.
func GetPreviousDayClose(symbol string) (float64, error) {
	var stockPrice models.StockPrice
	yesterday := time.Now().AddDate(0, 0, -1)
	if err := database.DB.Where("symbol = ? AND DATE(timestamp) = ?", symbol, yesterday.Format("2006-01-02")).Order("timestamp desc").First(&stockPrice).Error; err != nil {
		return 0, err
	}
	return stockPrice.Close, nil
}
