package purchase

import (
	dto "github.com/srv-api/detail/dto"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Create(req dto.PremiumPurchaseRequest) (dto.PremiumPurchaseResponse, error)
}

type purchaseRepository struct {
	DB *gorm.DB
}

func NewPinRepository(DB *gorm.DB) DomainRepository {
	return &purchaseRepository{
		DB: DB,
	}
}
