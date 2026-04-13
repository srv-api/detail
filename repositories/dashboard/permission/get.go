package permission

import (
	dto "github.com/srv-api/merchant/dto"
)

func (r *PermissionRepository) Get(req dto.RoleUserRequest) (dto.GetPermissionResponse, error) {
	var permissions []dto.PermissionResponse

	// Query untuk memvalidasi RoleUser dan Permission
	err := r.DB.Table("role_user_permissions").
		Select("permissions.label, permissions.icon, permissions.to").
		Joins("JOIN role_users ON role_user_permissions.role_user_id = role_users.id").
		Joins("JOIN permissions ON role_user_permissions.permission_id = permissions.id").
		Where("role_users.user_id = ? AND role_users.role_id = ?", req.UserID, "8gHwINv71XDy"). // Validasi RoleID dan UserID
		Scan(&permissions).Error

	if err != nil {
		return dto.GetPermissionResponse{}, err
	}

	// Bungkus izin dalam field 'items' seperti yang diharapkan
	return dto.GetPermissionResponse{
		Items: permissions,
	}, nil
}
