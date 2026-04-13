package user

import "github.com/srv-api/merchant/dto"

func (b *userService) Update(req dto.UserMerchantUpdateRequest) (dto.UserMerchantUpdateResponse, error) {
	request := dto.UserMerchantUpdateRequest{
		FullName:     req.FullName,
		UpdatedBy:    req.UpdatedBy,
		Email:        req.Email,
		Whatsapp:     req.Whatsapp,
		Password:     req.Password,
		AccessRoleID: req.AccessRoleID,
		Verified: dto.UserMerchantVerifiedByID{
			Verified:      req.Verified.Verified,
			StatusAccount: req.Verified.StatusAccount,
		},
	}

	user, err := b.Repo.Update(req)
	if err != nil {
		return user, err
	}

	response := dto.UserMerchantUpdateResponse{
		FullName:     request.FullName,
		UpdatedBy:    request.UpdatedBy,
		Email:        request.Email,
		Whatsapp:     request.Whatsapp,
		Password:     request.Password,
		AccessRoleID: request.AccessRoleID,
		Verified: dto.UserMerchantVerifiedByID{
			Verified:      request.Verified.Verified,
			StatusAccount: request.Verified.StatusAccount,
		},
	}

	return response, nil
}
