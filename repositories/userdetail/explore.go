package userdetail

import (
	"github.com/srv-api/auth/entity"
	dto "github.com/srv-api/detail/dto"
)

func (r *userdetailRepository) Explore(req dto.UserDetailRequest) ([]dto.ExploreUserResponse, error) {
	var results []dto.ExploreUserResponse

	query := `
		SELECT 
			ud.user_id,
			a.full_name,
			a.gender,
			ud.latitude,
			ud.longitude,
			COALESCE(ud.bio, '') as bio,
			ud.radius,
			ud.min_age,
			ud.max_age,
			ud.gender_target,
			a.age,
			COALESCE(uf.file_path, '') as profile_picture,
			(6371 * acos(
				LEAST(1, GREATEST(-1,
					cos(radians(current.latitude)) * cos(radians(ud.latitude)) * cos(radians(ud.longitude) - radians(current.longitude)) +
					sin(radians(current.latitude)) * sin(radians(ud.latitude))
				))
			)) AS distance
		FROM user_details current
		CROSS JOIN user_details ud
		JOIN access_doors a ON a.id = ud.user_id
		LEFT JOIN uploaded_files uf ON uf.user_id = ud.user_id 
			AND uf.deleted_at IS NULL
		WHERE current.user_id = ?
			AND ud.user_id != current.user_id
			AND ud.latitude IS NOT NULL 
			AND ud.longitude IS NOT NULL
			AND ud.latitude != 0
			AND ud.longitude != 0
			AND a.age BETWEEN current.min_age AND current.max_age
			AND (6371 * acos(
				LEAST(1, GREATEST(-1,
					cos(radians(current.latitude)) * cos(radians(ud.latitude)) * cos(radians(ud.longitude) - radians(current.longitude)) +
					sin(radians(current.latitude)) * sin(radians(ud.latitude))
				))
			)) <= current.radius
		ORDER BY distance
	`

	err := r.DB.Raw(query, req.UserID).Scan(&results).Error
	if err != nil {
		return nil, err
	}

	// Filter gender target jika tidak 'both'
	var currentUser entity.AccessDoor
	r.DB.Where("id = ?", req.UserID).First(&currentUser)

	if currentUser.Merchant.GenderTarget != "both" && currentUser.Merchant.GenderTarget != "" {
		filtered := make([]dto.ExploreUserResponse, 0)
		for _, u := range results {
			if u.Gender == currentUser.Merchant.GenderTarget {
				filtered = append(filtered, u)
			}
		}
		results = filtered
	}

	return results, nil
}
