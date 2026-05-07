package premium

import (
	dto "github.com/srv-api/detail/dto"
	util "github.com/srv-api/util/s"
)

func (s *premiumPurchase) Create(req dto.PremiumPurchaseRequest) (dto.PremiumPurchaseResponse, error) {

	// Proses pembuatan data Pin
	create := dto.PremiumPurchaseRequest{
		ID:        util.GenerateRandomString(),
		UserID:    req.UserID,
		DetailID:  req.DetailID,
		CreatedBy: req.CreatedBy,
	}

	created, err := s.Repo.Create(create)
	if err != nil {
		return dto.PremiumPurchaseResponse{}, err
	}

	response := dto.PremiumPurchaseResponse{
		ID:        created.ID,
		DetailID:  created.DetailID,
		UserID:    created.UserID,
		CreatedBy: created.CreatedBy,
	}

	return response, nil
}
