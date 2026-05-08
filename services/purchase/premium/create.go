package premium

import (
	dto "github.com/srv-api/detail/dto"
	util "github.com/srv-api/util/s"
)

const UNLIMITED = -1

func (s *premiumPurchase) Create(req dto.PremiumPurchaseRequest) (dto.PremiumPurchaseResponse, error) {
	// Proses pembuatan data Pin
	create := dto.PremiumPurchaseRequest{
		ID:            util.GenerateRandomString(),
		UserID:        req.UserID,
		DetailID:      req.DetailID,
		ProductID:     req.ProductID,
		TransactionID: req.TransactionID,
		PurchaseToken: req.PurchaseToken,
		ReceiptData:   req.ReceiptData,
		Signature:     req.Signature,
		CreatedBy:     req.CreatedBy,
	}

	// Create purchase record AND update user limit to unlimited
	// Semua logic transaction sudah di handle di repository
	created, err := s.Repo.CreateWithPremium(create)
	if err != nil {
		return dto.PremiumPurchaseResponse{}, err
	}

	return created, nil
}
