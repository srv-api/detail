package premium

import (
	dto "github.com/srv-api/detail/dto"
	m "github.com/srv-api/middlewares/middlewares"

	r "github.com/srv-api/detail/repositories/purchase/premium"
)

type PurchaseService interface {
	Create(req dto.PremiumPurchaseRequest) (dto.PremiumPurchaseResponse, error)
}

type premiumPurchase struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewPurchaseService(Repo r.DomainRepository, jwtS m.JWTService) PurchaseService {
	return &premiumPurchase{
		Repo: Repo,
		jwt:  jwtS,
	}
}
