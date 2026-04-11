package paymentmethod

import (
	dto "github.com/srv-api/merchant/dto"
	"github.com/srv-api/merchant/entity"
)

func (b *paymentmethodRepository) GetById(req dto.GetByIdPaymentRequest) (*dto.PaymentMethodResponse, error) {
	tr := entity.PaymentMethod{
		ID: req.ID,
	}

	if err := b.DB.Where("id = ?", tr.ID).Take(&tr).Error; err != nil {
		return nil, err
	}

	response := &dto.PaymentMethodResponse{
		PaymentMethod: tr.PaymentMethod,
	}

	return response, nil
}
