package pin

import (
	dto "github.com/srv-api/detail/dto"
	"github.com/srv-api/detail/entity"
)

func (r *pinRepository) Create(req dto.PinRequest) (dto.PinResponse, error) {

	create := entity.Pin{
		ID:        req.ID,
		Pin:       req.Pin,
		DetailID:  req.DetailID,
		UserID:    req.UserID,
		CreatedBy: req.CreatedBy,
	}

	if err := r.DB.Save(&create).Error; err != nil {
		return dto.PinResponse{}, err
	}

	response := dto.PinResponse{
		ID:        req.ID,
		Pin:       create.Pin,
		DetailID:  req.DetailID,
		UserID:    req.UserID,
		CreatedBy: req.CreatedBy,
	}

	return response, nil

}
