package pin

import (
	"github.com/labstack/echo/v4"
	dto "github.com/srv-api/merchant/dto"
	m "github.com/srv-api/middlewares/middlewares"

	r "github.com/srv-api/merchant/repositories/pin"
)

type PinService interface {
	Create(req dto.PinRequest) (dto.PinResponse, error)
	BulkDelete(req dto.BulkDeletePinRequest) (dto.BulkDeletePinResponse, error)
	Delete(req dto.DeletePinRequest) (dto.DeletePinResponse, error)
	Get(context echo.Context, req *dto.Pagination) dto.Response
	GetById(req dto.GetByIdPinRequest) (*dto.PinResponse, error)
	Update(req dto.PinUpdateRequest) (dto.PinUpdateResponse, error)
	VerifyPIN(req dto.VerifyPinRequest) (*dto.VerifyPinResponse, error)
	GetPinStatus(req dto.PinRequest) (dto.PinStatusResponse, error)
}

type pinService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewPinService(Repo r.DomainRepository, jwtS m.JWTService) PinService {
	return &pinService{
		Repo: Repo,
		jwt:  jwtS,
	}
}
