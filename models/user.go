package models

// User represents a user in the system
type User struct {
	ID           uint    `gorm:"primaryKey"`
	Username     string  `gorm:"unique;not null"`
	PasswordHash string  `gorm:"not null"`
	CashBalance  float64 `gorm:"not null"`
}
