package paymentmethod

import (
	dto "github.com/srv-api/merchant/dto"
	"github.com/srv-api/merchant/entity"
)

func (r *paymentmethodRepository) Get(req dto.PaymentMethodRequest) ([]dto.PaymentMethodResponse, error) {
	var paymentmethods []entity.PaymentMethod
	if err := r.DB.Preload("Image").Find(&paymentmethods).Error; err != nil {
		return nil, err
	}

	var result []dto.PaymentMethodResponse
	for _, pm := range paymentmethods {
		var image *dto.UploadedPayment
		if pm.Image != nil {
			image = &dto.UploadedPayment{
				ID:        pm.Image.ID,
				PaymentID: pm.Image.PaymentID,
				FileName:  pm.Image.FileName,
				FilePath:  pm.Image.FilePath,
			}
		}

		result = append(result, dto.PaymentMethodResponse{
			ID:            pm.ID,
			UserID:        pm.UserID,
			MerchantID:    pm.MerchantID,
			PaymentMethod: pm.PaymentMethod,
			Status:        pm.Status,
			CreatedBy:     pm.CreatedBy,
			UpdatedBy:     pm.UpdatedBy,
			DeletedBy:     pm.DeletedBy,
			CreatedAt:     pm.CreatedAt,
			Image:         image,
		})
	}

	return result, nil
}
