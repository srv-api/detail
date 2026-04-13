package pin

import (
	dto "github.com/srv-api/merchant/dto"
	"github.com/srv-api/merchant/entity"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Create(req dto.PinRequest) (dto.PinResponse, error)
	Delete(req dto.DeletePinRequest) (dto.DeletePinResponse, error)
	BulkDelete(req dto.BulkDeletePinRequest) (int, error)
	Get(req *dto.Pagination) (RepositoryResult, int)
	GetById(req dto.GetByIdPinRequest) (*dto.PinResponse, error)
	Update(req dto.PinUpdateRequest) (dto.PinUpdateResponse, error)
	Verify(req dto.VerifyPinRequest) (*entity.Pin, error)
	GetPinStatus(req dto.PinRequest) (*entity.Pin, error)
}

type pinRepository struct {
	DB *gorm.DB
}

func NewPinRepository(DB *gorm.DB) DomainRepository {
	return &pinRepository{
		DB: DB,
	}
}
