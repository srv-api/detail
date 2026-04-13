package roleuserpermission

import (
	dto "github.com/srv-api/merchant/dto"
)

func (b *roleuserpermissionService) GetById(req dto.GetRoleUserPermissionByIdRequest) (*dto.RoleUserPermissionResponse, error) {
	transaction, err := b.Repo.GetById(req)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
