package dto

import "time"

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

type UserLimitResponse struct {
	UserID             string `json:"userId"`
	RemainingSwipe     int    `json:"remainingSwipe"`
	RemainingSuperLike int    `json:"remainingSuperLike"`
	IsPremium          bool   `json:"isPremium"`
}

type PremiumStatusResponse struct {
	IsPremium      bool       `json:"is_premium"`
	PremiumExpiry  *time.Time `json:"premium_expiry,omitempty"`
	RemainingSwipe int        `json:"remaining_swipe"`
}
