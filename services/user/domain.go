package user

import (
	"github.com/labstack/echo/v4"
	dto "github.com/srv-api/merchant/dto"
	m "github.com/srv-api/middlewares/middlewares"

	r "github.com/srv-api/merchant/repositories/user"
)

type UserService interface {
	Create(req dto.UserRequest) (dto.UserResponse, error)
	Get(context echo.Context, req *dto.Pagination) dto.Response
	GetById(req dto.GetByIdRequest) (*dto.UserResponse, error)
	Delete(req dto.DeleteRequest) (dto.DeleteResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (dto.BulkDeleteResponse, error)
	Update(req dto.UserUpdateRequest) (dto.UserUpdateResponse, error)
}

type userService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewUserService(Repo r.DomainRepository, jwtS m.JWTService) UserService {
	return &userService{
		Repo: Repo,
		jwt:  jwtS,
	}
}
