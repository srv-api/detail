package contentsetting

import (
	dto "github.com/srv-api/merchant/dto"
)

func (s *contentsettingService) Get(req dto.ContentSettingRequest) (dto.ContentSettingResponse, error) {
	// Fetch comments from the repository layer based on post_id
	cs, err := s.Repo.Get(req)
	if err != nil {
		return dto.ContentSettingResponse{}, err
	}

	return cs, nil
}
