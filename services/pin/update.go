package pin

import (
	"github.com/srv-api/detail/dto"
)

func (b *pinService) Update(req dto.PinUpdateRequest) (dto.PinUpdateResponse, error) {
	request := dto.PinUpdateRequest{
		Pin:         req.Pin,
		UserID:      req.UserID,
		DetailID:    req.DetailID,
		UpdatedBy:   req.UpdatedBy,
		Description: req.Description,
		Status:      req.Status,
	}

	product, err := b.Repo.Update(req)
	if err != nil {
		return product, err
	}

	response := dto.PinUpdateResponse{
		Pin:         request.Pin,
		UserID:      request.UserID,
		DetailID:    request.DetailID,
		UpdatedBy:   request.UpdatedBy,
		Description: request.Description,
		Status:      request.Status,
	}

	return response, nil
}
