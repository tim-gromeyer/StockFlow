package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tim/StockFlow/services"
)

// GetPortfolio handles retrieving a user's portfolio.
// @Summary Get user portfolio
// @Description Get a user's portfolio and total value
// @Tags portfolio
// @Produce  json
// @Param   user_id  path    int  true  "User ID"
// @Success 200 {object} PortfolioResponse
// @Failure 400 {string} string "Invalid user ID"
// @Failure 404 {string} string "User not found"
// @Router /portfolio/{user_id} [get]
func GetPortfolio(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	portfolio, totalValue, err := services.GetPortfolio(uint(userID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, PortfolioResponse{
		Portfolio:  portfolio,
		TotalValue: totalValue,
	})
}
