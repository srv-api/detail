package verification_service

import (
	"context"
	"errors"
	"fmt"

	"github.com/srv-api/detail/dto"
	"google.golang.org/api/androidpublisher/v3"
	"google.golang.org/api/option"
)

func (s *verificationService) VerifyWithStore(req dto.VerifyPurchaseRequest) (bool, error) {
	switch req.Platform {
	case "google":
		return s.verifyGooglePlay(req)
	case "ios":
		return s.verifyAppleAppStore(req)
	default:
		return false, errors.New("unsupported platform")
	}
}

func (s *verificationService) verifyGooglePlay(req dto.VerifyPurchaseRequest) (bool, error) {
	ctx := context.Background()

	// Gunakan service account credentials
	service, err := androidpublisher.NewService(ctx, option.WithCredentialsFile("path/to/service-account-key.json"))
	if err != nil {
		return false, fmt.Errorf("failed to create android publisher service: %w", err)
	}

	// Verifikasi purchase
	purchase, err := service.Purchases.Products.Get(
		s.androidPackageName,
		req.ProductID,
		req.PurchaseToken,
	).Context(ctx).Do()

	if err != nil {
		return false, fmt.Errorf("failed to verify purchase: %w", err)
	}

	// PurchaseState: 0 = purchased, 1 = cancelled, 2 = pending
	if purchase.PurchaseState != 0 {
		return false, errors.New("purchase not in valid state")
	}

	// Cek apakah sudah dikonsumsi? (untuk consumable products)
	if purchase.ConsumptionState == 1 { // 1 = consumed
		return false, errors.New("product already consumed")
	}

	return true, nil
}

func (s *verificationService) verifyAppleAppStore(req dto.VerifyPurchaseRequest) (bool, error) {
	// Implementasi Apple receipt validation
	// Endpoint: https://sandbox.itunes.apple.com/verifyReceipt (sandbox)
	// atau https://buy.itunes.apple.com/verifyReceipt (production)

	// Contoh implementasi sederhana
	// return verifyAppleReceipt(req.ReceiptData)

	return true, nil
}
