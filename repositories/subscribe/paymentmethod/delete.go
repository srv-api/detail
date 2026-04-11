package paymentmethod

import (
	dto "github.com/srv-api/merchant/dto"
	"github.com/srv-api/merchant/entity"
)

func (b *paymentmethodRepository) Delete(req dto.DeletePaymentRequest) (dto.DeletePaymentResponse, error) {
	tr := dto.GetByIdPaymentRequest{
		ID: req.ID,
	}

	_, err := b.GetById(tr)
	if err != nil {
		return dto.DeletePaymentResponse{}, err
	}

	// Use GORM BeforeDelete hook to set DeletedBy
	if err := b.DB.Where("id = ?", req.ID).Delete(&entity.PaymentMethod{}).Error; err != nil {
		return dto.DeletePaymentResponse{}, err
	}

	response := dto.DeletePaymentResponse{
		DeletedBy: req.DeletedBy,
	}

	return response, nil
}
