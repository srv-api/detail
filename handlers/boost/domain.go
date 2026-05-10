package premium

import (
	s "github.com/srv-api/detail/services/boost"

	"github.com/labstack/echo/v4"
)

type DomainHandler interface {
	Create(c echo.Context) error
}

type domainHandler struct {
	serviceBoost s.BoostService
}

func NewBoostHandler(service s.BoostService) DomainHandler {
	return &domainHandler{
		serviceBoost: service,
	}
}
