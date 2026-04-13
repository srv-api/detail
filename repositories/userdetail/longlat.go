package userdetail

import (
	dto "github.com/srv-api/detail/dto"
	"github.com/srv-api/detail/entity"
)

func (b *userdetailRepository) LongLat(req dto.UpdateUserDetailRequest) (dto.UpdateUserDetailResponse, error) {
	tr := dto.GetUserDetailByIdRequest{
		ID: req.ID,
	}

	request := entity.UserDetail{
		ID:        tr.ID,
		UserID:    req.UserID,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
		UpdatedBy: req.UpdatedBy,
	}

	mer, err := b.GetById(tr)
	if err != nil {
		return dto.UpdateUserDetailResponse{}, err
	}

	err = b.DB.Where("ID = ?", req.ID).Updates(entity.UserDetail{
		UserID:    request.UserID,
		Latitude:  request.Latitude,
		Longitude: request.Longitude,
		UpdatedBy: request.UpdatedBy,
	}).Error
	if err != nil {
		return dto.UpdateUserDetailResponse{}, err
	}

	response := dto.UpdateUserDetailResponse{
		ID:        mer.ID,
		UserID:    request.UserID,
		Latitude:  request.Latitude,
		Longitude: request.Longitude,
		UpdatedBy: request.UpdatedBy,
	}

	return response, nil
}
