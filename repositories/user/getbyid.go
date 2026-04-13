package user

import (
	"github.com/srv-api/auth/entity"
	dto "github.com/srv-api/merchant/dto"
)

func (b *userRepository) GetById(req dto.GetByIdRequest) (*dto.UserResponse, error) {
	tr := entity.AccessDoor{
		ID: req.ID,
	}

	if err := b.DB.Where("id = ?", tr.ID).Take(&tr).Error; err != nil {
		return nil, err
	}

	response := &dto.UserResponse{
		FullName: tr.FullName,
	}

	return response, nil
}
