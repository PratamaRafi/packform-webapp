package models

type (
	Deliveries struct {
		ID                uint `gorm:"primary_key" json:"delivery_id"`
		OrderItemID       uint `gorm:"foreign_key" json:"order_item_id"`
		DeliveredQuantity uint `gorm:"not null" json:"delivered_quantity"`
	}
)
