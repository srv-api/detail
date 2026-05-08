package dto

import "time"

type PremiumPurchaseRequest struct {
	ID            string `json:"id"`
	UserID        string `json:"user_id"`
	DetailID      string `json:"detail_id"`
	ProductID     string `json:"product_id"`
	TransactionID string `json:"transaction_id"`
	PurchaseToken string `json:"purchase_token"`
	ReceiptData   string `json:"receipt_data"`
	Signature     string `json:"signature"`
	CreatedBy     string `json:"created_by"`
}

type PremiumPurchaseResponse struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	DetailID      string    `json:"detail_id"`
	ProductID     string    `json:"product_id"`
	TransactionID string    `json:"transaction_id"`
	IsPremium     bool      `json:"is_premium"`
	ExpiredAt     time.Time `json:"expired_at"`
	PurchaseToken string    `json:"purchase_token"`
	ReceiptData   string    `json:"receipt_data"`
	Signature     string    `json:"signature"`
	CreatedBy     string    `json:"created_by"`
}

type GooglePayVerificationRequest struct {
	TransactionID string `json:"transaction_id"`
	PurchaseToken string `json:"purchase_token"`
	ProductID     string `json:"product_id"`
	ReceiptData   string `json:"receipt_data"`
	Signature     string `json:"signature"`
}

type GooglePayVerificationResponse struct {
	IsValid          bool   `json:"is_valid"`
	PurchaseState    int    `json:"purchase_state"`
	ConsumptionState int    `json:"consumption_state"`
	Kind             string `json:"kind"`
}
