package services

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/tim/StockFlow/types"
)

var (
	// symbolsData stores the loaded stock symbols and company names.
	symbolsData map[string]string
)

// LoadSymbols loads stock symbols and company names from the embedded JSON file.
func LoadSymbols(content embed.FS) error {
	fileContent, err := content.ReadFile("assets/symbols.json")
	if err != nil {
		return fmt.Errorf("failed to read embedded symbols file: %w", err)
	}

	if err := json.Unmarshal(fileContent, &symbolsData); err != nil {
		return fmt.Errorf("failed to unmarshal symbols data: %w", err)
	}

	log.Printf("Loaded %d stock symbols.", len(symbolsData))
	return nil
}

// SearchStocks searches for stock symbols and company names.
func SearchStocks(query string) []types.StockSearchResult {
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
