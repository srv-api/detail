package userdetail

import (
	"errors"

	"github.com/srv-api/auth/entity"
	limit "github.com/srv-api/detail/entity"
	"gorm.io/gorm"

	dto "github.com/srv-api/detail/dto"
)

func (r *userdetailRepository) Explore(req dto.UserDetailRequest) ([]dto.ExploreUserResponse, error) {
	var results []dto.ExploreUserResponse

	// 🔥 1. Ambil user limit
	var userLimit limit.UserLimit
	r.DB.Where("user_id = ?", req.UserID).First(&userLimit)

	// default fallback kalau belum ada record
	limit := userLimit.RemainingSwipe
	if limit <= 0 {
		// kalau swipe habis → tidak kasih user
		return []dto.ExploreUserResponse{}, nil
	}

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
		LEFT JOIN profile_pictures uf ON uf.user_id = ud.user_id 
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
			AND NOT EXISTS (
				SELECT 1 FROM likes l 
				WHERE l.user_id = current.user_id 
					AND l.target_user_id = ud.user_id
			)
			AND NOT EXISTS (
				SELECT 1 FROM matches m 
				WHERE (m.user1_id = current.user_id AND m.user2_id = ud.user_id)
					OR (m.user1_id = ud.user_id AND m.user2_id = current.user_id)
			)
		ORDER BY distance
		LIMIT ?
	`

	// 🔥 2. pakai LIMIT dari user_limit
	err := r.DB.Raw(query, req.UserID, limit).Scan(&results).Error
	if err != nil {
		return nil, err
	}

	// 🔥 3. Filter gender target (TETAP)
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

func (r *userdetailRepository) GetUserLimit(userID string) (*limit.UserLimit, error) {
	var userLimit limit.UserLimit
	err := r.DB.Where("user_id = ?", userID).First(&userLimit).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Auto create jika belum ada (sama seperti di like service)
			userLimit = limit.UserLimit{
				UserID:             userID,
				RemainingSwipe:     50,
				RemainingSuperLike: 1,
			}
			if createErr := r.DB.Create(&userLimit).Error; createErr != nil {
				return nil, createErr
			}
			return &userLimit, nil
		}
		return nil, err
	}

	return &userLimit, nil
}
