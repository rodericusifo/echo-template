package sql

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Employee struct {
	// Migrated Fields
	ID        uint   `gorm:"primaryKey"`
	XID       string `gorm:"column:xid"`
	Name      string
	Email     string
	Address   *string
	Age       *int
	Birthday  *time.Time
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Relations
	User User
}

func (e *Employee) BeforeCreate(tx *gorm.DB) error {
	if e.XID == "" {
		e.XID = uuid.NewString()
	}
	return nil
}

func (Employee) TableName() string {
	return "employees"
}
