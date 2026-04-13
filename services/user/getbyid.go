package user

import (
	dto "github.com/srv-api/merchant/dto"
)

func (b *userService) GetById(req dto.GetByIdRequest) (*dto.UserResponse, error) {
	transaction, err := b.Repo.GetById(req)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
