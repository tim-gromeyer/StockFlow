package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tim/StockFlow/services"
)

// GetPortfolio handles retrieving a user's portfolio.
// @Summary Get user portfolio
// @Description Get a user's portfolio and total value
// @Tags portfolio
// @Produce  json
// @Success 200 {object} PortfolioResponse
// @Failure 400 {object} ErrorResponse "Invalid user ID"
// @Failure 404 {object} ErrorResponse "User not found"
// @Router /api/portfolio [get]
func GetPortfolio(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "User ID not found in context"})
		return
	}

	portfolio, totalValue, err := services.GetPortfolio(userID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: err.Error()})
		return
	}

	cashBalance, err := services.GetUserBalance(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	var portfolioItems []PortfolioItem
	for _, item := range portfolio {
		portfolioItems = append(portfolioItems, PortfolioItem{
			StockSymbol: item.StockSymbol,
			Quantity:    item.Quantity,
		})
	}

	c.JSON(http.StatusOK, PortfolioResponse{
		Portfolio:   portfolioItems,
		TotalValue:  totalValue,
		CashBalance: cashBalance,
	})
}
