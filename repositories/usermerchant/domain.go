package user

import (
	dto "github.com/srv-api/merchant/dto"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Create(req dto.UserMerchantRequest) (dto.UserMerchantResponse, error)
	Get(req *dto.Pagination) (dto.UserMerchantPaginationResponse, int)
	GetById(req dto.GetByIdRequest) (*dto.UserMerchantByIdResponse, error)
	Delete(req dto.DeleteRequest) (dto.DeleteResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (int, error)
	Update(req dto.UserMerchantUpdateRequest) (dto.UserMerchantUpdateResponse, error)
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserMerchantRepository(DB *gorm.DB) DomainRepository {
	return &userRepository{
		DB: DB,
	}
}
