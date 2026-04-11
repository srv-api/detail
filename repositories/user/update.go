package user

import (
	"github.com/srv-api/auth/entity"
	"github.com/srv-api/merchant/dto"
)

func (b *userRepository) Update(req dto.UserUpdateRequest) (dto.UserUpdateResponse, error) {
	// Menyiapkan struktur update untuk produk
	updateUser := entity.AccessDoor{
		FullName: req.FullName,
	}

	// Cek apakah produk ada terlebih dahulu
	var existingUser entity.AccessDoor
	err := b.DB.Where("id = ?", req.ID).First(&existingUser).Error
	if err != nil {
		return dto.UserUpdateResponse{}, err
	}

	// Update produk dengan nilai yang baru
	err = b.DB.Model(&existingUser).Updates(updateUser).Error
	if err != nil {
		return dto.UserUpdateResponse{}, err
	}

	// Menyiapkan response setelah pembaruan berhasil
	response := dto.UserUpdateResponse{
		FullName: updateUser.FullName,
	}

	return response, nil
}
