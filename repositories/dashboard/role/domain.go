package role

import (
	dto "github.com/srv-api/merchant/dto"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Create(req dto.RoleRequest) (dto.RoleResponse, error)
	Get(req dto.RoleUserRequest) (dto.GetRoleResponse, error)
	Pagination(req *dto.Pagination) (RepositoryResult, int)
	GetById(req dto.GetRoleByIdRequest) (*dto.RoleResponse, error)
	Delete(req dto.DeleteRoleRequest) (dto.DeleteResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (int, error)
	Update(req dto.RoleUpdateRequest) (dto.RoleUpdateResponse, error)
	RoleUser(req dto.GetRoleRequest) (dto.GetRoleResponse, error)
}

type RoleRepository struct {
	DB *gorm.DB
}

func NewRoleRepository(DB *gorm.DB) DomainRepository {
	return &RoleRepository{
		DB: DB,
	}
}
