package deleteaccount

import (
	dto "github.com/srv-api/merchant/dto"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Create(req dto.DeleteAccountRequest) (dto.DeleteAccountResponse, error)
	Get(req *dto.Pagination) (RepositoryResult, int)
	GetById(req dto.GetDeleteAccountByIdRequest) (*dto.DeleteAccountResponse, error)
	Delete(req dto.DeleteDeleteAccountRequest) (dto.DeleteDeleteAccountResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (int, error)
	Update(req dto.DeleteAccountUpdateRequest) (dto.DeleteAccountUpdateResponse, error)
}

type deleteaccountRepository struct {
	DB *gorm.DB
}

func NewDeleteAccountRepository(DB *gorm.DB) DomainRepository {
	return &deleteaccountRepository{
		DB: DB,
	}
}
