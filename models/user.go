package models

// User represents a user in the system
// with a unique ID and a cash balance.
// swagger:model
type User struct {
	ID          uint    `gorm:"primaryKey"`
	CashBalance float64 `gorm:"not null"`
}
