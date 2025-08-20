package types

// StockSearchResult represents a single search result for a stock.
type StockSearchResult struct {
	Symbol      string `json:"symbol"`
	CompanyName string `json:"companyName"`
}
