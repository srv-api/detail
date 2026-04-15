package service

import (
	repository "github.com/srv-api/detail/repositories/match"
)

type MatchService interface {
	CreateMatch(user1, user2 string) error
	GetMatches(userID string) ([]map[string]interface{}, error)
}

type matchService struct {
	Repo repository.MatchRepository
}

func NewMatchService(repo repository.MatchRepository) MatchService {
	return &matchService{Repo: repo}
}

func (s *matchService) CreateMatch(user1, user2 string) error {
	return s.Repo.CreateMatch(user1, user2)
}

func (s *matchService) GetMatches(userID string) ([]map[string]interface{}, error) {
	return s.Repo.GetMatches(userID)
}
