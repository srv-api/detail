package user

import (
	"github.com/srv-api/auth/entity"
	"github.com/srv-api/detail/dto"
)

func (b *userRepository) Update(req dto.UserDetailUpdateRequest) (dto.UserDetailUpdateResponse, error) {
	// Menyiapkan struktur update untuk produk
	updateUserDetail := entity.AccessDoor{
		FullName:     req.FullName,
		UpdatedBy:    req.UpdatedBy,
		Email:        req.Email,
		Whatsapp:     req.Whatsapp,
		Password:     req.Password,
		AccessRoleID: req.AccessRoleID,
		Verified: entity.UserVerified{
			Verified:      req.Verified.Verified,
			StatusAccount: req.Verified.StatusAccount,
		},
	}

	// Cek apakah produk ada terlebih dahulu
	var existingUserDetail entity.AccessDoor
	err := b.DB.Where("id = ?", req.ID).First(&existingUserDetail).Error
	if err != nil {
		return dto.UserDetailUpdateResponse{}, err
	}

	// Update produk dengan nilai yang baru
	err = b.DB.Model(&existingUserDetail).Updates(updateUserDetail).Error
	if err != nil {
		return dto.UserDetailUpdateResponse{}, err
	}

	// Menyiapkan response setelah pembaruan berhasil
	response := dto.UserDetailUpdateResponse{
		FullName:     updateUserDetail.FullName,
		UpdatedBy:    updateUserDetail.UpdatedBy,
		Email:        updateUserDetail.Email,
		Whatsapp:     updateUserDetail.Whatsapp,
		Password:     updateUserDetail.Password,
		AccessRoleID: updateUserDetail.AccessRoleID,
		Verified: dto.UserDetailVerifiedByID{
			Verified:      updateUserDetail.Verified.Verified,
			StatusAccount: updateUserDetail.Verified.StatusAccount,
		},
	}

	return response, nil
}
