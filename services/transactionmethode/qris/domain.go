package qris

import (
	dto "github.com/srv-api/merchant/dto"
	m "github.com/srv-api/middlewares/middlewares"

	r "github.com/srv-api/merchant/repositories/transactionmethode/qris"
)

type QrisService interface {
	Create(req dto.CoQrisRequest) (dto.CoQrisResponse, error)
}

type qrisService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewQrisService(Repo r.DomainRepository, jwtS m.JWTService) QrisService {
	return &qrisService{
		Repo: Repo,
		jwt:  jwtS,
	}
}
