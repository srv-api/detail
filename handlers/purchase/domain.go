package purchase

import (
	s "github.com/srv-api/detail/services/purchase"

	"github.com/labstack/echo/v4"
)

type PurchaseHandler interface {
	VerifyPurchase(c echo.Context) error
}

type purchaseHandler struct {
	servicePurchase s.PurchaseService
}

func NewPurchaseHandler(service s.PurchaseService) PurchaseHandler {
	return &purchaseHandler{
		servicePurchase: service,
	}
}
