package subscribe

import (
	dto "github.com/srv-api/merchant/dto"
	m "github.com/srv-api/middlewares/middlewares"

	r "github.com/srv-api/merchant/repositories/subscribe"
)

type SubscribeService interface {
	Create(req dto.SubscribeRequest) (dto.SubscribeResponse, error)
	UpdateStatus(orderID string, transactionStatus string) error
	ChargeQris(req dto.ChargeRequest) (*dto.QrisResponse, error)
	CheckTransactionStatus(request dto.CreateTransactionRequest) (dto.TransactionStatusResponse, error)
	ChargePermata(req dto.ChargeRequest) (*dto.VAPermataResponse, error)
	ChargeBri(req dto.ChargeRequest) (*dto.VAResponse, error)
	ChargeCimb(req dto.ChargeRequest) (*dto.VAResponse, error)
	ChargeBni(req dto.ChargeRequest) (*dto.VAResponse, error)
	CardPayment(cardData dto.CreditCardChargeRequest) (*dto.TokenizeResponse, error)
	ChargeGopay(req dto.ChargeRequest) (*dto.GopayResponse, error)
	CancelPay(req dto.GetorderID) (map[string]interface{}, int, error)
	CreatePaypalOrder(req dto.PaypalCreateRequest) (*dto.PaypalOrderResponse, error)
	CapturePaypalOrder(orderID string) (*dto.PaypalCaptureResponse, error)
}

type subscribeService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewSubscribeService(Repo r.DomainRepository, jwtS m.JWTService) SubscribeService {
	return &subscribeService{
		Repo: Repo,
		jwt:  jwtS,
	}
}
