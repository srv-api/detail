package entity

import (
	"time"

	"gorm.io/gorm"
)

type UploadedPayment struct {
	ID         uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID     string         `gorm:"type:varchar(36);index" json:"user_id"`
	MerchantID string         `gorm:"type:varchar(36);index" json:"merchant_id"`
	PaymentID  uint           `gorm:"payment_id,omitempty" json:"payment_id"`
	FileName   string         `gorm:"size:255;not null" json:"file_name"`
	FilePath   string         `gorm:"file_path,omitempty" json:"file_path"`
	CreatedBy  string         `gorm:"created_by" json:"created_by"`
	UpdatedBy  string         `gorm:"updated_by" json:"updated_by"`
	DeletedBy  string         `gorm:"deleted_by" json:"deleted_by"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
