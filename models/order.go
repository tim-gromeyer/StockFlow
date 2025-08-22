package models

import (
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// OrderType defines the type of order (e.g., Market, Limit, Stop)
type OrderType int

const (
	OrderTypeMarket OrderType = iota
	OrderTypeLimit
	OrderTypeStop
)

// String returns the string representation of the OrderType.
func (ot OrderType) String() string {
	switch ot {
	case OrderTypeMarket:
		return "MARKET"
	case OrderTypeLimit:
		return "LIMIT"
	case OrderTypeStop:
	    return "STOP"
	default:
		return fmt.Sprintf("UNKNOWN_ORDER_TYPE(%d)", ot)
	}
}

// MarshalJSON implements the json.Marshaler interface.
func (ot OrderType) MarshalJSON() ([]byte, error) {
	return json.Marshal(ot.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (ot *OrderType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "MARKET":
		*ot = OrderTypeMarket
	case "LIMIT":
		*ot = OrderTypeLimit
	case "STOP":
		*ot = OrderTypeStop
	default:
		return fmt.Errorf("invalid OrderType: %s", s)
	}
	return nil
}

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
