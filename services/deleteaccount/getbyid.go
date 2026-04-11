package deleteaccount

import (
	dto "github.com/srv-api/merchant/dto"
)

func (b *deleteaccountService) GetById(req dto.GetDeleteAccountByIdRequest) (*dto.DeleteAccountResponse, error) {
	transaction, err := b.Repo.GetById(req)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
