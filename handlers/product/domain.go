package product

import (
	s "github.com/srv-api/merchant/services/product"

	"github.com/labstack/echo/v4"
)

type DomainHandler interface {
	Create(c echo.Context) error
	Get(c echo.Context) error
	Delete(c echo.Context) error
	BulkDelete(c echo.Context) error
	BulkEdit(c echo.Context) error
	GetById(c echo.Context) error
	Update(c echo.Context) error
	UploadImage(c echo.Context) error
	GetPicture(c echo.Context) error
}

type domainHandler struct {
	serviceProduct s.ProductService
}

func NewProductHandler(service s.ProductService) DomainHandler {
	return &domainHandler{
		serviceProduct: service,
	}
}
