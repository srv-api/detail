package roleuser

import (
	dto "github.com/srv-api/detail/dto"
	"github.com/srv-api/detail/entity"
)

func (b *RoleUserRepository) GetById(req dto.GetRoleUserByIdRequest) (*dto.RoleUserResponse, error) {
	tr := entity.RoleUser{
		ID: req.ID,
	}

	if err := b.DB.Where("id = ?", tr.ID).Take(&tr).Error; err != nil {
		return nil, err
	}

	response := &dto.RoleUserResponse{
		RoleID: tr.RoleID,
	}

	return response, nil
}
