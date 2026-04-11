package deleteaccount

import (
	s "github.com/srv-api/merchant/services/deleteaccount"

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
	serviceDeleteAccount s.DeleteAccountService
}

func NewRequestDeleteHandler(service s.DeleteAccountService) DomainHandler {
	return &domainHandler{
		serviceDeleteAccount: service,
	}
}
