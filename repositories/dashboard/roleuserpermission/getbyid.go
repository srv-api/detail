package roleuserpermission

import (
	dto "github.com/srv-api/detail/dto"
	"github.com/srv-api/detail/entity"
)

func (b *RoleUserPermissionRepository) GetById(req dto.GetRoleUserPermissionByIdRequest) (*dto.RoleUserPermissionResponse, error) {
	tr := entity.RoleUserPermission{
		ID: req.ID,
	}

	if err := b.DB.Where("id = ?", tr.ID).Take(&tr).Error; err != nil {
		return nil, err
	}

	response := &dto.RoleUserPermissionResponse{
		ID: tr.ID,
	}

	return response, nil
}
