package pin

import (
	dto "github.com/srv-api/merchant/dto"
)

func (b *pinService) Delete(req dto.DeletePinRequest) (dto.DeletePinResponse, error) {
	transactionBody := dto.DeleteRequest{
		ID:        req.ID,
		DeletedBy: req.DeletedBy,
	}

	_, err := b.Repo.Delete(req)
	if err != nil {
		return dto.DeletePinResponse{}, err
	}

	response := dto.DeletePinResponse{
		ID:        transactionBody.ID,
		DeletedBy: transactionBody.DeletedBy,
	}

	return response, nil
}
