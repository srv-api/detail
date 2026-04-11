package permission

import (
	s "github.com/srv-api/merchant/services/dashboard/permission"

	"github.com/labstack/echo/v4"
)

type DomainHandler interface {
	Create(c echo.Context) error
	Get(c echo.Context) error
	Delete(c echo.Context) error
	BulkDelete(c echo.Context) error
	GetById(c echo.Context) error
	Update(c echo.Context) error
	Pagination(c echo.Context) error
}

type domainHandler struct {
	servicePermission s.PermissionService
}

func NewPermissionHandler(service s.PermissionService) DomainHandler {
	return &domainHandler{
		servicePermission: service,
	}
}
