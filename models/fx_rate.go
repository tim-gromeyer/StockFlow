package models

import "time"

// FXRate represents the exchange rate between two currencies at a specific time.
type FXRate struct {
	FromCurrency string    `gorm:"primaryKey"`
	ToCurrency   string    `gorm:"primaryKey"`
	Timestamp    time.Time `gorm:"primaryKey"`
	Rate         float64   `gorm:"not null"`
}
