package roleuser

import (
	"encoding/json"

	dto "github.com/srv-api/detail/dto"
)

func (r *RoleUserRepository) Get(req dto.RoleUserRequest) (dto.GetRoleUserResponse, error) {
	var rows []struct {
		ID           uint
		RoleID       string
		RoleName     string
		PermissionID []byte
		UserID       string
		DetailID     string
		CreatedBy    string

		// Permission Items
		PermID uint
		Label  string
		Icon   string
		To     string
	}

	err := r.DB.Table("role_users AS ru").
		Select(`
			ru.id,
			ru.role_id,
			r.role AS role_name,
			ru.permission_id,
			ru.user_id,
			ru.detail_id,
			p.id AS perm_id,
			p.label,
			p.icon,
			p.to
		`).
		Joins(`JOIN roles r ON r.id = ru.role_id`).
		Joins(`
			LEFT JOIN permissions AS p 
				ON ru.permission_id::jsonb @> ('[' || p.id || ']')::jsonb
		`).
		Where("ru.user_id = ?", req.UserID).
		Where("ru.detail_id = ?", req.DetailID).
		Scan(&rows).Error

	if err != nil {
		return dto.GetRoleUserResponse{}, err
	}

	roleMap := map[uint]*dto.RoleUserResponse{}

	for _, row := range rows {

		if _, ok := roleMap[row.ID]; !ok {

			var permIDs []uint
			json.Unmarshal(row.PermissionID, &permIDs)

			roleMap[row.ID] = &dto.RoleUserResponse{
				ID:          row.ID,
				RoleID:      row.RoleID,
				RoleName:    row.RoleName, // ← Tambahkan
				UserID:      row.UserID,
				DetailID:    row.DetailID,
				CreatedBy:   row.CreatedBy,
				Permissions: []dto.PermissionItem{},
			}
		}

		if row.PermID != 0 {
			roleMap[row.ID].Permissions = append(
				roleMap[row.ID].Permissions,
				dto.PermissionItem{
					ID:    row.PermID,
					Label: row.Label,
					Icon:  row.Icon,
					To:    row.To,
				},
			)
		}
	}

	var result []dto.RoleUserResponse
	for _, v := range roleMap {
		result = append(result, *v)
	}

	return dto.GetRoleUserResponse{
		Roles: result,
	}, nil
}
