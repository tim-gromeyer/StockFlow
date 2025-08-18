package services

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/tim/StockFlow/websocket"
)

var (
	stockPrices = map[string]float64{
		"AAPL":  150.00,
		"GOOGL": 2800.00,
		"TSLA":  700.00,
	}
	mu = &sync.RWMutex{}
)

// StartMarketSimulation simulates market movements and broadcasts updates.
func StartMarketSimulation(hub *websocket.Hub) {
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			mu.Lock()
			var updates []websocket.StockPriceUpdate
			for symbol, price := range stockPrices {
				prevPrice := price
				change := (rand.Float64() - 0.5) * 0.1 // Fluctuation between -5% and +5%
				newPrice := price * (1 + change)
				stockPrices[symbol] = newPrice

				updates = append(updates, websocket.StockPriceUpdate{
					Symbol:    symbol,
					Price:     newPrice,
					PrevPrice: prevPrice,
				})
			}
			mu.Unlock()

			jsonUpdate, err := json.Marshal(updates)
			if err != nil {
				log.Printf("Error marshalling stock update: %v", err)
				continue
			}
			hub.Broadcast(jsonUpdate)
		}
	}()
}

// GetStockPrice safely retrieves the current price of a stock.
func GetStockPrice(symbol string) (float64, error) {
	mu.RLock()
	defer mu.RUnlock()

	price, ok := stockPrices[symbol]
	if !ok {
		return 0, fmt.Errorf("stock symbol %s not found", symbol)
	}
	return price, nil
}
