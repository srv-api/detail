package entity

import (
	"time"

	"gorm.io/gorm"
)

type Permission struct {
	ID         uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID     string         `gorm:"type:varchar(36);index" json:"user_id"`
	MerchantID string         `gorm:"type:varchar(36);index" json:"merchant_id"`
	Label      string         `gorm:"type:varchar(100)" json:"label"`
	Icon       string         `gorm:"type:varchar(255)" json:"icon"`
	To         string         `gorm:"type:varchar(255)" json:"to"`
	CreatedBy  string         `json:"created_by"`
	UpdatedBy  string         `json:"updated_by"`
	DeletedBy  string         `json:"deleted_by"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
