package pin

import (
	s "github.com/srv-api/merchant/services/pin"

	"github.com/labstack/echo/v4"
)

type DomainHandler interface {
	Create(c echo.Context) error
	Get(c echo.Context) error
	GetById(c echo.Context) error
	BulkDelete(c echo.Context) error
	Delete(c echo.Context) error
	Update(c echo.Context) error
	VerifyPIN(c echo.Context) error
	GetPinStatus(c echo.Context) error
}

type domainHandler struct {
	servicePin s.PinService
}

func NewPinHandler(service s.PinService) DomainHandler {
	return &domainHandler{
		servicePin: service,
	}
}
