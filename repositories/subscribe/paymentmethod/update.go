package paymentmethod

import (
	"github.com/srv-api/detail/dto"
	"github.com/srv-api/detail/entity"
)

func (b *paymentmethodRepository) Update(req dto.PaymentMethodUpdateRequest) (dto.PaymentMethodUpdateResponse, error) {
	// Menyiapkan struktur update untuk produk
	updatePayment := entity.PaymentMethod{
		PaymentMethod: req.PaymentMethod,
		Status:        req.Status, // Pastikan status boolean diterima dengan benar
		UpdatedBy:     req.UpdatedBy,
		UserID:        req.UserID,
		DetailID:      req.DetailID,
	}

	// Cek apakah produk ada terlebih dahulu
	var existingPayment entity.PaymentMethod
	err := b.DB.Where("id = ?", req.ID).First(&existingPayment).Error
	if err != nil {
		return dto.PaymentMethodUpdateResponse{}, err
	}

	// Update produk dengan nilai yang baru
	err = b.DB.Model(&existingPayment).Updates(updatePayment).Error
	if err != nil {
		return dto.PaymentMethodUpdateResponse{}, err
	}

	// Menyiapkan response setelah pembaruan berhasil
	response := dto.PaymentMethodUpdateResponse{
		PaymentMethod: updatePayment.PaymentMethod,
		Status:        updatePayment.Status,
		UpdatedBy:     updatePayment.UpdatedBy,
		UserID:        updatePayment.UserID,
		DetailID:      updatePayment.DetailID,
	}

	return response, nil
}
