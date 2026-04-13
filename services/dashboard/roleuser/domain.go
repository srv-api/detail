package roleuser

import (
	"github.com/labstack/echo/v4"
	dto "github.com/srv-api/merchant/dto"
	m "github.com/srv-api/middlewares/middlewares"

	r "github.com/srv-api/merchant/repositories/dashboard/roleuser"
)

type RoleUserService interface {
	Create(req dto.RoleUserRequest) (dto.RoleUserResponse, error)
	Get(req dto.RoleUserRequest) (dto.GetRoleUserResponse, error)
	Pagination(context echo.Context, req *dto.Pagination) dto.Response
	GetById(req dto.GetRoleUserByIdRequest) (*dto.RoleUserResponse, error)
	Delete(req dto.DeleteRoleUserRequest) (dto.DeleteRoleUserResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (dto.BulkDeleteResponse, error)
	Update(req dto.RoleUserUpdateRequest) (dto.RoleUserUpdateResponse, error)
}

type roleuserService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewRoleUserService(Repo r.DomainRepository, jwtS m.JWTService) RoleUserService {
	return &roleuserService{
		Repo: Repo,
		jwt:  jwtS,
	}
}
