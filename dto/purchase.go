package dto

type VerifyPurchaseRequest struct {
	UserID        string `json:"userId"`
	PurchaseToken string `json:"purchaseToken"`
	ProductID     string `json:"productId"`
	Platform      string `json:"platform"`
	ReceiptData   string `json:"receiptData"`
}

type PurchaseResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
