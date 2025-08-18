package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tim/StockFlow/services"
)

// BuyRequest represents the request body for buying a stock.
// swagger:parameters buyStock
type BuyRequest struct {
	UserID      uint   `json:"user_id" binding:"required"`
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
// @Router /buy [post]
func BuyStock(c *gin.Context) {
	var req BuyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	if err := services.BuyStock(req.UserID, req.StockSymbol, req.Quantity); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{Message: "Successfully bought stock"})
}
