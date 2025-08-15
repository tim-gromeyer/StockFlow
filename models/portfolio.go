package models

// Portfolio represents a user's holding of a specific stock.
// swagger:model
type Portfolio struct {
	UserID      uint   `gorm:"primaryKey"`
	StockSymbol string `gorm:"primaryKey"`
	Quantity    int    `gorm:"not null"`
}
