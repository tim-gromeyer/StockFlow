package services

import (
	"fmt"
	"time"

	"github.com/tim/StockFlow/database"
	"github.com/tim/StockFlow/models"
	"github.com/tim/StockFlow/websocket"
)

var (
	// a map to store the current price of each stock
	currentPrices = make(map[string]float64)
)

// GetCurrentPrice returns the current price of a stock.
func GetCurrentPrice(symbol string) float64 {
	return currentPrices[symbol]
}

// StartMarketSimulation simulates market movements by iterating through historical data.
func StartMarketSimulation(hub *websocket.Hub) {
	for {
		// Get all distinct stock symbols from the database
		var symbols []string
		database.DB.Model(&models.StockPrice{}).Distinct().Pluck("symbol", &symbols)

		for _, symbol := range symbols {
			go simulateStock(symbol, hub)
		}

		// Re-fetch symbols every hour
		time.Sleep(1 * time.Hour)
	}
}

func simulateStock(symbol string, hub *websocket.Hub) {
	var prices []models.StockPrice
	database.DB.Where("symbol = ?", symbol).Order("timestamp asc").Find(&prices)

	if len(prices) == 0 {
		return
	}

	for _, price := range prices {
		currentPrices[symbol] = price.Close

		// Broadcast the new price to all clients
		hub.Broadcast([]byte(fmt.Sprintf("{\"symbol\": \"%s\", \"price\": %f}", symbol, price.Close)))

		// Simulate real-time market movement
		time.Sleep(1 * time.Second)
	}
}
