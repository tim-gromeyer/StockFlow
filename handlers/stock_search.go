package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tim/StockFlow/services"
	"github.com/tim/StockFlow/types"
)

// SearchStocks handles searching for stock symbols and company names.
// @Summary Search stocks
// @Description Search for stock symbols or company names based on a query string
// @Tags stocks
// @Produce  json
// @Param   q  query      string  true  "Search query for stock symbol or company name"
// @Success 200 {array} types.StockSearchResult
// @Failure 400 {object} ErrorResponse "Missing search query"
// @Router /api/stocks/search [get]
func SearchStocks(c *gin.Context) {
	// Explicitly use types.StockSearchResult to avoid "imported and not used" error
	var _ types.StockSearchResult

	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Missing search query"})
		return
	}

	results := services.SearchStocks(query)
	c.JSON(http.StatusOK, results)
}
