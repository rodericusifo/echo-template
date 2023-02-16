package sql

import (
	"time"

	"gorm.io/gorm"

	"github.com/google/uuid"
	"github.com/rodericusifo/echo-template/internal/pkg/constant"
)

type User struct {
	// Migrated Fields
	ID        uint   `gorm:"primaryKey"`
	XID       string `gorm:"column:xid"`
	Name      string
	Email     string
	Password  string
	Role      constant.UserRole
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.XID == "" {
		u.XID = uuid.NewString()
	}
	return nil
}

func (User) TableName() string {
	return "users"
}
