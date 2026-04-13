package user

import (
	dto "github.com/srv-api/merchant/dto"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Create(req dto.UserRequest) (dto.UserResponse, error)
	Get(req *dto.Pagination) (dto.UserPaginationResponse, int)
	GetById(req dto.GetByIdRequest) (*dto.UserResponse, error)
	Delete(req dto.DeleteRequest) (dto.DeleteResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (int, error)
	Update(req dto.UserUpdateRequest) (dto.UserUpdateResponse, error)
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) DomainRepository {
	return &userRepository{
		DB: DB,
	}
}
