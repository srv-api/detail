package boost

import (
	dto "github.com/srv-api/detail/dto"
	m "github.com/srv-api/middlewares/middlewares"

	r "github.com/srv-api/detail/repositories/boost"
)

type BoostService interface {
	Create(req dto.BoostRequest) (dto.BoostResponse, error)
}

type boostService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewBoostService(Repo r.DomainRepository, jwtS m.JWTService) BoostService {
	return &boostService{
		Repo: Repo,
		jwt:  jwtS,
	}
}
