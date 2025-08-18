package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tim/StockFlow/database"
)

// HealthCheck handles the health check endpoint.
// @Summary Health check
// @Description Checks the health of the API and its dependencies (e.g., database).
// @Tags health
// @Produce  json
// @Success 200 {object} SuccessResponse "API is healthy"
// @Failure 500 {object} ErrorResponse "API is unhealthy"
// @Router /health [get]
func HealthCheck(c *gin.Context) {
	// Check database connection
	sqlDB, err := database.DB.DB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Database connection error"})
		return
	}
	if err = sqlDB.Ping(); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Database not reachable"})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{Message: "API is healthy"})
}
