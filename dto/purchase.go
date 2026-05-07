package dto

import "time"

type PremiumPurchaseRequest struct {
	ID          string    `json:"id"`
	UserID      string    `json:"userId"`
	DetailID    string    `json:"detail_id"`
	Platform    string    `json:"platform"`
	ReceiptData string    `json:"receiptData"`
	CreatedBy   string    `json:"created_by"`
	UpdatedAt   time.Time `json:"updated_by"`
}

type PremiumPurchaseResponse struct {
	ID          string `json:"id"`
	UserID      string `json:"userId"`
	DetailID    string `json:"detail_id"`
	Platform    string `json:"platform"`
	ReceiptData string `json:"receiptData"`
	CreatedBy   string `json:"created_by"`
}
