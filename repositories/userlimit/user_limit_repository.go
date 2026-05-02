package userlimit

import (
	"time"

	"github.com/srv-api/detail/entity"
	"gorm.io/gorm"
)

func (r *userlimitRepository) FindByUserID(userID string) (*entity.UserLimit, error) {
	var userLimit entity.UserLimit
	err := r.DB.Where("user_id = ?", userID).First(&userLimit).Error
	if err != nil {
		return nil, err
	}
	return &userLimit, nil
}

func (r *userlimitRepository) Update(userLimit *entity.UserLimit) error {
	return r.DB.Save(userLimit).Error
}

func (r *userlimitRepository) IncrementLimits(userID string, swipeIncrement, superLikeIncrement int) error {
	return r.DB.Model(&entity.UserLimit{}).
		Where("user_id = ?", userID).
		Updates(map[string]interface{}{
			"remaining_swipe":      gorm.Expr("remaining_swipe + ?", swipeIncrement),
			"remaining_super_like": gorm.Expr("remaining_super_like + ?", superLikeIncrement),
			"is_premium":           true,
			"updated_at":           time.Now(),
		}).Error
}

func (r *userlimitRepository) UpdatePremiumStatus(userID string, isPremium bool, expiry *time.Time) error {
	return r.DB.Model(&entity.UserLimit{}).
		Where("user_id = ?", userID).
		Updates(map[string]interface{}{
			"is_premium":     isPremium,
			"premium_expiry": expiry,
			"updated_at":     time.Now(),
		}).Error
}

func (r *userlimitRepository) CreateOrUpdate(userLimit *entity.UserLimit) error {
	return r.DB.Where("user_id = ?", userLimit.UserID).
		Assign(userLimit).
		FirstOrCreate(userLimit).Error
}
