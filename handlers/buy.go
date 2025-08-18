package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tim/StockFlow/services"
)

// BuyRequest represents the request body for buying a stock.
// swagger:parameters buyStock
type BuyRequest struct {
	StockSymbol string `json:"stock_symbol" binding:"required"`
	Quantity    int    `json:"quantity" binding:"required"`
}

// BuyStock handles buying a stock for a user.
// @Summary Buy a stock
// @Description Buy a specified quantity of a stock for a user
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param   buy_request  body    BuyRequest  true  "Buy Request"
// @Success 200 {object} SuccessResponse "Successfully bought stock"
// @Failure 400 {object} ErrorResponse "Invalid request"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/buy [post]
func BuyStock(c *gin.Context) {
	var req BuyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "User ID not found in context"})
		return
	}

	if err := services.BuyStock(userID.(uint), req.StockSymbol, req.Quantity); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{Message: "Successfully bought stock"})
}
