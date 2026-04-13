package roleuserpermission

import (
	"github.com/srv-api/merchant/dto"
	"github.com/srv-api/merchant/entity"
)

func (b *RoleUserPermissionRepository) Update(req dto.RoleUserPermissionUpdateRequest) (dto.RoleUserPermissionUpdateResponse, error) {
	// Menyiapkan struktur update untuk produk
	updateRole := entity.RoleUserPermission{
		ID: req.ID,
	}

	// Cek apakah produk ada terlebih dahulu
	var existingRole entity.RoleUserPermission
	err := b.DB.Where("id = ?", req.ID).First(&existingRole).Error
	if err != nil {
		return dto.RoleUserPermissionUpdateResponse{}, err
	}

	// Update produk dengan nilai yang baru
	err = b.DB.Model(&existingRole).Updates(updateRole).Error
	if err != nil {
		return dto.RoleUserPermissionUpdateResponse{}, err
	}

	// Menyiapkan response setelah pembaruan berhasil
	response := dto.RoleUserPermissionUpdateResponse{
		ID: updateRole.ID,
	}

	return response, nil
}
