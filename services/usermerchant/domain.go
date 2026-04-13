package user

import (
	"github.com/labstack/echo/v4"
	dto "github.com/srv-api/detail/dto"
	m "github.com/srv-api/middlewares/middlewares"

	r "github.com/srv-api/detail/repositories/usermerchant"
)

type UserDetailService interface {
	Create(req dto.UserDetailRequest) (dto.UserDetailResponse, error)
	Get(context echo.Context, req *dto.Pagination) dto.Response
	GetById(req dto.GetByIdRequest) (*dto.UserDetailByIdResponse, error)
	Delete(req dto.DeleteRequest) (dto.DeleteResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (dto.BulkDeleteResponse, error)
	Update(req dto.UserDetailUpdateRequest) (dto.UserDetailUpdateResponse, error)
}

type userService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewUserDetailService(Repo r.DomainRepository, jwtS m.JWTService) UserDetailService {
	return &userService{
		Repo: Repo,
		jwt:  jwtS,
	}
}
