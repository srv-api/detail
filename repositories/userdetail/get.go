package userdetail

import (
	dto "github.com/srv-api/detail/dto"
	"github.com/srv-api/detail/entity"
)

func (r *userdetailRepository) Get(req dto.UserDetailRequest) (dto.UserDetailResponse, error) {

	var data entity.UserDetail

	if err := r.DB.Where("user_id = ?", req.UserID).Find(&data).Error; err != nil {
		return dto.UserDetailResponse{}, err
	}

	response := dto.UserDetailResponse{
		ID:           data.ID,
		UserID:       data.UserID,
		Latitude:     data.Latitude,
		Longitude:    data.Longitude,
		Radius:       data.Radius,
		MinAge:       data.MinAge,
		MaxAge:       data.MaxAge,
		GenderTarget: data.GenderTarget,
		UpdatedAt:    data.UpdatedAt,
	}

	return response, nil
}
