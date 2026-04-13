package deleteaccount

import (
	dto "github.com/srv-api/detail/dto"
)

func (b *deleteaccountService) Delete(req dto.DeleteDeleteAccountRequest) (dto.DeleteDeleteAccountResponse, error) {
	transactionBody := dto.DeleteDeleteAccountRequest{
		ID:        req.ID,
		DeletedBy: req.DeletedBy,
	}

	_, err := b.Repo.Delete(req)
	if err != nil {
		return dto.DeleteDeleteAccountResponse{}, err
	}

	response := dto.DeleteDeleteAccountResponse{
		ID:        transactionBody.ID,
		DeletedBy: transactionBody.DeletedBy,
	}

	return response, nil
}
