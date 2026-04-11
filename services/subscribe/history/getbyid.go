package history

import (
	dto "github.com/srv-api/merchant/dto"
)

func (b *historyService) GetById(req dto.GetHistory) (*dto.VAResponse, error) {
	transaction, err := b.Repo.GetById(req)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
