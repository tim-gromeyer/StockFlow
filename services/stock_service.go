package services

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	stockPrices = map[string]float64{
		"AAPL":  150.00,
		"GOOGL": 2800.00,
		"TSLA":  700.00,
	}
	mu = &sync.RWMutex{}
)

// StartMarketSimulation simulates market movements.
func StartMarketSimulation() {
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			mu.Lock()
			for symbol, price := range stockPrices {
				change := (rand.Float64() - 0.5) * 0.1 // Fluctuation between -5% and +5%
				stockPrices[symbol] = price * (1 + change)
			}
			mu.Unlock()
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
