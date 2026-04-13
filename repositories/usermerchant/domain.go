package user

import (
	dto "github.com/srv-api/detail/dto"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Create(req dto.UserDetailRequest) (dto.UserDetailResponse, error)
	Get(req *dto.Pagination) (dto.UserDetailPaginationResponse, int)
	GetById(req dto.GetByIdRequest) (*dto.UserDetailByIdResponse, error)
	Delete(req dto.DeleteRequest) (dto.DeleteResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (int, error)
	Update(req dto.UserDetailUpdateRequest) (dto.UserDetailUpdateResponse, error)
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserDetailRepository(DB *gorm.DB) DomainRepository {
	return &userRepository{
		DB: DB,
	}
}
