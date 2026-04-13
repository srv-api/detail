package roleuserpermission

import "github.com/srv-api/merchant/dto"

func (b *roleuserpermissionService) Update(req dto.RoleUserPermissionUpdateRequest) (dto.RoleUserPermissionUpdateResponse, error) {
	request := dto.RoleUserPermissionUpdateRequest{
		RoleUserID: req.RoleUserID,
		UpdatedBy:  req.UpdatedBy,
		UserID:     req.UserID,
	}

	role, err := b.Repo.Update(req)
	if err != nil {
		return role, err
	}

	response := dto.RoleUserPermissionUpdateResponse{
		RoleUserID: req.RoleUserID,
		UpdatedBy:  request.UpdatedBy,
		UserID:     request.UserID,
	}

	return response, nil
}
