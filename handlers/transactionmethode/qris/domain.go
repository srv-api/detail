package qris

import (
	s "github.com/srv-api/merchant/services/transactionmethode/qris"

	"github.com/labstack/echo/v4"
)

type DomainHandler interface {
	Create(c echo.Context) error
}

type domainHandler struct {
	serviceQris s.QrisService
}

func NewQrisHandler(service s.QrisService) DomainHandler {
	return &domainHandler{
		serviceQris: service,
	}
}
