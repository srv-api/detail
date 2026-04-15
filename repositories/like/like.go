package repository

import (
	"github.com/srv-api/detail/entity"
	"gorm.io/gorm"
)

type LikeRepository interface {
	CreateLike(userID, targetUserID string) error
	IsMatch(userID, targetUserID string) (bool, error)
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
