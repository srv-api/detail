package permission

import (
	dto "github.com/srv-api/merchant/dto"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Create(req dto.PermissionRequest) (dto.PermissionResponse, error)
	Get(req dto.RoleUserRequest) (dto.GetPermissionResponse, error)
	Pagination(req *dto.Pagination) (RepositoryResult, int)
	GetById(req dto.GetPermissionByIdRequest) (*dto.PermissionResponse, error)
	Delete(req dto.DeletePermissionRequest) (dto.DeleteResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (int, error)
	Update(req dto.PermissionUpdateRequest) (dto.PermissionUpdateResponse, error)
}

type PermissionRepository struct {
	DB *gorm.DB
}

func NewPermissionRepository(DB *gorm.DB) DomainRepository {
	return &PermissionRepository{
		DB: DB,
	}
}
