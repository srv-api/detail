package paymentmethod

import (
	s "github.com/srv-api/merchant/services/subscribe/paymentmethod"

	"github.com/labstack/echo/v4"
)

type DomainHandler interface {
	Create(c echo.Context) error
	Get(c echo.Context) error
	Delete(c echo.Context) error
	BulkDelete(c echo.Context) error
	GetById(c echo.Context) error
	Update(c echo.Context) error
	GetPicture(c echo.Context) error
}

type domainHandler struct {
	servicePayment s.PaymentMethodService
}

func NewPaymentHandler(service s.PaymentMethodService) DomainHandler {
	return &domainHandler{
		servicePayment: service,
	}
}
