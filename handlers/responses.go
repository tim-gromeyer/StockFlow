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

// RegisterResponse represents the response for the register endpoint.
// swagger:model
type RegisterResponse struct {
	ID          uint    `json:"id"`
	Username    string  `json:"username"`
	CashBalance float64 `json:"cash_balance"`
}

// LoginResponse represents the response for the login endpoint.
// swagger:model
type LoginResponse struct {
	Token string `json:"token"`
}

// ErrorResponse represents the response for an error.
// swagger:model
type ErrorResponse struct {
	Error string `json:"error"`
}

// SuccessResponse represents a generic success response.
// swagger:model
type SuccessResponse struct {
	Message string `json:"message"`
}
