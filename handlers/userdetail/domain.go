package userdetail

import (
	s "github.com/srv-api/merchant/services/userdetail"

	"github.com/labstack/echo/v4"
)

type DomainHandler interface {
	Get(c echo.Context) error
	Explore(c echo.Context) error
	Update(c echo.Context) error
	LongLat(c echo.Context) error
}

type domainHandler struct {
	serviceUserDetail s.UserDetailService
}

func NewUserDetailHandler(service s.UserDetailService) DomainHandler {
	return &domainHandler{
		serviceUserDetail: service,
	}
}
