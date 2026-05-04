package userdetail

import (
	"errors"

	dto "github.com/srv-api/detail/dto"
)

func (s *merchantService) Explore(req dto.UserDetailRequest) (*dto.ExploreResponse, error) {
	// 🔵 AMBIL DATA LIMIT USER DARI DATABASE
	userLimit, err := s.Repo.GetUserLimit(req.UserID)
	if err != nil {
		return nil, err
	}

	// 🟢 CEK: Apakah masih ada swipe?
	if userLimit.RemainingSwipe <= 0 {
		return nil, errors.New("daily swipe limit exceeded")
	}

	// 📋 Kalau masih ada swipe, lanjut ambil data explore
	response, err := s.Repo.Explore(req)
	if err != nil {
		return nil, err
	}

	// ✅ Tambahkan sisa swipe ke response
	response.RemainingSwipe = userLimit.RemainingSwipe

	return response, nil

}
