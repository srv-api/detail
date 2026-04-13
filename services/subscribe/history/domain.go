package history

import (
	"github.com/labstack/echo/v4"
	dto "github.com/srv-api/merchant/dto"
	m "github.com/srv-api/middlewares/middlewares"

	r "github.com/srv-api/merchant/repositories/subscribe/history"
)

type HistoryService interface {
	Get(context echo.Context, req *dto.Pagination) dto.Response
	ExpireTransaction(orderID string) error
	GetById(req dto.GetHistory) (*dto.VAResponse, error)
}

type historyService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewHistoryService(Repo r.DomainRepository, jwtS m.JWTService) HistoryService {
	return &historyService{
		Repo: Repo,
		jwt:  jwtS,
	}
}
