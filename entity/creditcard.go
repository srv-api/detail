package entity

import (
	"time"

	"gorm.io/gorm"
)

type CreditCard struct {
	ID            string         `gorm:"primary_key,omitempty" json:"id"`
	OrderID       string         `gorm:"order_id" json:"order_id"`
	UserID        string         `gorm:"type:varchar(36);index" json:"user_id"`
	CardNumber    string         `gorm:"card_number" json:"card_number"`
	ExpiryMonth   string         `gorm:"expiry_month" json:"expiry_month"`
	ExpiryYear    string         `gorm:"expiry_year" json:"expiry_year"`
	CVV           string         `gorm:"cvv" json:"cvv"`
	GrossAmount   int            `gorm:"gross_amout" json:"gross_amout"`
	TokenID       string         `gorm:"token_id" json:"token_id"`
	TransactionID string         `gorm:"transaction_id" json:"transaction_id"`
	Status        string         `gorm:"status" json:"status"`
	CreatedBy     string         `gorm:"created_by" json:"created_by"`
	UpdatedBy     string         `gorm:"updated_by" json:"updated_by"`
	DeletedBy     string         `gorm:"deleted_by" json:"deleted_by"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
