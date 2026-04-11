package entity

import (
	"time"

	"gorm.io/gorm"
)

type Pin struct {
	ID         string         `gorm:"primary_key" json:"id"`
	UserID     string         `gorm:"type:varchar(36);index" json:"user_id"`
	MerchantID string         `gorm:"type:varchar(36);index" json:"merchant_id"`
	Pin        string         `gorm:"pin" json:"pin"`
	CreatedBy  string         `gorm:"created_by" json:"created_by"`
	UpdatedBy  string         `gorm:"update_by" json:"update_by"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
