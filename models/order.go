package models

import "time"

type (
	Orders struct {
		ID         uint      `gorm:"primary_key" json:"order_id"`
		CreatedAt  time.Time `json:"created_at"`
		OrderName  string    `gorm:"not null" json:"order_name"`
		CustomerID string    `gorm:"foreign_key" json:"customer_id"`
	}
)
