package userdetail

import (
	dto "github.com/srv-api/detail/dto"
	"github.com/srv-api/detail/entity"
)

func (b *userdetailRepository) GetById(req dto.GetUserDetailByIdRequest) (*dto.UserDetailRequest, error) {
	tr := entity.UserDetail{
		ID: req.ID,
	}

	if err := b.DB.Where("id = ?", tr.ID).Take(&tr).Error; err != nil {
		return nil, err
	}

	response := &dto.UserDetailRequest{
		ID:           tr.ID,
		UserID:       tr.UserID,
		Latitude:     tr.Latitude,
		Longitude:    tr.Longitude,
		Radius:       tr.Radius,
		MinAge:       tr.MinAge,
		MaxAge:       tr.MaxAge,
		GenderTarget: tr.GenderTarget,
	}

	return response, nil
}
