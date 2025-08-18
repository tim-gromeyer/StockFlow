package handlers

import "github.com/tim/StockFlow/models"

// PortfolioResponse represents the response for the get portfolio endpoint.
// swagger:model
type PortfolioResponse struct {
	Portfolio  []models.Portfolio `json:"portfolio"`
	TotalValue float64            `json:"total_value"`
}

// BalanceResponse represents the response for the get balance endpoint.
// swagger:model
type BalanceResponse struct {
	Balance float64 `json:"balance"`
}
