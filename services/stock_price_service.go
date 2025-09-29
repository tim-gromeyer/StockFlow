package services

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/tim/StockFlow/database"
	"github.com/tim/StockFlow/models"
)

// FetchAndStoreStockPrices fetches the latest 1-minute interval stock data from Alpha Vantage
// and stores it in the database.
// FetchAndStoreStockPrices fetches the latest 1-minute interval stock data from Alpha Vantage
// and stores it in the database.
func FetchAndStoreStockPrices(symbol string) error {
	if apiKey == "YOUR_API_KEY" {
		return fmt.Errorf("please replace 'YOUR_API_KEY' with your actual Alpha Vantage API key")
	}

	url := fmt.Sprintf(alphaVantageURL, symbol, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch data from Alpha Vantage: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("alpha Vantage API request failed with status code: %d", resp.StatusCode)
	}

	reader := csv.NewReader(resp.Body)
	// Skip header row
	if _, err := reader.Read(); err != nil {
		return fmt.Errorf("failed to read header row from CSV: %w", err)
	}

	var stockPrices []models.StockPrice
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read record from CSV: %w", err)
		}

		timestamp, err := time.Parse("2006-01-02 15:04:05", record[0])
		if err != nil {
			return fmt.Errorf("failed to parse timestamp: %w", err)
		}

		open, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			return fmt.Errorf("failed to parse open price: %w", err)
		}

		high, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			return fmt.Errorf("failed to parse high price: %w", err)
		}

		low, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			return fmt.Errorf("failed to parse low price: %w", err)
		}

		close, err := strconv.ParseFloat(record[4], 64)
		if err != nil {
			return fmt.Errorf("failed to parse close price: %w", err)
		}

		volume, err := strconv.ParseInt(record[5], 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse volume: %w", err)
		}

		stockPrices = append(stockPrices, models.StockPrice{
			Symbol:    symbol,
			Timestamp: timestamp,
			Open:      open,
			High:      high,
			Low:       low,
			Close:     close,
			Volume:    volume,
		})
	}

	if len(stockPrices) > 0 {
		result := database.DB.Create(&stockPrices)
		if result.Error != nil {
			return fmt.Errorf("failed to save stock prices to database: %w", result.Error)
		}
	}

	return nil
}
