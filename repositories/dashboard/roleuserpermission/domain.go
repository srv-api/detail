package roleuserpermission

import (
	dto "github.com/srv-api/merchant/dto"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Create(req dto.RoleUserPermissionRequest) (dto.RoleUserPermissionResponse, error)
	Get(req dto.RoleUserPermissionRequest) (dto.GetRoleUserPermissionResponse, error)
	Pagination(req *dto.Pagination) (RepositoryResult, int)
	GetById(req dto.GetRoleUserPermissionByIdRequest) (*dto.RoleUserPermissionResponse, error)
	Delete(req dto.DeleteRoleUserPermissionRequest) (dto.DeleteResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (int, error)
	Update(req dto.RoleUserPermissionUpdateRequest) (dto.RoleUserPermissionUpdateResponse, error)
}

type RoleUserPermissionRepository struct {
	DB *gorm.DB
}

func NewRoleUserPermissionRepository(DB *gorm.DB) DomainRepository {
	return &RoleUserPermissionRepository{
		DB: DB,
	}
}
