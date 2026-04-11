package roleuser

import "github.com/srv-api/merchant/dto"

func (b *roleuserService) Update(req dto.RoleUserUpdateRequest) (dto.RoleUserUpdateResponse, error) {
	request := dto.RoleUserUpdateRequest{
		RoleID:    req.RoleID,
		UpdatedBy: req.UpdatedBy,
		UserID:    req.UserID,
	}

	role, err := b.Repo.Update(req)
	if err != nil {
		return role, err
	}

	response := dto.RoleUserUpdateResponse{
		RoleID:    req.RoleID,
		UpdatedBy: request.UpdatedBy,
		UserID:    request.UserID,
	}

	return response, nil
}
