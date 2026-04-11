package deleteaccount

import (
	"github.com/srv-api/merchant/dto"
	"github.com/srv-api/merchant/entity"
)

func (b *deleteaccountRepository) Update(req dto.DeleteAccountUpdateRequest) (dto.DeleteAccountUpdateResponse, error) {
	// Menyiapkan struktur update untuk produk
	updateTable := entity.DeleteAccount{
		Email:      req.Email,
		Reason:     req.Reason,
		UpdatedBy:  req.UpdatedBy,
		UserID:     req.UserID,
		MerchantID: req.MerchantID,
	}

	// Cek apakah produk ada terlebih dahulu
	var existingTable entity.DeleteAccount
	err := b.DB.Where("id = ?", req.ID).First(&existingTable).Error
	if err != nil {
		return dto.DeleteAccountUpdateResponse{}, err
	}

	// Update produk dengan nilai yang baru
	err = b.DB.Model(&existingTable).Updates(updateTable).Error
	if err != nil {
		return dto.DeleteAccountUpdateResponse{}, err
	}

	// Menyiapkan response setelah pembaruan berhasil
	response := dto.DeleteAccountUpdateResponse{
		Email:      updateTable.Email,
		Reason:     updateTable.Reason,
		UpdatedBy:  updateTable.UpdatedBy,
		UserID:     updateTable.UserID,
		MerchantID: updateTable.MerchantID,
	}

	return response, nil
}
