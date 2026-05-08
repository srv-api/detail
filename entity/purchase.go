package entity

import "time"

type PurchasePremium struct {
	ID            string    `gorm:"type:varchar(36);primaryKey"`
	UserID        string    `gorm:"type:varchar(50);index;not null"`
	DetailID      string    `gorm:"type:varchar(36);index" json:"detail_id"`
	ProductID     string    `gorm:"product_id" json:"product_id"`
	TransactionID string    `gorm:"transaction_id" json:"transaction_id"`
	PurchaseToken string    `gorm:"purchase_token" json:"purchase_token"`
	ReceiptData   string    `gorm:"receipt_data" json:"receipt_data"`
	Signature     string    `gorm:"signature" json:"signature"`
	CreatedBy     string    `json:"created_by"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}
