package history

import (
	s "github.com/srv-api/merchant/services/subscribe/history"

	"github.com/labstack/echo/v4"
)

type DomainHandler interface {
	Get(c echo.Context) error
	GetById(c echo.Context) error
	CheckExpire(c echo.Context) error
}

type domainHandler struct {
	serviceHistory s.HistoryService
}

func NewHistoryHandler(service s.HistoryService) DomainHandler {
	return &domainHandler{
		serviceHistory: service,
	}
}
