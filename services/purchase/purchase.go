package purchase

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/srv-api/detail/dto"
	"github.com/srv-api/detail/entity"
)

func (s *purchaseService) ProcessPurchase(req dto.VerifyPurchaseRequest) (*dto.PurchaseResponse, error) {
	// 1. Cek apakah token sudah pernah dipakai (prevent replay attack)
	exists, err := s.Repo.ExistsByToken(req.PurchaseToken)
	if err != nil {
		return nil, fmt.Errorf("failed to check existing purchase: %w", err)
	}
	if exists {
		return &dto.PurchaseResponse{
			Success: false,
			Message: "Purchase token already used",
		}, nil
	}

	// 2. Verifikasi purchase dengan Google/Apple
	isValid, err := s.verificationSvc.VerifyWithStore(req)
	if err != nil || !isValid {
		return &dto.PurchaseResponse{
			Success: false,
			Message: "Invalid purchase",
		}, nil
	}

	// 3. Dapatkan benefit berdasarkan product
	swipeBonus, superLikeBonus := s.getProductBenefits(req.ProductID)
	amount := s.GetProductAmount(req.ProductID)

	// 4. Update user limit (tambah limit atau set premium)
	err = s.userLimitRepo.IncrementLimits(req.UserID, swipeBonus, superLikeBonus)
	if err != nil {
		return nil, fmt.Errorf("failed to update user limit: %w", err)
	}

	// 5. Catat purchase history
	purchaseRecord := &entity.PurchaseHistory{
		ID:        s.GeneratePurchaseID(),
		UserID:    req.UserID,
		ProductID: req.ProductID,
		Token:     req.PurchaseToken,
		Amount:    amount,
		Status:    "completed",
		CreatedAt: time.Now(),
	}

	if err := s.Repo.Create(purchaseRecord); err != nil {
		// Log error tapi jangan rollback karena user limit sudah ke-update
		// Tapi idealnya pakai transaction
		return nil, fmt.Errorf("failed to save purchase history: %w", err)
	}

	return &dto.PurchaseResponse{
		Success: true,
		Message: "Premium activated successfully",
	}, nil
}

func (s *purchaseService) getProductBenefits(productID string) (swipeBonus int, superLikeBonus int) {
	switch productID {
	case "p1": // premium
		return 100, 10
	case "p2": // super like pack
		return 0, 5
	case "p3": // swipe pack
		return 50, 0
	default:
		return 0, 0
	}
}

func (s *purchaseService) GetProductAmount(productID string) float64 {
	switch productID {
	case "p1":
		return 9.99
	case "p2":
		return 4.99
	case "p3":
		return 2.99
	default:
		return 0
	}
}

func (s *purchaseService) GeneratePurchaseID() string {
	return uuid.New().String()
}
