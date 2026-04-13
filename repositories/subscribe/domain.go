package subscribe

import (
	"github.com/plutov/paypal/v4"
	dto "github.com/srv-api/merchant/dto"
	"gorm.io/gorm"
)

type DomainRepository interface {
	Create(req dto.SubscribeRequest) (dto.SubscribeResponse, error)
	UpdateStatus(orderID string, status string) error
	UpdateUserVerified(orderID string) error
	ChargeQris(req dto.ChargeRequest) (*dto.QrisResponse, error)
	CheckStatus(orderID string) (map[string]interface{}, error)
	ChargePermata(req dto.ChargeRequest) (*dto.VAPermataResponse, error)
	ChargeBri(req dto.ChargeRequest) (*dto.VAResponse, error)
	ChargeCimb(req dto.ChargeRequest) (*dto.VAResponse, error)
	ChargeBni(req dto.ChargeRequest) (*dto.VAResponse, error)
	CardPayment(cardData dto.CreditCardChargeRequest) (*dto.TokenizeResponse, error)
	ChargeGopay(req dto.ChargeRequest) (*dto.GopayResponse, error)
	CancelPay(req dto.GetorderID) ([]byte, int, error)
	UpdateSubscribeByOrderID(data dto.MidtransCancelResponse) error

	// Tambahan untuk PayPal
	CreatePaypalOrder(req dto.PaypalCreateRequest) (*paypal.Order, error)
	CapturePaypalOrder(orderID string) (*paypal.CaptureOrderResponse, error)
	UpdateSubscribeStatus(orderID string, status string) error
}

// type subscribeRepository struct {
// 	DB           *gorm.DB
// 	paypalClient *paypal.Client
// }

// func NewSubscribeRepository(DB *gorm.DB, paypalClient *paypal.Client) DomainRepository {
// 	return &subscribeRepository{
// 		DB:           DB,
// 		paypalClient: paypalClient,
// 	}
// }

type subscribeRepository struct {
	DB           *gorm.DB
	paypalClient *paypal.Client
}

func NewSubscribeRepository(DB *gorm.DB, paypalClient *paypal.Client) DomainRepository {
	return &subscribeRepository{
		DB:           DB,
		paypalClient: paypalClient,
	}
}
