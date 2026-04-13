package paymentmethod

import (
	dto "github.com/srv-api/merchant/dto"
	"github.com/srv-api/merchant/entity"
)

func (r *paymentmethodRepository) Create(req dto.PaymentMethodRequest) (dto.PaymentMethodResponse, error) {

	create := entity.PaymentMethod{
		ID:            req.ID,
		PaymentMethod: req.PaymentMethod,
		Status:        req.Status,
		UserID:        req.UserID,
		MerchantID:    req.MerchantID,
		CreatedBy:     req.CreatedBy,
		Category:      req.Category,
	}

	if err := r.DB.Create(&create).Error; err != nil {
		return dto.PaymentMethodResponse{}, err
	}

	response := dto.PaymentMethodResponse{
		ID:            create.ID,
		PaymentMethod: create.PaymentMethod,
		Status:        create.Status,
		UserID:        create.UserID,
		MerchantID:    create.MerchantID,
		CreatedBy:     create.CreatedBy,
		Category:      req.Category,
	}

	return response, nil
}

func (r *paymentmethodRepository) SaveImage(img entity.UploadedPayment) error {
	return r.DB.Create(&img).Error
}
