package models

import (
	"time"

	"gorm.io/gorm"
)

// OrderType defines the type of order (e.g., Market, Limit, Stop)
type OrderType string

const (
	OrderTypeMarket OrderType = "MARKET"
	OrderTypeLimit  OrderType = "LIMIT"
	OrderTypeStop   OrderType = "STOP"
)

// OrderStatus defines the status of an order
type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "PENDING"
	OrderStatusExecuted  OrderStatus = "EXECUTED"
	OrderStatusCancelled OrderStatus = "CANCELLED"
)

// Order represents a pending or executed stock order
type Order struct {
	gorm.Model
	UserID      uint        `gorm:"not null"`
	StockSymbol string      `gorm:"not null"`
	Quantity    int         `gorm:"not null"`
	OrderType   OrderType   `gorm:"not null"`
	LimitPrice  float64     // Only for LIMIT orders
	StopPrice   float64     // Only for STOP orders
	IsBuy       bool        `gorm:"not null"` // true for buy, false for sell
	Status      OrderStatus `gorm:"not null"`
	ExecutedPrice float64     // Price at which the order was executed
	ExecutedAt  *time.Time  // Timestamp of execution
}
