package merchant

import (
	dto "github.com/srv-api/merchant/dto"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Get(req dto.UserDetailRequest) (dto.UserDetailResponse, error)
	GetById(req dto.GetUserDetailByIdRequest) (*dto.UserDetailRequest, error)
	Update(req dto.UpdateUserDetailRequest) (dto.UpdateUserDetailResponse, error)
}

type merchantRepository struct {
	DB *gorm.DB
}

func NewMerchantRepository(DB *gorm.DB) DomainRepository {
	return &merchantRepository{
		DB: DB,
	}
}
