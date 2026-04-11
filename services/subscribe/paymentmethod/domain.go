package paymentmethod

import (
	dto "github.com/srv-api/merchant/dto"
	m "github.com/srv-api/middlewares/middlewares"

	r "github.com/srv-api/merchant/repositories/subscribe/paymentmethod"
)

type PaymentMethodService interface {
	Create(req dto.PaymentMethodRequest) (dto.PaymentMethodResponse, error)
	Get(req dto.PaymentMethodRequest) ([]dto.PaymentMethodResponse, error)
	GetById(req dto.GetByIdPaymentRequest) (dto.PaymentMethodResponse, error)
	Delete(req dto.DeletePaymentRequest) (dto.DeletePaymentResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (dto.BulkDeleteResponse, error)
	Update(req dto.PaymentMethodUpdateRequest) (dto.PaymentMethodUpdateResponse, error)
	GetPicture(req dto.GetPaymentploadRequest) (*dto.GetPaymentUploadResponse, error)
}

type paymentmethodService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewPaymentMethodService(Repo r.DomainRepository, jwtS m.JWTService) PaymentMethodService {
	return &paymentmethodService{
		Repo: Repo,
		jwt:  jwtS,
	}
}
