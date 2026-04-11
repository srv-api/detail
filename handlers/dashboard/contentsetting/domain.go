package contentsetting

import (
	s "github.com/srv-api/merchant/services/dashboard/contentsetting"

	"github.com/labstack/echo/v4"
)

type DomainHandler interface {
	Get(c echo.Context) error
	Update(c echo.Context) error
}

type domainHandler struct {
	serviceContentSetting s.ContentSettingService
}

func NewContentSettingHandler(service s.ContentSettingService) DomainHandler {
	return &domainHandler{
		serviceContentSetting: service,
	}
}
