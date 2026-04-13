package paymentmethod

import (
	dto "github.com/srv-api/detail/dto"
	"github.com/srv-api/detail/entity"
)

func (r *paymentmethodRepository) Create(req dto.PaymentMethodRequest) (dto.PaymentMethodResponse, error) {

	create := entity.PaymentMethod{
		ID:            req.ID,
		PaymentMethod: req.PaymentMethod,
		Status:        req.Status,
		UserID:        req.UserID,
		DetailID:      req.DetailID,
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
		DetailID:      create.DetailID,
		CreatedBy:     create.CreatedBy,
		Category:      req.Category,
	}

	return response, nil
}

func (r *paymentmethodRepository) SaveImage(img entity.UploadedPayment) error {
	return r.DB.Create(&img).Error
}
