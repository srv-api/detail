package purchase

import (
	"github.com/srv-api/detail/entity"
	"gorm.io/gorm"
)

type DomainRepository interface {
	Create(purchase *entity.PurchaseHistory) error
	FindByToken(token string) (*entity.PurchaseHistory, error)
	FindByUserID(userID string) ([]entity.PurchaseHistory, error)
	UpdateStatus(id string, status string) error
	ExistsByToken(token string) (bool, error)
}

type purchaseRepository struct {
	DB *gorm.DB
}

func NewUserDetailRepository(DB *gorm.DB) DomainRepository {
	return &purchaseRepository{
		DB: DB,
	}
}
