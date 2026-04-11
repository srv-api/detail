package subscribe

import (
	s "github.com/srv-api/merchant/services/subscribe"

	"github.com/labstack/echo/v4"
)

type DomainHandler interface {
	TokenizeCardHandler(c echo.Context) error
	CardPayment(c echo.Context) error

	Create(c echo.Context) error
	MidtransCallback(c echo.Context) error
	ChargeBni(c echo.Context) error
	ChargePermata(c echo.Context) error
	ChargeMandiri(c echo.Context) error
	ChargeBri(c echo.Context) error
	ChargeCimb(c echo.Context) error

	ChargeGpay(c echo.Context) error

	ChargeQris(c echo.Context) error
	ChargeGopay(c echo.Context) error
	ChargeShopeePay(c echo.Context) error

	CheckTransactionStatus(c echo.Context) error
	CancelPay(c echo.Context) error

	PayPal(c echo.Context) error
	CapturePaypalOrder(c echo.Context) error
}

type domainHandler struct {
	serviceSubscribe s.SubscribeService
}

func NewSubscribeHandler(service s.SubscribeService) DomainHandler {
	return &domainHandler{
		serviceSubscribe: service,
	}
}
