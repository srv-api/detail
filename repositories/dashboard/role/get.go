package role

import (
	dto "github.com/srv-api/merchant/dto"
)

func (r *RoleRepository) Get(req dto.RoleUserRequest) (dto.GetRoleResponse, error) {
	var roles []dto.RoleResponse

	// Query untuk memvalidasi RoleUser dan Role
	err := r.DB.Table("role_user_roles").
		Select("roles.label, roles.icon, roles.to").
		Joins("JOIN role_users ON role_user_roles.role_user_id = role_users.id").
		Joins("JOIN roles ON role_user_roles.role_id = roles.id").
		Where("role_users.user_id = ? AND role_users.role_id = ?", req.UserID, "8gHwINv71XDy"). // Validasi RoleID dan UserID
		Scan(&roles).Error

	if err != nil {
		return dto.GetRoleResponse{}, err
	}

	// Bungkus izin dalam field 'items' seperti yang diharapkan
	return dto.GetRoleResponse{
		Role: roles,
	}, nil
}
