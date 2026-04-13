package paymentmethod

import (
	dto "github.com/srv-api/merchant/dto"
)

func (b *paymentmethodService) GetById(req dto.GetByIdPaymentRequest) (dto.PaymentMethodResponse, error) {
	pay, err := b.Repo.GetById(req)
	if err != nil {
		return dto.PaymentMethodResponse{}, err
	}
	res := dto.PaymentMethodResponse{
		PaymentMethod: pay.PaymentMethod,
	}

	return res, nil
}
