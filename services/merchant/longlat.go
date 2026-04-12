package product

import (
	dto "github.com/srv-api/merchant/dto"
)

func (b *merchantService) LongLat(req dto.UpdateUserDetailRequest) (dto.UpdateUserDetailResponse, error) {

	request := dto.UpdateUserDetailRequest{
		ID:        req.ID,
		Longitude: req.Longitude,
		Latitude:  req.Latitude,
		UpdatedBy: req.UpdatedBy,
	}

	product, err := b.Repo.Update(req)
	if err != nil {
		return product, err
	}

	response := dto.UpdateUserDetailResponse{
		ID:        request.ID,
		Longitude: request.Longitude,
		Latitude:  request.Latitude,
		UpdatedBy: request.UpdatedBy,
	}

	return response, nil
}
