package history

import (
	dto "github.com/srv-api/merchant/dto"
	"github.com/srv-api/merchant/entity"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Get(req *dto.Pagination) (RepositoryResult, int)
	GetById(req dto.GetHistory) (*dto.VAResponse, error)
	FindByOrderID(orderID string) (*entity.Subscribe, error)
	UpdateStatus(orderID, status string) error
}

type historyRepository struct {
	DB *gorm.DB
}

func NewHistoryRepository(DB *gorm.DB) DomainRepository {
	return &historyRepository{
		DB: DB,
	}
}
