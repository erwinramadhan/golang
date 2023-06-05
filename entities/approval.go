package entities

import "time"

type Approval struct {
	ID           uint  `gorm:"primary_key"`
	AdminID      uint  `gorm:"column:admin_id"`
	SuperAdminId *uint `gorm:"column:super_admin_id"`
	Status       *bool `gorm:"column:status"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
