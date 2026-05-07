package premium

import (
	s "github.com/srv-api/detail/services/purchase/premium"

	"github.com/labstack/echo/v4"
)

type DomainHandler interface {
	Create(c echo.Context) error
}

type domainHandler struct {
	PremiumPurchase s.PurchaseService
}

func NewPinHandler(service s.PurchaseService) DomainHandler {
	return &domainHandler{
		PremiumPurchase: service,
	}
}
