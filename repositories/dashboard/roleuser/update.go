package roleuser

import (
	"github.com/srv-api/merchant/dto"
	"github.com/srv-api/merchant/entity"
)

func (b *RoleUserRepository) Update(req dto.RoleUserUpdateRequest) (dto.RoleUserUpdateResponse, error) {
	// Menyiapkan struktur update untuk produk
	updateRole := entity.RoleUser{
		RoleID: req.RoleID,
	}

	// Cek apakah produk ada terlebih dahulu
	var existingRole entity.RoleUser
	err := b.DB.Where("id = ?", req.ID).First(&existingRole).Error
	if err != nil {
		return dto.RoleUserUpdateResponse{}, err
	}

	// Update produk dengan nilai yang baru
	err = b.DB.Model(&existingRole).Updates(updateRole).Error
	if err != nil {
		return dto.RoleUserUpdateResponse{}, err
	}

	// Menyiapkan response setelah pembaruan berhasil
	response := dto.RoleUserUpdateResponse{
		RoleID: updateRole.RoleID,
	}

	return response, nil
}
