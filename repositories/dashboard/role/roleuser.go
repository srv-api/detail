package role

import (
	dto "github.com/srv-api/merchant/dto"
	"github.com/srv-api/merchant/entity"
)

func (r *RoleRepository) RoleUser(req dto.GetRoleRequest) (dto.GetRoleResponse, error) {
	var roles []entity.Role

	if err := r.DB.
		Where("role <> ?", "God Cashpay").
		Find(&roles).Error; err != nil {
		return dto.GetRoleResponse{}, err
	}

	roler := make([]dto.RoleResponse, len(roles))
	for i, role := range roles {
		roler[i] = dto.RoleResponse{
			ID:   role.ID,
			Role: role.Role,
		}
	}

	return dto.GetRoleResponse{Role: roler}, nil
}
