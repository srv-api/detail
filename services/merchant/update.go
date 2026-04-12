package product

import (
	dto "github.com/srv-api/merchant/dto"
)

func (b *merchantService) Update(req dto.UpdateUserDetailRequest) (dto.UpdateUserDetailResponse, error) {

	request := dto.UpdateUserDetailRequest{
		ID:           req.ID,
		UserID:       req.UserID,
		Longitude:    req.Longitude,
		Latitude:     req.Latitude,
		Radius:       req.Radius,
		MinAge:       req.MinAge,
		MaxAge:       req.MaxAge,
		GenderTarget: req.GenderTarget,
		UpdatedBy:    req.UpdatedBy,
	}

	product, err := b.Repo.LongLat(req)
	if err != nil {
		return product, err
	}

	response := dto.UpdateUserDetailResponse{
		ID:        request.ID,
		UserID:    req.UserID,
		Longitude: request.Longitude,
		Latitude:  request.Latitude,
		UpdatedBy: request.UpdatedBy,
	}

	return response, nil
}
