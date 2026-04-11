package permission

import (
	dto "github.com/srv-api/merchant/dto"
	"github.com/srv-api/merchant/entity"
)

func (b *PermissionRepository) GetById(req dto.GetPermissionByIdRequest) (*dto.PermissionResponse, error) {
	tr := entity.Permission{
		ID: req.ID,
	}

	if err := b.DB.Where("id = ?", tr.ID).Take(&tr).Error; err != nil {
		return nil, err
	}

	response := &dto.PermissionResponse{
		Label: tr.Label,
	}

	return response, nil
}
