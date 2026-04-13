package roleuserpermission

import (
	dto "github.com/srv-api/detail/dto"
)

func (s *roleuserpermissionService) Get(req dto.RoleUserPermissionRequest) (dto.GetRoleUserPermissionResponse, error) {
	products, _ := s.Repo.Get(req)

	return products, nil
}
