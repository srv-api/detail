package contentsetting

import (
	"time"

	dto "github.com/srv-api/merchant/dto"
)

func (s *contentsettingService) Update(req dto.UpdateContentSettingRequest) (dto.UpdateContentSettingResponse, error) {
	err := s.Repo.Update(req)
	if err != nil {
		return dto.UpdateContentSettingResponse{}, err
	}

	return dto.UpdateContentSettingResponse{
		ID:           req.ID,
		UserID:       req.UserID,
		MerchantID:   req.MerchantID,
		TopHeader:    req.TopHeader,
		ButtonHeader: req.ButtonHeader,
		Feature:      req.Feature,
		Footer:       req.Footer,
		UpdatedBy:    req.UpdatedBy,
		UpdatedAt:    time.Now(),
	}, nil
}
