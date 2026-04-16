package repository

import (
	"errors"

	"github.com/srv-api/detail/entity"
	"gorm.io/gorm"
)

type LikeRepository interface {
	CreateLike(userID, targetUserID string) error
	IsMatch(userID, targetUserID string) (bool, error)
	DeductSwipe(userID string) error
	DeductSuperLike(userID string) error
}

type likeRepository struct {
	DB *gorm.DB
}

func NewLikeRepository(db *gorm.DB) LikeRepository {
	return &likeRepository{DB: db}
}

// insert like
func (r *likeRepository) CreateLike(userID, targetUserID string) error {
	like := entity.Like{
		UserID:       userID,
		TargetUserID: targetUserID,
	}

	// hindari duplicate
	err := r.DB.Create(&like).Error
	if err != nil {
		return err
	}

	return nil
}

// cek apakah match (mutual like)
func (r *likeRepository) IsMatch(userID, targetUserID string) (bool, error) {
	var count int64

	err := r.DB.Model(&entity.Like{}).
		Where("user_id = ? AND target_user_id = ?", targetUserID, userID).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *likeRepository) DeductSwipe(userID string) error {
	result := r.DB.Model(&entity.UserLimit{}).
		Where("user_id = ? AND remaining_swipe > 0", userID).
		Update("remaining_swipe", gorm.Expr("remaining_swipe - 1"))

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no swipe remaining")
	}

	return nil
}

func (r *likeRepository) DeductSuperLike(userID string) error {
	result := r.DB.Model(&entity.UserLimit{}).
		Where("user_id = ? AND remaining_super_like > 0", userID).
		Update("remaining_super_like", gorm.Expr("remaining_super_like - 1"))

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no super like remaining")
	}

	return nil
}
