package roleuser

import (
	"fmt"
	"math"
	"strings"

	dto "github.com/srv-api/detail/dto"
)

func (r *RoleUserRepository) Pagination(req *dto.Pagination) (RepositoryResult, int) {
	var roleusers []dto.RoleUserResponse
	var totalRows int64

	offset := (req.Page - 1) * req.Limit

	// Perbaikan: Gunakan kolom yang benar dari masing-masing tabel
	find := r.DB.Table("role_users AS ru").
		Select(`
			ru.id,
			ru.role_id,
			ru.user_id,
			ru.detail_id,
			ru.created_by,
			ru.permission_id,
			roles.role AS role_name,
			permissions.id AS permission_id,
			permissions.label AS permission_label,
			permissions.icon AS permission_icon,
			permissions.to AS permission_to,
			access_doors.full_name AS user_full_name
		`).
		Joins("LEFT JOIN roles ON roles.id = ru.role_id"). // Ubah dari roles.role_id menjadi roles.id
		Joins("LEFT JOIN permissions ON permissions.id = ru.permission_id").
		Joins("LEFT JOIN access_doors ON access_doors.user_id = ru.user_id").
		Limit(req.Limit).
		Offset(offset)

	// Search filter
	if req.Searchs != nil {
		for _, s := range req.Searchs {
			switch s.Action {
			case "equals":
				find = find.Where(fmt.Sprintf("%s = ?", s.Column), s.Query)
			case "contains":
				find = find.Where(fmt.Sprintf("%s LIKE ?", s.Column), "%"+s.Query+"%")
			case "in":
				find = find.Where(fmt.Sprintf("%s IN (?)", s.Column), strings.Split(s.Query, ","))
			}
		}
	}

	find = find.Order(req.Sort)

	// Eksekusi
	if err := find.Scan(&roleusers).Error; err != nil {
		return RepositoryResult{Error: err}, 0
	}

	req.Rows = roleusers

	// COUNT query
	countQuery := r.DB.Table("role_users AS ru").
		Joins("LEFT JOIN roles ON roles.id = ru.role_id").
		Joins("LEFT JOIN permissions ON permissions.id = ru.permission_id").
		Joins("LEFT JOIN access_doors ON access_doors.user_id = ru.user_id")

	if err := countQuery.Count(&totalRows).Error; err != nil {
		return RepositoryResult{Error: err}, 0
	}

	req.TotalRows = int(totalRows)
	totalPages := int(math.Ceil(float64(totalRows) / float64(req.Limit)))
	req.TotalPages = totalPages

	// fromRow / toRow
	if req.Page == 1 {
		req.FromRow = 1
		req.ToRow = len(roleusers)
	} else {
		req.FromRow = (req.Page-1)*req.Limit + 1
		req.ToRow = req.FromRow + len(roleusers) - 1
	}

	return RepositoryResult{Result: req}, totalPages
}
