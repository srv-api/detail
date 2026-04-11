package role

import "github.com/srv-api/merchant/dto"

func (b *roleService) Update(req dto.RoleUpdateRequest) (dto.RoleUpdateResponse, error) {
	request := dto.RoleUpdateRequest{
		Role:      req.Role,
		UpdatedBy: req.UpdatedBy,
		UserID:    req.UserID,
	}

	role, err := b.Repo.Update(req)
	if err != nil {
		return role, err
	}

	response := dto.RoleUpdateResponse{
		Role:      req.Role,
		UpdatedBy: request.UpdatedBy,
		UserID:    request.UserID,
	}

	return response, nil
}
