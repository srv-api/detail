package roleuserpermission

import (
	dto "github.com/srv-api/merchant/dto"
	"github.com/srv-api/merchant/entity"
)

func (r *RoleUserPermissionRepository) Create(req dto.RoleUserPermissionRequest) (dto.RoleUserPermissionResponse, error) {

	// Create the new Role entry
	create := entity.RoleUserPermission{
		ID:           req.ID,
		RoleUserID:   req.RoleUserID,
		PermissionID: req.PermissionID,
	}

	// Save the new Role to the database
	if err := r.DB.Save(&create).Error; err != nil {
		return dto.RoleUserPermissionResponse{}, err
	}

	// Build the response for the created Role
	response := dto.RoleUserPermissionResponse{
		ID:           create.ID,
		RoleUserID:   create.RoleUserID,
		PermissionID: create.PermissionID,
	}

	return response, nil
}
