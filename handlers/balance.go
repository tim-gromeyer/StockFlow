package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tim/StockFlow/services"
)

// GetBalance handles retrieving a user's cash balance.
// @Summary Get user balance
// @Description Get a user's cash balance
// @Tags portfolio
// @Produce  json
// @Param   user_id  path    int  true  "User ID"
// @Success 200 {object} BalanceResponse
// @Failure 400 {string} string "Invalid user ID"
// @Failure 404 {string} string "User not found"
// @Router /balance/{user_id} [get]
func GetBalance(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	balance, err := services.GetUserBalance(uint(userID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, BalanceResponse{Balance: balance})
}
