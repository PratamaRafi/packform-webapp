package models

import (
	"github.com/lib/pq"
)

type (
	Customer struct {
		UserID     string         `gorm:"primary_key" json:"user_id"`
		Login      string         `gorm:"not null" json:"login"`
		Password   string         `gorm:"not null" json:"password"`
		Name       string         `gorm:"not null" json:"name"`
		CompanyID  uint           `gorm:"foreign_key" json:"company_id"`
		CreditCard pq.StringArray `gorm:"type:text[]" json:"credit_card"`
	}
)
