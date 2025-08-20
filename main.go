package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tim/StockFlow/database"
	"github.com/tim/StockFlow/handlers"
	"github.com/tim/StockFlow/middleware"
	"github.com/tim/StockFlow/services"
	"github.com/tim/StockFlow/websocket"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/tim/StockFlow/docs"
)

// @title StockFlow API
// @version 1.0
// @description This is the API for the StockFlow application.
// @host localhost:8080
// @BasePath /
func main() {
	// Initialize database
	database.InitDatabase()

	// Load stock symbols
	if err := services.LoadSymbols("./assets/symbols.json"); err != nil {
		log.Fatalf("Failed to load stock symbols: %v", err)
	}

	r := gin.Default()

	// CORS Middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	// WebSocket Hub
	hub := websocket.NewHub()
	go hub.Run()

	// Start market simulation
	go services.StartMarketSimulation(hub)

	// Routes
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Auth routes
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/register", handlers.Register)
		authRoutes.POST("/login", handlers.Login)
	}

	// API routes
	apiRoutes := r.Group("/api")
	apiRoutes.Use(middleware.AuthMiddleware())
	{
		apiRoutes.GET("/health", handlers.HealthCheck)
		apiRoutes.GET("/portfolio", handlers.GetPortfolio)
		apiRoutes.GET("/balance", handlers.GetBalance)
		apiRoutes.POST("/buy", handlers.BuyStock)
		apiRoutes.POST("/sell", handlers.SellStock)
		apiRoutes.GET("/stocks/search", handlers.SearchStocks) // New search endpoint
	}

	// WebSocket route
	r.GET("/ws", func(c *gin.Context) {
		websocket.ServeWs(hub, c.Writer, c.Request)
	})

	r.Run(":8080")
}
