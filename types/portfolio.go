package types

// PortfolioItem represents a single item in a user's portfolio.
type PortfolioItem struct {
	StockSymbol      string  `json:"stockSymbol"`
	Quantity         int     `json:"quantity"`
	AveragePrice     float64 `json:"avg_price"`
	CurrentPrice     float64 `json:"curr_price"`
	PercentageChange float64 `json:"percentage_change"`
}
