package entity

import (
	"time"

	"gorm.io/gorm"
)

type Merk struct {
	ID          string         `gorm:"primary_key,omitempty" json:"id"`
	UserID      string         `gorm:"type:varchar(36);index" json:"user_id"`
	MerchantID  string         `gorm:"type:varchar(36);index;not null" json:"merchant_id"`
	MerkName    string         `gorm:"merk_name,omitempty" json:"merk_name"`
	Status      int            `gorm:"status" json:"status"`
	Description string         `gorm:"description" json:"description"`
	CreatedBy   string         `gorm:"created_by" json:"created_by"`
	UpdatedBy   string         `gorm:"updated_by" json:"updated_by"`
	DeletedBy   string         `gorm:"deleted_by" json:"deleted_by"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
