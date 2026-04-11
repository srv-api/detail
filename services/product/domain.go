package product

import (
	"github.com/labstack/echo/v4"
	dto "github.com/srv-api/merchant/dto"
	m "github.com/srv-api/middlewares/middlewares"

	r "github.com/srv-api/merchant/repositories/product"
)

type ProductService interface {
	Create(req dto.ProductRequest) (dto.ProductResponse, error)
	Get(context echo.Context, req *dto.Pagination) dto.Response
	GetById(req dto.GetByIdRequest) (*dto.ProductResponse, error)
	Delete(req dto.DeleteRequest) (dto.DeleteResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (dto.BulkDeleteResponse, error)
	BulkEdit(req dto.BulkEditRequest) (dto.BulkEditResponse, error)
	Update(req dto.ProductUpdateRequest) (dto.ProductUpdateResponse, error)
	Upload(req dto.ProductUploadRequest) (dto.ProductUploadResponse, error)
	GetPicture(req dto.GetProductUploadRequest) (*dto.GetProductUploadResponse, error)
}

type productService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewProductService(Repo r.DomainRepository, jwtS m.JWTService) ProductService {
	return &productService{
		Repo: Repo,
		jwt:  jwtS,
	}
}
