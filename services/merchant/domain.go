package product

import (
	dto "github.com/srv-api/merchant/dto"
	m "github.com/srv-api/middlewares/middlewares"

	r "github.com/srv-api/merchant/repositories/merchant"
)

type MerchantService interface {
	Get(req dto.UserDetailRequest) (dto.UserDetailResponse, error)
	Update(req dto.UpdateUserDetailRequest) (dto.UpdateUserDetailResponse, error)
	LongLat(req dto.UpdateUserDetailRequest) (dto.UpdateUserDetailResponse, error)
}

type merchantService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewMerchantService(Repo r.DomainRepository, jwtS m.JWTService) MerchantService {
	return &merchantService{
		Repo: Repo,
		jwt:  jwtS,
	}
}
