package pin

import (
	"github.com/srv-api/merchant/dto"
	"github.com/srv-api/merchant/entity"
)

func (b *pinRepository) Update(req dto.PinUpdateRequest) (dto.PinUpdateResponse, error) {
	// Menyiapkan struktur update untuk produk
	updatePin := entity.Pin{
		Pin:        req.Pin,
		UpdatedBy:  req.UpdatedBy,
		UserID:     req.UserID,
		MerchantID: req.MerchantID,
	}

	var existingProduct entity.Pin
	err := b.DB.Where("id = ?", req.ID).First(&existingProduct).Error
	if err != nil {
		return dto.PinUpdateResponse{}, err
	}

	// Update produk dengan nilai yang baru
	err = b.DB.Model(&existingProduct).Updates(updatePin).Error
	if err != nil {
		return dto.PinUpdateResponse{}, err
	}

	// Menyiapkan response setelah pembaruan berhasil
	response := dto.PinUpdateResponse{
		Pin:        updatePin.Pin,
		UpdatedBy:  updatePin.UpdatedBy,
		UserID:     updatePin.UserID,
		MerchantID: updatePin.MerchantID,
	}

	return response, nil
}
