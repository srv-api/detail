package pin

import (
	dto "github.com/srv-api/merchant/dto"
)

func (b *pinService) GetById(req dto.GetByIdPinRequest) (*dto.PinResponse, error) {
	transaction, err := b.Repo.GetById(req)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
