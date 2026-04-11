package role

import (
	dto "github.com/srv-api/merchant/dto"
)

func (b *roleService) GetById(req dto.GetRoleByIdRequest) (*dto.RoleResponse, error) {
	transaction, err := b.Repo.GetById(req)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
