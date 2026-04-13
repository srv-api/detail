package roleuserpermission

import (
	dto "github.com/srv-api/detail/dto"
)

func (s *roleuserpermissionService) Create(req dto.RoleUserPermissionRequest) (dto.RoleUserPermissionResponse, error) {

	create := dto.RoleUserPermissionRequest{
		PermissionID: req.PermissionID,
		RoleUserID:   req.RoleUserID,
		UserID:       req.UserID,
		DetailID:     req.DetailID,
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
		DetailID:     created.DetailID,
		CreatedBy:    created.CreatedBy,
	}

	return response, nil
}
