package pin

import (
	"errors"

	"github.com/srv-api/merchant/dto"
	"gorm.io/gorm"
)

func (s *pinService) GetPinStatus(req dto.PinRequest) (dto.PinStatusResponse, error) {
	pin, err := s.Repo.GetPinStatus(req)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return dto.PinStatusResponse{}, err
	}

	status := dto.PinStatusResponse{
		IsPinEnabled: pin != nil,
	}

	return status, nil
}
