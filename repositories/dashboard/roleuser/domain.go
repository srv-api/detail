package roleuser

import (
	dto "github.com/srv-api/merchant/dto"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Create(req dto.RoleUserRequest) (dto.RoleUserResponse, error)
	Get(req dto.RoleUserRequest) (dto.GetRoleUserResponse, error)
	Pagination(req *dto.Pagination) (RepositoryResult, int)
	GetById(req dto.GetRoleUserByIdRequest) (*dto.RoleUserResponse, error)
	Delete(req dto.DeleteRoleUserRequest) (dto.DeleteResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (int, error)
	Update(req dto.RoleUserUpdateRequest) (dto.RoleUserUpdateResponse, error)
}

type RoleUserRepository struct {
	DB *gorm.DB
}

func NewRoleUserRepository(DB *gorm.DB) DomainRepository {
	return &RoleUserRepository{
		DB: DB,
	}
}
