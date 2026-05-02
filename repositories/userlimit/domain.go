package userlimit

import (
	"time"

	"github.com/srv-api/detail/entity"
	"gorm.io/gorm"
)

type DomainRepository interface {
	FindByUserID(userID string) (*entity.UserLimit, error)
	Update(userLimit *entity.UserLimit) error
	IncrementLimits(userID string, swipeIncrement, superLikeIncrement int) error
	UpdatePremiumStatus(userID string, isPremium bool, expiry *time.Time) error
	CreateOrUpdate(userLimit *entity.UserLimit) error
}

type userlimitRepository struct {
	DB *gorm.DB
}

func NewUserDetailRepository(DB *gorm.DB) DomainRepository {
	return &userlimitRepository{
		DB: DB,
	}
}
