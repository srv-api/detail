package service

import (
	"errors"

	"github.com/srv-api/detail/dto"
	repository "github.com/srv-api/detail/repositories/like"
)

type LikeService interface {
	LikeUser(userID string, req dto.LikeRequest) (dto.LikeResponse, error)
}

type likeService struct {
	Repo repository.LikeRepository
}

func NewLikeService(repo repository.LikeRepository) LikeService {
	return &likeService{Repo: repo}
}

func (s *likeService) LikeUser(userID string, req dto.LikeRequest) (dto.LikeResponse, error) {

	// validasi diri sendiri
	if userID == req.TargetUserID {
		return dto.LikeResponse{}, errors.New("cannot like yourself")
	}

	// simpan like
	err := s.Repo.CreateLike(userID, req.TargetUserID)
	if err != nil {
		return dto.LikeResponse{}, err
	}

	// cek match
	isMatch, err := s.Repo.IsMatch(userID, req.TargetUserID)
	if err != nil {
		return dto.LikeResponse{}, err
	}

	if isMatch {
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
