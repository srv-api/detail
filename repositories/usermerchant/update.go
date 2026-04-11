package user

import (
	"github.com/srv-api/auth/entity"
	"github.com/srv-api/merchant/dto"
)

func (b *userRepository) Update(req dto.UserMerchantUpdateRequest) (dto.UserMerchantUpdateResponse, error) {
	// Menyiapkan struktur update untuk produk
	updateUserMerchant := entity.AccessDoor{
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
	var existingUserMerchant entity.AccessDoor
	err := b.DB.Where("id = ?", req.ID).First(&existingUserMerchant).Error
	if err != nil {
		return dto.UserMerchantUpdateResponse{}, err
	}

	// Update produk dengan nilai yang baru
	err = b.DB.Model(&existingUserMerchant).Updates(updateUserMerchant).Error
	if err != nil {
		return dto.UserMerchantUpdateResponse{}, err
	}

	// Menyiapkan response setelah pembaruan berhasil
	response := dto.UserMerchantUpdateResponse{
		FullName:     updateUserMerchant.FullName,
		UpdatedBy:    updateUserMerchant.UpdatedBy,
		Email:        updateUserMerchant.Email,
		Whatsapp:     updateUserMerchant.Whatsapp,
		Password:     updateUserMerchant.Password,
		AccessRoleID: updateUserMerchant.AccessRoleID,
		Verified: dto.UserMerchantVerifiedByID{
			Verified:      updateUserMerchant.Verified.Verified,
			StatusAccount: updateUserMerchant.Verified.StatusAccount,
		},
	}

	return response, nil
}
