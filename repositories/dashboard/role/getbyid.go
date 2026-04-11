package role

import (
	dto "github.com/srv-api/merchant/dto"
	"github.com/srv-api/merchant/entity"
)

func (b *RoleRepository) GetById(req dto.GetRoleByIdRequest) (*dto.RoleResponse, error) {
	tr := entity.Role{
		ID: req.ID,
	}

	if err := b.DB.Where("id = ?", tr.ID).Take(&tr).Error; err != nil {
		return nil, err
	}

	response := &dto.RoleResponse{
		Role: tr.Role,
	}

	return response, nil
}
