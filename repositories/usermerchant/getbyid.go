package user

import (
	"github.com/srv-api/auth/entity"
	dto "github.com/srv-api/merchant/dto"
	mentity "github.com/srv-api/merchant/entity"
	util "github.com/srv-api/util/s"
)

func (b *userRepository) GetById(req dto.GetByIdRequest) (*dto.UserMerchantByIdResponse, error) {
	var user entity.AccessDoor

	// Ambil data user + relasi Verified dan Merchant (kalau perlu)
	if err := b.DB.
		Preload("Verified").
		Preload("Merchant").
		Where("id = ?", req.ID).
		Take(&user).Error; err != nil {
		return nil, err
	}

	// ✅ Ambil role name berdasarkan access_role_id
	var role mentity.Role
	_ = b.DB.First(&role, "id = ?", user.AccessRoleID)

	// ✅ Ubah boolean menjadi string

	// ✅ Dekripsi data jika terenkripsi
	decryptedWa, _ := util.Decrypt(user.Whatsapp)
	decryptedEmail, _ := util.Decrypt(user.Email)

	// ✅ Buat response DTO
	response := &dto.UserMerchantByIdResponse{
		ID:            user.ID,
		MerchantID:    user.MerchantID,
		FullName:      user.FullName,
		Whatsapp:      decryptedWa,
		Email:         decryptedEmail,
		Password:      user.Password,
		AccessRoleID:  user.AccessRoleID,
		RoleName:      role.Role,
		LoginAttempts: user.LoginAttempts,
		Suspended:     user.Suspended,
		LastAttempt:   user.LastAttempt,
		CreatedBy:     user.CreatedBy,
		UpdatedBy:     user.UpdatedBy,
		DeletedBy:     user.DeletedBy,
		CreatedAt:     user.CreatedAt,
		Verified: dto.UserMerchantVerifiedByID{
			ID:             user.Verified.ID,
			UserID:         user.Verified.UserID,
			Token:          user.Verified.Token,
			Verified:       user.Verified.Verified,      // ✅ ubah ke string
			StatusAccount:  user.Verified.StatusAccount, // ✅ ubah ke string
			AccountExpired: user.Verified.AccountExpired,
			Otp:            user.Verified.Otp,
			ExpiredAt:      user.Verified.ExpiredAt,
		},
	}

	return response, nil
}
