package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tim/StockFlow/models"
	"github.com/tim/StockFlow/services"
)

// SellRequest represents the request body for selling a stock.
// swagger:parameters sellStock
type SellRequest struct {
	StockSymbol string  `json:"stockSymbol" binding:"required"`
	Quantity    int     `json:"quantity" binding:"required,min=1"`
	OrderType   models.OrderType  `json:"orderType" binding:"required"`
	LimitPrice  float64 `json:"limitPrice"` // Optional, for LIMIT orders
	StopPrice   float64 `json:"stopPrice"`  // Optional, for STOP orders
}

// SellStock handles selling a stock for a user.
// @Summary Sell a stock
// @Description Sell a specified quantity of a stock for a user
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param   sell_request  body    SellRequest  true  "Sell Request"
// @Success 200 {object} SuccessResponse "Successfully sold stock"
// @Failure 400 {object} ErrorResponse "Invalid request"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/sell [post]
func SellStock(c *gin.Context) {
	var req SellRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "User ID not found in context"})
		return
	}

	if err := services.SellStock(userID.(uint), req.StockSymbol, req.Quantity, models.OrderType(req.OrderType), req.LimitPrice, req.StopPrice); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{Message: "Successfully sold stock"})
}
