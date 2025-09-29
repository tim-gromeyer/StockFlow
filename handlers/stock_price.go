package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tim/StockFlow/services"
)

// FetchStockPrices handles the request to fetch and store stock prices.
// @Summary Fetch stock prices
// @Description Fetches and stores the latest 1-minute interval stock data for a given symbol.
// @Tags stocks
// @Produce  json
// @Param   symbol    path    string  true        "Stock Symbol"
// @Success 200 {object} SuccessResponse "Successfully fetched and stored stock prices"
// @Failure 400 {object} ErrorResponse "Invalid stock symbol"
// @Failure 500 {object} ErrorResponse "Failed to fetch or store stock prices"
// @Router /api/stocks/{symbol}/fetch [post]
func FetchStockPrices(c *gin.Context) {
	symbol := c.Param("symbol")
	if symbol == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Stock symbol is required"})
		return
	}

	err := services.FetchAndStoreStockPrices(symbol)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{Message: "Successfully fetched and stored stock prices"})
}
