package product

import (
	dto "github.com/srv-api/merchant/dto"
)

func (b *merchantService) Update(req dto.UpdateUserDetailRequest) (dto.UpdateUserDetailResponse, error) {

	request := dto.UpdateUserDetailRequest{
		ID:           req.ID,
		UserID:       req.UserID,
		Radius:       req.Radius,
		MinAge:       req.MinAge,
		MaxAge:       req.MaxAge,
		GenderTarget: req.GenderTarget,
		UpdatedBy:    req.UpdatedBy,
	}

	product, err := b.Repo.Update(req)
	if err != nil {
		return product, err
	}

	response := dto.UpdateUserDetailResponse{
		ID:           request.ID,
		UserID:       req.UserID,
		UpdatedBy:    request.UpdatedBy,
		Radius:       request.Radius,
		MinAge:       request.MinAge,
		MaxAge:       request.MaxAge,
		GenderTarget: request.GenderTarget,
	}

	return response, nil
}
