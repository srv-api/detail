package userdetail

import (
	"errors"

	dto "github.com/srv-api/detail/dto"
)

func (s *merchantService) Explore(req dto.UserDetailRequest) (dto.ExploreResponse, error) {
	users, err := s.Repo.Explore(req)
	if err != nil {
		return dto.ExploreResponse{}, err
	}
	return dto.ExploreResponse{Users: users}, nil
}

// Contoh di service/like_service.go
func (s *merchantService) LikeUser(userID, targetUserID string, isSuperLike bool) error {
	// Cek limit user
	userLimit, err := s.Repo.GetUserLimit(userID)
	if err != nil {
		return err
	}

	if isSuperLike {
		if userLimit.RemainingSuperLike <= 0 {
			return errors.New("no super like remaining")
		}
		// Proses super like
		if err := s.userDetailRepo.DeductSuperLike(userID); err != nil {
			return err
		}
	} else {
		if userLimit.RemainingSwipe <= 0 {
			return errors.New("daily swipe limit exceeded")
		}
		// Proses like biasa
		if err := s.userDetailRepo.DeductSwipe(userID); err != nil {
			return err
		}
	}

	// Lanjutkan proses like...
	return nil
}
