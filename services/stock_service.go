package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/tim/StockFlow/types"
	"github.com/tim/StockFlow/websocket"
)

var (
	stockPrices = map[string]float64{
		"AAPL":  150.00,
		"GOOGL": 2800.00,
		"TSLA":  700.00,
	}
	mu = &sync.RWMutex{}

	// symbolsData stores the loaded stock symbols and company names.
	symbolsData map[string]string
)

// LoadSymbols loads stock symbols and company names from a JSON file.
func LoadSymbols(filePath string) error {
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read symbols file: %w", err)
	}

	if err := json.Unmarshal(fileContent, &symbolsData); err != nil {
		return fmt.Errorf("failed to unmarshal symbols data: %w", err)
	}

	log.Printf("Loaded %d stock symbols.", len(symbolsData))
	return nil
}

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

// SearchStocks searches for stock symbols and company names.
func SearchStocks(query string) []types.StockSearchResult {
	mu.RLock()
	defer mu.RUnlock()

	query = strings.ToLower(query)
	var results []types.StockSearchResult

	for symbol, companyName := range symbolsData {
		if strings.Contains(strings.ToLower(symbol), query) || strings.Contains(strings.ToLower(companyName), query) {
			results = append(results, types.StockSearchResult{
				Symbol:      symbol,
				CompanyName: companyName,
			})
		}
	}

	return results
}
