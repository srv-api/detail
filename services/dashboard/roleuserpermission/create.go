package roleuserpermission

import (
	dto "github.com/srv-api/merchant/dto"
)

func (s *roleuserpermissionService) Create(req dto.RoleUserPermissionRequest) (dto.RoleUserPermissionResponse, error) {

	create := dto.RoleUserPermissionRequest{
		PermissionID: req.PermissionID,
		RoleUserID:   req.RoleUserID,
		UserID:       req.UserID,
		MerchantID:   req.MerchantID,
		CreatedBy:    req.CreatedBy,
	}

	created, err := s.Repo.Create(create)
	if err != nil {
		return dto.RoleUserPermissionResponse{}, err
	}

	response := dto.RoleUserPermissionResponse{
		PermissionID: created.PermissionID,
		RoleUserID:   created.RoleUserID,
		UserID:       created.UserID,
		MerchantID:   created.MerchantID,
		CreatedBy:    created.CreatedBy,
	}

	return response, nil
}
