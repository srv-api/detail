package userdetail

import (
	"errors"

	"github.com/srv-api/auth/entity"
	dto "github.com/srv-api/detail/dto"
	explore "github.com/srv-api/detail/entity"
	"gorm.io/gorm"
)

func (r *userdetailRepository) Explore(req dto.UserDetailRequest) ([]dto.ExploreUserResponse, error) {
	var results []dto.ExploreUserResponse

	// Cek user limit terlebih dahulu
	var userLimit explore.UserLimit
	err := r.DB.Where("user_id = ?", req.UserID).First(&userLimit).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Buat user limit baru jika belum ada
			userLimit = explore.UserLimit{
				UserID:             req.UserID,
				RemainingSwipe:     50,
				RemainingSuperLike: 1,
			}
			if createErr := r.DB.Create(&userLimit).Error; createErr != nil {
				return nil, createErr
			}
		} else {
			return nil, err
		}
	}

	// Cek apakah masih punya sisa swipe
	if userLimit.RemainingSwipe <= 0 {
		return nil, errors.New("daily swipe limit exceeded. Please try again tomorrow")
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
			-- Exclude users who have been liked by current user (table: likes)
			AND NOT EXISTS (
				SELECT 1 FROM likes l 
				WHERE l.user_id = current.user_id 
					AND l.target_user_id = ud.user_id
			)
			-- Exclude users who have matched with current user (table: matches)
			AND NOT EXISTS (
				SELECT 1 FROM matches m 
				WHERE (m.user1_id = current.user_id AND m.user2_id = ud.user_id)
					OR (m.user1_id = ud.user_id AND m.user2_id = current.user_id)
			)
		ORDER BY distance
	`

	err = r.DB.Raw(query, req.UserID).Scan(&results).Error
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

// DeductSwipe mengurangi remaining swipe user
func (r *userdetailRepository) DeductSwipe(userID string) error {
	result := r.DB.Model(&explore.UserLimit{}).
		Where("user_id = ? AND remaining_swipe > 0", userID).
		Update("remaining_swipe", gorm.Expr("remaining_swipe - ?", 1))

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no remaining swipe available")
	}

	return nil
}

// DeductSuperLike mengurangi remaining super like user
func (r *userdetailRepository) DeductSuperLike(userID string) error {
	result := r.DB.Model(&explore.UserLimit{}).
		Where("user_id = ? AND remaining_super_like > 0", userID).
		Update("remaining_super_like", gorm.Expr("remaining_super_like - ?", 1))

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no remaining super like available")
	}

	return nil
}

// ResetDailySwipe mereset swipe setiap hari (bisa dijalankan via cron job)
func (r *userdetailRepository) ResetDailySwipe() error {
	return r.DB.Model(&explore.UserLimit{}).
		Where("updated_at < DATE_SUB(NOW(), INTERVAL 1 DAY)").
		Updates(map[string]interface{}{
			"remaining_swipe":      50,
			"remaining_super_like": 1,
		}).Error
}

// GetUserLimit mendapatkan informasi limit user
func (r *userdetailRepository) GetUserLimit(userID string) (*explore.UserLimit, error) {
	var userLimit explore.UserLimit
	err := r.DB.Where("user_id = ?", userID).First(&userLimit).Error
	if err != nil {
		return nil, err
	}
	return &userLimit, nil
}
