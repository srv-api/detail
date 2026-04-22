package repository

import (
	"errors"

	"github.com/srv-api/detail/dto"
	"github.com/srv-api/detail/entity"
	"gorm.io/gorm"
)

type LikeRepository interface {
	CreateLike(req dto.LikeRequest) error
	IsMatch(req dto.LikeRequest) (bool, error)
	DeductSwipe(req dto.LikeRequest) error
	DeductSuperLike(req dto.LikeRequest) error
}

type likeRepository struct {
	DB *gorm.DB
}

func NewLikeRepository(db *gorm.DB) LikeRepository {
	return &likeRepository{DB: db}
}

// insert like
func (r *likeRepository) CreateLike(req dto.LikeRequest) error {
	like := entity.Like{
		UserID:       req.UserID,
		TargetUserID: req.TargetUserID,
	}

	// hindari duplicate
	err := r.DB.Create(&like).Error
	if err != nil {
		return err
	}

	return nil
}

// cek apakah match (mutual like)
func (r *likeRepository) IsMatch(req dto.LikeRequest) (bool, error) {
	var count int64

	err := r.DB.Model(&entity.Like{}).
		Where("user_id = ? AND target_user_id = ?", req.TargetUserID, req.UserID).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *likeRepository) DeductSwipe(req dto.LikeRequest) error {
	// 🔥 LOG 1: Cek function terpanggil
	println("=== DEDUCT SWIPE CALLED ===")
	println("UserID:", req.UserID)
	println("TargetUserID:", req.TargetUserID)
	println("IsSuperLike:", req.IsSuperLike)

	result := r.DB.Model(&entity.UserLimit{}).
		Where("user_id = ? AND remaining_swipe > 0", req.UserID).
		Update("remaining_swipe", gorm.Expr("remaining_swipe - 1"))

	// 🔥 LOG 2: Cek result
	println("RowsAffected:", result.RowsAffected)
	println("Error:", result.Error)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		// 🔥 LOG 3: Cek apakah user_limit ada
		var count int64
		r.DB.Model(&entity.UserLimit{}).Where("user_id = ?", req.UserID).Count(&count)
		println("UserLimit record exists:", count > 0)

		if count == 0 {
			println("USER LIMIT NOT FOUND FOR USER:", req.UserID)
		} else {
			// Cek nilai remaining_swipe
			var userLimit entity.UserLimit
			r.DB.Where("user_id = ?", req.UserID).First(&userLimit)
			println("Current remaining_swipe:", userLimit.RemainingSwipe)
		}

		return errors.New("no swipe remaining")
	}

	println("✅ SWIPE DEDUCTED SUCCESSFULLY")
	return nil
}

func (r *likeRepository) DeductSuperLike(req dto.LikeRequest) error {
	result := r.DB.Model(&entity.UserLimit{}).
		Where("user_id = ? AND remaining_super_like > 0", req.UserID).
		Update("remaining_super_like", gorm.Expr("remaining_super_like - 1"))

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no super like remaining")
	}

	return nil
}
