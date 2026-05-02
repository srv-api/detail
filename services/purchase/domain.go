package purchase

import (
	dto "github.com/srv-api/detail/dto"
	l "github.com/srv-api/detail/repositories/userlimit"
	vs "github.com/srv-api/detail/services/verification_service"
	m "github.com/srv-api/middlewares/middlewares"

	r "github.com/srv-api/detail/repositories/purchase"
)

type PurchaseService interface {
	ProcessPurchase(req dto.VerifyPurchaseRequest) (*dto.PurchaseResponse, error)
	GetProductAmount(productID string) float64
	GeneratePurchaseID() string
}

type purchaseService struct {
	userLimitRepo   l.DomainRepository
	Repo            r.DomainRepository
	jwt             m.JWTService
	verificationSvc vs.VerificationService
}

func NewPurchaseService(Repo r.DomainRepository, userLimitRepo l.DomainRepository, verificationSvc vs.VerificationService, jwtS m.JWTService) PurchaseService {
	return &purchaseService{
		userLimitRepo:   userLimitRepo,
		Repo:            Repo,
		jwt:             jwtS,
		verificationSvc: verificationSvc,
	}
}
