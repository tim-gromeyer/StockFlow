package models

import "time"

// StockPrice represents the price of a stock at a specific time.
type StockPrice struct {
	Symbol    string    `gorm:"primaryKey"`
	Timestamp time.Time `gorm:"primaryKey"`
	Open      float64   `gorm:"not null"`
	High      float64   `gorm:"not null"`
	Low       float64   `gorm:"not null"`
	Close     float64   `gorm:"not null"`
	Volume    int64     `gorm:"not null"`
}
