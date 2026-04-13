package userdetail

import (
	dto "github.com/srv-api/merchant/dto"
)

func (s *merchantService) Explore(req dto.UserDetailRequest) (dto.ExploreResponse, error) {
	users, err := s.Repo.Explore(req)
	if err != nil {
		return dto.ExploreResponse{}, err
	}
	return dto.ExploreResponse{Users: users}, nil
}
