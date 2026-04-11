package role

import (
	dto "github.com/srv-api/merchant/dto"
)

func (s *roleService) Create(req dto.RoleRequest) (dto.RoleResponse, error) {

	create := dto.RoleRequest{
		Role:       req.Role,
		UserID:     req.UserID,
		MerchantID: req.MerchantID,
		CreatedBy:  req.CreatedBy,
	}

	created, err := s.Repo.Create(create)
	if err != nil {
		return dto.RoleResponse{}, err
	}

	response := dto.RoleResponse{
		Role: created.Role,
	}

	return response, nil
}
