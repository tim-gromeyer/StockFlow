package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tim/StockFlow/services"
)

// FXRateResponse represents the response for the FX rate endpoint.
// swagger:response fxRateResponse
type FXRateResponse struct {
	FromCurrency   string  `json:"from_currency"`
	ToCurrency     string  `json:"to_currency"`
	Rate           float64 `json:"rate"`
	DailyChange    float64 `json:"daily_change"`
	PercentChange  float64 `json:"percent_change"`
}

// GetFXRate handles the request to get the exchange rate and daily change for a currency pair.
// @Summary Get FX rate
// @Description Get the exchange rate and daily change for a currency pair.
// @Tags fx
// @Produce  json
// @Param   from    query   string  true        "From Currency"
// @Param   to      query   string  true        "To Currency"
// @Success 200 {object} FXRateResponse
// @Failure 400 {object} ErrorResponse "Invalid currency codes"
// @Failure 500 {object} ErrorResponse "Failed to fetch FX data"
// @Router /api/fx [get]
func GetFXRate(c *gin.Context) {
	from := c.Query("from")
	to := c.Query("to")

	if from == "" || to == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "'from' and 'to' query parameters are required"})
		return
	}

	rate, err := services.GetExchangeRate(from, to)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	dailyChange, percentChange, err := services.GetDailyChange(from, to)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, FXRateResponse{
		FromCurrency:  from,
		ToCurrency:    to,
		Rate:          rate,
		DailyChange:   dailyChange,
		PercentChange: percentChange,
	})
}
