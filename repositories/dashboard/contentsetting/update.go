package contentsetting

import (
	"encoding/json"
	"time"

	dto "github.com/srv-api/merchant/dto"
)

func (r *contentsettingRepository) Update(req dto.UpdateContentSettingRequest) error {
	setting, err := r.GetById(req.ID)
	if err != nil {
		return err
	}

	topHeaderJSON, _ := json.Marshal(req.TopHeader)
	buttonHeaderJSON, _ := json.Marshal(req.ButtonHeader)
	featureJSON, _ := json.Marshal(req.Feature)
	footerJSON, _ := json.Marshal(req.Footer)

	setting.TopHeader = topHeaderJSON
	setting.ButtonHeader = buttonHeaderJSON
	setting.Feature = featureJSON
	setting.Footer = footerJSON
	setting.UpdatedBy = req.UpdatedBy
	setting.UpdatedAt = time.Now()

	return r.DB.Save(setting).Error
}
