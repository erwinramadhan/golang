package entities

import "time"

type boolType string

const (
	TRUE  boolType = "true"
	FALSE boolType = "false"
)

type Account struct {
	ID        uint   `gorm:"primary_key"`
	Username  string `gorm:"column:username"`
	Password  string `gorm:"column:password"`
	RoleID    uint   `gorm:"column:role_id"`
	Verified  bool   `gorm:"column:verified;type:boolean"`
	Active    bool   `gorm:"column:active;type:boolean"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
