package userdetail

import (
	"errors"

	dto "github.com/srv-api/detail/dto"
)

func (s *merchantService) Explore(req dto.UserDetailRequest) (dto.ExploreResponse, error) {
	// 🔵 AMBIL DATA LIMIT USER DARI DATABASE
	userLimit, err := s.Repo.GetUserLimit(req.UserID)
	if err != nil {
		return dto.ExploreResponse{}, err
	}

	// 🟢 CEK: Apakah masih ada swipe?
	if userLimit.RemainingSwipe <= 0 {
		// Kembalikan error kalau limit habis
		return dto.ExploreResponse{}, errors.New("daily swipe limit exceeded")
	}

	// 📋 Kalau masih ada swipe, lanjut ambil data explore
	users, err := s.Repo.Explore(req)
	if err != nil {
		return dto.ExploreResponse{}, err
	}

	// ✅ Kembalikan data user + sisa swipe (TANPA mengurangi)
	return dto.ExploreResponse{
		Users: users,
	}, nil
}
