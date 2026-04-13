package role

import (
	dto "github.com/srv-api/detail/dto"
)

func (s *roleService) RoleUser(req dto.GetRoleRequest) (dto.GetRoleResponse, error) {
	return s.Repo.RoleUser(req)
}
