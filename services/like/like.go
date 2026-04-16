package service

import (
	"errors"

	"github.com/srv-api/detail/dto"
	repository "github.com/srv-api/detail/repositories/like"
	matchService "github.com/srv-api/detail/services/match"
)

type LikeService interface {
	LikeUser(userID string, req dto.LikeRequest) (dto.LikeResponse, error)
}

type likeService struct {
	Repo         repository.LikeRepository
	MatchService matchService.MatchService
}

func NewLikeService(
	repo repository.LikeRepository,
	matchService matchService.MatchService,
) LikeService {
	return &likeService{
		Repo:         repo,
		MatchService: matchService,
	}
}

func (s *likeService) LikeUser(userID string, req dto.LikeRequest) (dto.LikeResponse, error) {

	if userID == req.TargetUserID {
		return dto.LikeResponse{}, errors.New("cannot like yourself")
	}

	// 🔥 1. Deduct limit dulu
	if req.IsSuperLike {
		if err := s.Repo.DeductSuperLike(userID); err != nil {
			return dto.LikeResponse{}, err
		}
	} else {
		if err := s.Repo.DeductSwipe(userID); err != nil {
			return dto.LikeResponse{}, err
		}
	}

	// 🔥 2. Insert like
	err := s.Repo.CreateLike(userID, req.TargetUserID)
	if err != nil {
		return dto.LikeResponse{}, err
	}

	// 🔥 3. Check match
	isMatch, err := s.Repo.IsMatch(userID, req.TargetUserID)
	if err != nil {
		return dto.LikeResponse{}, err
	}

	if isMatch {
		err := s.MatchService.CreateMatch(userID, req.TargetUserID)
		if err != nil {
			return dto.LikeResponse{}, err
		}

		return dto.LikeResponse{
			IsMatch: true,
			Message: "It's a match! 🎉",
		}, nil
	}

	return dto.LikeResponse{
		IsMatch: false,
		Message: "Liked successfully",
	}, nil
}
