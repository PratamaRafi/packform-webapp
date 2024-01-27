package models

type (
	Companies struct {
		ID           uint   `gorm:"primary_key" json:"task_id"`
		Company_Name string `gorm:"not null" json:"company_name"`
	}
)
