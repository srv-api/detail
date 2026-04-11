package permission

import (
	dto "github.com/srv-api/merchant/dto"
)

func (b *permissionService) GetById(req dto.GetPermissionByIdRequest) (*dto.PermissionResponse, error) {
	transaction, err := b.Repo.GetById(req)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
