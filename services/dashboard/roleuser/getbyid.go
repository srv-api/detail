package roleuser

import (
	dto "github.com/srv-api/merchant/dto"
)

func (b *roleuserService) GetById(req dto.GetRoleUserByIdRequest) (*dto.RoleUserResponse, error) {
	transaction, err := b.Repo.GetById(req)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
