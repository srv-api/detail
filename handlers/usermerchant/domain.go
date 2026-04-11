package user

import (
	s "github.com/srv-api/merchant/services/usermerchant"

	"github.com/labstack/echo/v4"
)

type DomainHandler interface {
	Create(c echo.Context) error
	Get(c echo.Context) error
	Delete(c echo.Context) error
	BulkDelete(c echo.Context) error
	GetById(c echo.Context) error
	Update(c echo.Context) error
}

type domainHandler struct {
	serviceUserMerchant s.UserMerchantService
}

func NewUserMerchantHandler(service s.UserMerchantService) DomainHandler {
	return &domainHandler{
		serviceUserMerchant: service,
	}
}
