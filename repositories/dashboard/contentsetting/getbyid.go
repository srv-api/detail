package contentsetting

import (
	"github.com/srv-api/merchant/entity"
)

func (r *contentsettingRepository) GetById(id string) (*entity.ContentSetting, error) {
	var setting entity.ContentSetting
	if err := r.DB.First(&setting, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &setting, nil
}
