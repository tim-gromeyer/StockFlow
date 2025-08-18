package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tim/StockFlow/services"
)

// GetBalance handles retrieving a user's cash balance.
// @Summary Get user balance
// @Description Get a user's cash balance
// @Tags portfolio
// @Produce  json
// @Success 200 {object} BalanceResponse
// @Failure 400 {object} ErrorResponse "Invalid user ID"
// @Failure 404 {object} ErrorResponse "User not found"
// @Router /api/balance [get]
func GetBalance(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "User ID not found in context"})
		return
	}

	balance, err := services.GetUserBalance(userID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, BalanceResponse{Balance: balance})
}
