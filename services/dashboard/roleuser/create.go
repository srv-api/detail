package roleuser

import (
	dto "github.com/srv-api/merchant/dto"
)

func (s *roleuserService) Create(req dto.RoleUserRequest) (dto.RoleUserResponse, error) {

	create := dto.RoleUserRequest{
		RoleID:       req.RoleID,
		UserID:       req.UserID,
		PermissionID: req.PermissionID,
		MerchantID:   req.MerchantID,
		CreatedBy:    req.CreatedBy,
	}

	created, err := s.Repo.Create(create)
	if err != nil {
		return dto.RoleUserResponse{}, err
	}

	response := dto.RoleUserResponse{
		RoleID:       created.RoleID,
		UserID:       created.UserID,
		PermissionID: created.PermissionID,
		MerchantID:   created.MerchantID,
		CreatedBy:    created.CreatedBy,
	}

	return response, nil
}
