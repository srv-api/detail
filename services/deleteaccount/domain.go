package deleteaccount

import (
	"github.com/labstack/echo/v4"
	dto "github.com/srv-api/merchant/dto"
	m "github.com/srv-api/middlewares/middlewares"

	r "github.com/srv-api/merchant/repositories/deleteaccount"
)

type DeleteAccountService interface {
	Create(req dto.DeleteAccountRequest) (dto.DeleteAccountResponse, error)
	Get(context echo.Context, req *dto.Pagination) dto.Response
	GetById(req dto.GetDeleteAccountByIdRequest) (*dto.DeleteAccountResponse, error)
	Delete(req dto.DeleteDeleteAccountRequest) (dto.DeleteDeleteAccountResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (dto.BulkDeleteResponse, error)
	Update(req dto.DeleteAccountUpdateRequest) (dto.DeleteAccountUpdateResponse, error)
}

type deleteaccountService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewDeleteAccountService(Repo r.DomainRepository, jwtS m.JWTService) DeleteAccountService {
	return &deleteaccountService{
		Repo: Repo,
		jwt:  jwtS,
	}
}
