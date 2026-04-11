package paymentmethod

import (
	dto "github.com/srv-api/merchant/dto"
)

func (b *paymentmethodService) Delete(req dto.DeletePaymentRequest) (dto.DeletePaymentResponse, error) {
	transactionBody := dto.DeletePaymentRequest{
		ID:        req.ID,
		DeletedBy: req.DeletedBy,
	}

	_, err := b.Repo.Delete(req)
	if err != nil {
		return dto.DeletePaymentResponse{}, err
	}

	response := dto.DeletePaymentResponse{
		ID:        transactionBody.ID,
		DeletedBy: transactionBody.DeletedBy,
	}

	return response, nil
}
