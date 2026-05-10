package boost

import (
	dto "github.com/srv-api/detail/dto"
	"github.com/srv-api/detail/entity"
)

func (r *boostRepository) Create(req dto.BoostRequest) (dto.BoostResponse, error) {

	create := entity.Boost{
		ID:        req.ID,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		UserID:    req.UserID,
	}

	if err := r.DB.Save(&create).Error; err != nil {
		return dto.BoostResponse{}, err
	}

	response := dto.BoostResponse{
		ID:        req.ID,
		StartTime: create.StartTime,
		EndTime:   create.EndTime,
		UserID:    req.UserID,
	}

	return response, nil

}
