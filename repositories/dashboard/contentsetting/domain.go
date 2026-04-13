package contentsetting

import (
	dto "github.com/srv-api/merchant/dto"
	"github.com/srv-api/merchant/entity"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Get(req dto.ContentSettingRequest) (dto.ContentSettingResponse, error)
	GetById(id string) (*entity.ContentSetting, error)
	Update(req dto.UpdateContentSettingRequest) error
}

type contentsettingRepository struct {
	DB *gorm.DB
}

func NewContentSettingRepository(DB *gorm.DB) DomainRepository {
	return &contentsettingRepository{
		DB: DB,
	}
}
