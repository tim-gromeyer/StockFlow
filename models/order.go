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
	// OrderTypeMarket represents a market order
	// @name MARKET
	// @enum 1
	OrderTypeMarket OrderType = iota + 1
	// OrderTypeLimit represents a limit order
	// @name LIMIT
	// @enum 2
	OrderTypeLimit
	// OrderTypeStop represents a stop order
	// @name STOP
	// @enum 3
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
	// Try to unmarshal as a string first
	var s string
	if err := json.Unmarshal(b, &s); err == nil {
		switch s {
		case "MARKET":
			*ot = OrderTypeMarket
		case "LIMIT":
			*ot = OrderTypeLimit
		case "STOP":
			*ot = OrderTypeStop
		default:
			return fmt.Errorf("invalid OrderType string: %s", s)
		}
		return nil
	}

	// If string unmarshalling fails, try to unmarshal as an integer
	var i int
	if err := json.Unmarshal(b, &i); err == nil {
		switch i {
		case int(OrderTypeMarket):
			*ot = OrderTypeMarket
		case int(OrderTypeLimit):
			*ot = OrderTypeLimit
		case int(OrderTypeStop):
			*ot = OrderTypeStop
		default:
			return fmt.Errorf("invalid OrderType integer: %d", i)
		}
		return nil
	}

	return fmt.Errorf("OrderType must be a string or an integer")
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
