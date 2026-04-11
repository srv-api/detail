package entity

import (
	"time"

	"gorm.io/gorm"
)

type PaymentMethod struct {
	ID            uint             `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID        string           `gorm:"type:varchar(36);index" json:"user_id"`
	MerchantID    string           `gorm:"type:varchar(36);index" json:"merchant_id"`
	PaymentMethod string           `json:"payment_method"`
	PaymentLabel  string           `json:"payment_label"`
	Category      string           `json:"category" validate:"required,oneof='E-wallet' 'Virtual Account' 'Bank Transfer(VA)' 'QRIS'"`
	Image         *UploadedPayment `gorm:"foreignKey:PaymentID;references:ID" json:"image"`
	Status        int              `json:"status"`
	CreatedBy     string           `json:"created_by"`
	UpdatedBy     string           `json:"updated_by"`
	DeletedBy     string           `json:"deleted_by"`
	CreatedAt     time.Time        `json:"created_at"`
	UpdatedAt     time.Time        `json:"updated_at"`
	DeletedAt     gorm.DeletedAt   `gorm:"index" json:"deleted_at"`
}
