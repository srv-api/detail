package role

import (
	"github.com/srv-api/merchant/dto"
	"github.com/srv-api/merchant/entity"
)

func (b *RoleRepository) Update(req dto.RoleUpdateRequest) (dto.RoleUpdateResponse, error) {
	// Menyiapkan struktur update untuk produk
	updateRole := entity.Role{
		Role: req.Role,
	}

	// Cek apakah produk ada terlebih dahulu
	var existingRole entity.Role
	err := b.DB.Where("id = ?", req.ID).First(&existingRole).Error
	if err != nil {
		return dto.RoleUpdateResponse{}, err
	}

	// Update produk dengan nilai yang baru
	err = b.DB.Model(&existingRole).Updates(updateRole).Error
	if err != nil {
		return dto.RoleUpdateResponse{}, err
	}

	// Menyiapkan response setelah pembaruan berhasil
	response := dto.RoleUpdateResponse{
		Role: updateRole.Role,
	}

	return response, nil
}
