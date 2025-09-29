package types

// StockSearchResult represents a single result of a stock search.
type StockSearchResult struct {
	Symbol      string `json:"symbol"`
	CompanyName string `json:"company_name"`
}