package contentsetting

import (
	"encoding/json"

	dto "github.com/srv-api/merchant/dto"
	"github.com/srv-api/merchant/entity"
)

func (r *contentsettingRepository) Get(req dto.ContentSettingRequest) (dto.ContentSettingResponse, error) {
	var data entity.ContentSetting

	// Ambil satu data global (misal untuk landing page utama)
	if err := r.DB.
		Where("deleted_at IS NULL").
		Order("updated_at DESC").
		First(&data).Error; err != nil {
		return dto.ContentSettingResponse{}, err
	}

	// decode JSONB ke struct
	var topHeader []dto.TopHeader
	var buttonHeader []dto.ButtonHeader
	var feature []dto.Feature
	var footer []dto.Footer

	if len(data.TopHeader) > 0 {
		_ = json.Unmarshal(data.TopHeader, &topHeader)
	}
	if len(data.ButtonHeader) > 0 {
		_ = json.Unmarshal(data.ButtonHeader, &buttonHeader)
	}
	if len(data.Feature) > 0 {
		_ = json.Unmarshal(data.Feature, &feature)
	}
	if len(data.Footer) > 0 {
		_ = json.Unmarshal(data.Footer, &footer)
	}

	response := dto.ContentSettingResponse{
		TopHeader:    topHeader,
		ButtonHeader: buttonHeader,
		Feature:      feature,
		Footer:       footer,
	}

	return response, nil
}
