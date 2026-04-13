package permission

import (
	dto "github.com/srv-api/merchant/dto"
)

func (s *permissionService) Get(req dto.RoleUserRequest) (dto.GetPermissionResponse, error) {
	products, _ := s.Repo.Get(req)

	return products, nil
}
