package permission

import (
	"github.com/srv-api/merchant/dto"
	"github.com/srv-api/merchant/entity"
)

func (b *PermissionRepository) Update(req dto.PermissionUpdateRequest) (dto.PermissionUpdateResponse, error) {
	// Menyiapkan struktur update untuk produk
	updatePermission := entity.Permission{
		Label:      req.Label,
		Icon:       req.Icon, // Pastikan status boolean diterima dengan benar
		UpdatedBy:  req.UpdatedBy,
		UserID:     req.UserID,
		To:         req.To,
		MerchantID: req.MerchantID,
	}

	// Cek apakah produk ada terlebih dahulu
	var existingPermission entity.Permission
	err := b.DB.Where("id = ?", req.ID).First(&existingPermission).Error
	if err != nil {
		return dto.PermissionUpdateResponse{}, err
	}

	// Update produk dengan nilai yang baru
	err = b.DB.Model(&existingPermission).Updates(updatePermission).Error
	if err != nil {
		return dto.PermissionUpdateResponse{}, err
	}

	// Menyiapkan response setelah pembaruan berhasil
	response := dto.PermissionUpdateResponse{
		Label:      updatePermission.Label,
		Icon:       updatePermission.Icon,
		UpdatedBy:  updatePermission.UpdatedBy,
		UserID:     updatePermission.UserID,
		MerchantID: updatePermission.MerchantID,
		To:         updatePermission.To,
	}

	return response, nil
}
