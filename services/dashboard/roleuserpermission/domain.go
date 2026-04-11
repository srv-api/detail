package roleuserpermission

import (
	"github.com/labstack/echo/v4"
	dto "github.com/srv-api/merchant/dto"
	m "github.com/srv-api/middlewares/middlewares"

	r "github.com/srv-api/merchant/repositories/dashboard/roleuserpermission"
)

type RoleUserPermissionService interface {
	Create(req dto.RoleUserPermissionRequest) (dto.RoleUserPermissionResponse, error)
	Get(req dto.RoleUserPermissionRequest) (dto.GetRoleUserPermissionResponse, error)
	Pagination(context echo.Context, req *dto.Pagination) dto.Response
	GetById(req dto.GetRoleUserPermissionByIdRequest) (*dto.RoleUserPermissionResponse, error)
	Delete(req dto.DeleteRoleUserPermissionRequest) (dto.DeleteRoleUserPermissionResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (dto.BulkDeleteResponse, error)
	Update(req dto.RoleUserPermissionUpdateRequest) (dto.RoleUserPermissionUpdateResponse, error)
}

type roleuserpermissionService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewRoleUserPermissionService(Repo r.DomainRepository, jwtS m.JWTService) RoleUserPermissionService {
	return &roleuserpermissionService{
		Repo: Repo,
		jwt:  jwtS,
	}
}
