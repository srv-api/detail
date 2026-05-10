package boost

import (
	dto "github.com/srv-api/detail/dto"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Create(req dto.BoostRequest) (dto.BoostResponse, error)
}

type boostRepository struct {
	DB *gorm.DB
}

func NewBoostRepository(DB *gorm.DB) DomainRepository {
	return &boostRepository{
		DB: DB,
	}
}
