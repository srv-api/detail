package userdetail

import (
	dto "github.com/srv-api/detail/dto"
	m "github.com/srv-api/middlewares/middlewares"

	r "github.com/srv-api/detail/repositories/userdetail"
)

type UserDetailService interface {
	Get(req dto.UserDetailRequest) (dto.UserDetailResponse, error)
	Explore(req dto.UserDetailRequest) (dto.ExploreResponse, error)
	Update(req dto.UpdateUserDetailRequest) (dto.UpdateUserDetailResponse, error)
	LongLat(req dto.UpdateUserDetailRequest) (dto.UpdateUserDetailResponse, error)
}

type merchantService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewUserDetailService(Repo r.DomainRepository, jwtS m.JWTService) UserDetailService {
	return &merchantService{
		Repo: Repo,
		jwt:  jwtS,
	}
}
