package entity

import (
	"time"

	"gorm.io/gorm"
)

type TopUp struct {
	ID              string         `gorm:"primary_key,omitempty" json:"id"`
	OrderID         string         `gorm:"order_id,omitempty" json:"order_id"`
	TransactionID   string         `gorm:"transaction_id,omitempty" json:"transaction_id"`
	TransactionTime time.Time      `gorm:"transaction_time,omitempty" json:"transaction_time"`
	PaymentType     string         `gorm:"payment_type,omitempty" json:"payment_type"`
	Url             string         `gorm:"url,omitempty" json:"url"`
	GrossAmount     int64          `gorm:"gross_amount" json:"gross_amount"`
	Status          string         `gorm:"status" json:"status"`
	VA              string         `gorm:"va" json:"va"`
	ExpiryTime      time.Time      `gorm:"expiry_time" json:"expiry_time"`
	Bank            string         `gorm:"bank" json:"bank"`
	RedirectURL     string         `gorm:"redirect_url" json:"redirect_url"`
	CreatedBy       string         `gorm:"created_by" json:"created_by"`
	DeletedBy       string         `gorm:"deleted_by" json:"deleted_by"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
