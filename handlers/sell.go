package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tim/StockFlow/services"
)

// SellRequest represents the request body for selling a stock.
// swagger:parameters sellStock
type SellRequest struct {
	UserID      uint   `json:"user_id" binding:"required"`
	StockSymbol string `json:"stock_symbol" binding:"required"`
	Quantity    int    `json:"quantity" binding:"required"`
}

// SellStock handles selling a stock for a user.
// @Summary Sell a stock
// @Description Sell a specified quantity of a stock for a user
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param   sell_request  body    SellRequest  true  "Sell Request"
// @Success 200 {string} string "Successfully sold stock"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /sell [post]
func SellStock(c *gin.Context) {
	var req SellRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.SellStock(req.UserID, req.StockSymbol, req.Quantity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully sold stock"})
}
