package paymentmethod

import "github.com/srv-api/merchant/dto"

func (b *paymentmethodService) Update(req dto.PaymentMethodUpdateRequest) (dto.PaymentMethodUpdateResponse, error) {
	request := dto.PaymentMethodUpdateRequest{
		PaymentMethod: req.PaymentMethod,
		Status:        req.Status,
		UpdatedBy:     req.UpdatedBy,

		UserID: req.UserID,
	}

	paymentmethod, err := b.Repo.Update(req)
	if err != nil {
		return paymentmethod, err
	}

	response := dto.PaymentMethodUpdateResponse{
		PaymentMethod: request.PaymentMethod,
		Status:        request.Status,
		UpdatedBy:     request.UpdatedBy,
		UserID:        request.UserID,
	}

	return response, nil
}
