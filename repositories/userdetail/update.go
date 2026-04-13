package userdetail

import (
	dto "github.com/srv-api/detail/dto"
	"github.com/srv-api/detail/entity"
)

func (b *userdetailRepository) Update(req dto.UpdateUserDetailRequest) (dto.UpdateUserDetailResponse, error) {
	tr := dto.GetUserDetailByIdRequest{
		ID: req.ID,
	}

	request := entity.UserDetail{
		ID:           tr.ID,
		UserID:       req.UserID,
		Radius:       req.Radius,
		MinAge:       req.MinAge,
		MaxAge:       req.MaxAge,
		GenderTarget: req.GenderTarget,
		UpdatedBy:    req.UpdatedBy,
	}

	mer, err := b.GetById(tr)
	if err != nil {
		return dto.UpdateUserDetailResponse{}, err
	}

	err = b.DB.Where("ID = ?", req.ID).Updates(entity.UserDetail{
		UserID:       request.UserID,
		Radius:       request.Radius,
		MinAge:       request.MinAge,
		MaxAge:       request.MaxAge,
		GenderTarget: request.GenderTarget,
		UpdatedBy:    request.UpdatedBy,
	}).Error
	if err != nil {
		return dto.UpdateUserDetailResponse{}, err
	}

	response := dto.UpdateUserDetailResponse{
		ID:           mer.ID,
		UserID:       request.UserID,
		Latitude:     request.Latitude,
		Longitude:    request.Longitude,
		Radius:       request.Radius,
		MinAge:       request.MinAge,
		MaxAge:       request.MaxAge,
		GenderTarget: request.GenderTarget,
		UpdatedBy:    request.UpdatedBy,
	}

	return response, nil
}
