package user

import (
	"github.com/labstack/echo/v4"
	dto "github.com/srv-api/merchant/dto"
	m "github.com/srv-api/middlewares/middlewares"

	r "github.com/srv-api/merchant/repositories/usermerchant"
)

type UserMerchantService interface {
	Create(req dto.UserMerchantRequest) (dto.UserMerchantResponse, error)
	Get(context echo.Context, req *dto.Pagination) dto.Response
	GetById(req dto.GetByIdRequest) (*dto.UserMerchantByIdResponse, error)
	Delete(req dto.DeleteRequest) (dto.DeleteResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (dto.BulkDeleteResponse, error)
	Update(req dto.UserMerchantUpdateRequest) (dto.UserMerchantUpdateResponse, error)
}

type userService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewUserMerchantService(Repo r.DomainRepository, jwtS m.JWTService) UserMerchantService {
	return &userService{
		Repo: Repo,
		jwt:  jwtS,
	}
}
