package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/tim/StockFlow/database"
	"github.com/tim/StockFlow/models"
)

// GetExchangeRate fetches the current exchange rate for a given currency pair.
func GetExchangeRate(from, to string) (float64, error) {
	url := fmt.Sprintf(alphaVantageFXURL, from, to, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch data from Alpha Vantage: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("alpha Vantage API request failed with status code: %d", resp.StatusCode)
	}

	var result struct {
		RealtimeCurrencyExchangeRate struct {
			ExchangeRate string `json:"5. Exchange Rate"`
		} `json:"Realtime Currency Exchange Rate"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, fmt.Errorf("failed to decode response from Alpha Vantage: %w", err)
	}

	rate, err := strconv.ParseFloat(result.RealtimeCurrencyExchangeRate.ExchangeRate, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse exchange rate: %w", err)
	}

	fxRate := models.FXRate{
		FromCurrency: from,
		ToCurrency:   to,
		Timestamp:    time.Now(),
		Rate:         rate,
	}

	if err := database.DB.Create(&fxRate).Error; err != nil {
		return 0, fmt.Errorf("failed to save exchange rate to database: %w", err)
	}

	return rate, nil
}

// GetDailyChange fetches the daily change in price and percentage for a given currency pair.
func GetDailyChange(from, to string) (float64, float64, error) {
	url := fmt.Sprintf(alphaVantageFXDailyURL, from, to, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to fetch data from Alpha Vantage: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, 0, fmt.Errorf("alpha Vantage API request failed with status code: %d", resp.StatusCode)
	}

	var result struct {
		TimeSeriesFXDaily map[string]struct {
			Close string `json:"4. close"`
		} `json:"Time Series FX (Daily)"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, 0, fmt.Errorf("failed to decode response from Alpha Vantage: %w", err)
	}

	if len(result.TimeSeriesFXDaily) < 2 {
		return 0, 0, fmt.Errorf("not enough data to calculate daily change")
	}

	// Get the dates for the last two days
	var dates []string
	for date := range result.TimeSeriesFXDaily {
		dates = append(dates, date)
	}
	// Sort dates in descending order
	for i := 0; i < len(dates); i++ {
		for j := i + 1; j < len(dates); j++ {
			if dates[i] < dates[j] {
				dates[i], dates[j] = dates[j], dates[i]
			}
		}
	}

	today, err := strconv.ParseFloat(result.TimeSeriesFXDaily[dates[0]].Close, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse today's close price: %w", err)
	}

	yesterday, err := strconv.ParseFloat(result.TimeSeriesFXDaily[dates[1]].Close, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse yesterday's close price: %w", err)
	}

	change := today - yesterday
	percentChange := (change / yesterday) * 100

	return change, percentChange, nil
}
