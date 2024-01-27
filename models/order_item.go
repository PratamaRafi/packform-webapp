package models

type (
	OrderItem struct {
		ID       uint    `gorm:"primary_key" json:"order_item_id"`
		OrderID  uint    `gorm:"foreign_key" json:"order_id"`
		Price    float32 `gorm:"default:null" json:"price_per_unit"`
		Quantity uint    `gorm:"not null" json:"quantity"`
		Prouct   string  `gorm:"not null" json:"product"`
	}
)
