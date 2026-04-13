package user

import "github.com/srv-api/detail/dto"

func (b *userService) Update(req dto.UserDetailUpdateRequest) (dto.UserDetailUpdateResponse, error) {
	request := dto.UserDetailUpdateRequest{
		FullName:     req.FullName,
		UpdatedBy:    req.UpdatedBy,
		Email:        req.Email,
		Whatsapp:     req.Whatsapp,
		Password:     req.Password,
		AccessRoleID: req.AccessRoleID,
		Verified: dto.UserDetailVerifiedByID{
			Verified:      req.Verified.Verified,
			StatusAccount: req.Verified.StatusAccount,
		},
	}

	user, err := b.Repo.Update(req)
	if err != nil {
		return user, err
	}

	response := dto.UserDetailUpdateResponse{
		FullName:     request.FullName,
		UpdatedBy:    request.UpdatedBy,
		Email:        request.Email,
		Whatsapp:     request.Whatsapp,
		Password:     request.Password,
		AccessRoleID: request.AccessRoleID,
		Verified: dto.UserDetailVerifiedByID{
			Verified:      request.Verified.Verified,
			StatusAccount: request.Verified.StatusAccount,
		},
	}

	return response, nil
}
