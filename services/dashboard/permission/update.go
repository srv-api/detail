package permission

import "github.com/srv-api/merchant/dto"

func (b *permissionService) Update(req dto.PermissionUpdateRequest) (dto.PermissionUpdateResponse, error) {
	request := dto.PermissionUpdateRequest{
		Label:     req.Label,
		Icon:      req.Icon,
		UpdatedBy: req.UpdatedBy,
		UserID:    req.UserID,
		To:        req.To,
	}

	permission, err := b.Repo.Update(req)
	if err != nil {
		return permission, err
	}

	response := dto.PermissionUpdateResponse{
		Label:     req.Label,
		Icon:      req.Icon,
		UpdatedBy: request.UpdatedBy,
		UserID:    request.UserID,
		To:        req.To,
	}

	return response, nil
}
