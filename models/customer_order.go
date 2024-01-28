package models

import "time"

type (
	CustomerOrder struct {
		OrderName         string    `gorm:"not null" json:"order_name"`
		Company_Name      string    `gorm:"not null" json:"company_name"`
		Name              string    `gorm:"not null" json:"name"`
		CreatedAt         time.Time `json:"created_at"`
		DeliveredAmount   float32   `gorm:"not null" json:"delivered_amount"`
		TotalAmount       float32   `gorm:"not null" json:"total_amount"`
		Price             float32   `gorm:"not null" json:"price"`
		DeliveredQuantity uint      `gorm:"not null" json:"delivered_quantity"`
	}
)
