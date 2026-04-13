package roleuserpermission

import (
	dto "github.com/srv-api/merchant/dto"
)

func (b *roleuserpermissionService) Delete(req dto.DeleteRoleUserPermissionRequest) (dto.DeleteRoleUserPermissionResponse, error) {
	transactionBody := dto.DeleteRoleUserPermissionRequest{
		ID:        req.ID,
		DeletedBy: req.DeletedBy,
	}

	_, err := b.Repo.Delete(req)
	if err != nil {
		return dto.DeleteRoleUserPermissionResponse{}, err
	}

	response := dto.DeleteRoleUserPermissionResponse{
		ID:        transactionBody.ID,
		DeletedBy: transactionBody.DeletedBy,
	}

	return response, nil
}
