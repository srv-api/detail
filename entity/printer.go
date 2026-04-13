package entity

import (
	"time"

	"gorm.io/gorm"
)

type Printer struct {
	ID          string         `gorm:"id" json:"id"`
	UserID      string         `gorm:"type:varchar(36);index" json:"user_id"`
	MerchantID  string         `gorm:"type:varchar(36);index" json:"merchant_id"`
	PrinterName string         `gorm:"type:varchar(50)" json:"printer_name"`
	Description string         `gorm:"description" json:"description"`
	CreatedBy   string         `gorm:"created_by" json:"created_by"`
	UpdatedBy   string         `gorm:"update_by" json:"update_by"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type PrinterAutoIncrement struct {
	MerchantID    string `gorm:"primary_key"`
	NextIncrement int    `gorm:"not null;default:1"`
}
