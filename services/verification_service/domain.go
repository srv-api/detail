package verification_service

import (
	"github.com/srv-api/detail/dto"
)

type VerificationService interface {
	VerifyWithStore(req dto.VerifyPurchaseRequest) (bool, error)
}

type verificationService struct {
	androidPackageName string
}

func NewVerificationService(packageName string) VerificationService {
	return &verificationService{
		androidPackageName: packageName,
	}
}
