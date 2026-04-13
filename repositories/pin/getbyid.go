package pin

import (
	dto "github.com/srv-api/detail/dto"
	"github.com/srv-api/detail/entity"
)

func (b *pinRepository) GetById(req dto.GetByIdPinRequest) (*dto.PinResponse, error) {
	tr := entity.Pin{
		ID: req.ID,
	}

	if err := b.DB.Where("id = ?", tr.ID).Take(&tr).Error; err != nil {
		return nil, err
	}

	response := &dto.PinResponse{
		Pin: tr.Pin,
	}

	return response, nil
}
