package permission

import (
	"github.com/labstack/echo/v4"
	dto "github.com/srv-api/merchant/dto"
	m "github.com/srv-api/middlewares/middlewares"

	r "github.com/srv-api/merchant/repositories/dashboard/permission"
)

type PermissionService interface {
	Create(req dto.PermissionRequest) (dto.PermissionResponse, error)
	Get(req dto.RoleUserRequest) (dto.GetPermissionResponse, error)
	Pagination(context echo.Context, req *dto.Pagination) dto.Response
	GetById(req dto.GetPermissionByIdRequest) (*dto.PermissionResponse, error)
	Delete(req dto.DeletePermissionRequest) (dto.DeletePermissionResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (dto.BulkDeleteResponse, error)
	Update(req dto.PermissionUpdateRequest) (dto.PermissionUpdateResponse, error)
}

type permissionService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewPermissionService(Repo r.DomainRepository, jwtS m.JWTService) PermissionService {
	return &permissionService{
		Repo: Repo,
		jwt:  jwtS,
	}
}
