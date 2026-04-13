package product

import (
	dto "github.com/srv-api/detail/dto"
	"github.com/srv-api/detail/entity"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Create(req dto.ProductRequest) (dto.ProductResponse, error)
	Get(req *dto.Pagination) (RepositoryResult, int)
	GetById(req dto.GetByIdRequest) (*dto.ProductResponse, error)
	Delete(req dto.DeleteRequest) (dto.DeleteResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (int, error)
	BulkEdit(req dto.BulkEditRequest) (int, error)
	Update(req dto.ProductUpdateRequest) (dto.ProductUpdateResponse, error)
	SaveFile(req dto.ProductUploadRequest) (dto.ProductUploadResponse, error)
	GetPicture(req dto.GetProductUploadRequest) (*dto.GetProductUploadResponse, error)
	CheckMerchantDetail(DetailID string, merchantDetail *entity.UserDetail) error
}

type productRepository struct {
	DB *gorm.DB
}

func NewProductRepository(DB *gorm.DB) DomainRepository {
	return &productRepository{
		DB: DB,
	}
}
