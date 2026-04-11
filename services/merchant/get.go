package product

import (
	dto "github.com/srv-api/merchant/dto"
)

func (s *merchantService) Get(req dto.UserDetailRequest) (dto.UserDetailResponse, error) {
	// Fetch comments from the repository layer based on post_id
	comments, err := s.Repo.Get(req)
	if err != nil {
		return dto.UserDetailResponse{}, err
	}

	return comments, nil
}
