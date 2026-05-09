package purchase

import (
	"errors"
	"time"

	dto "github.com/srv-api/detail/dto"
	"github.com/srv-api/detail/entity"
	"gorm.io/gorm"
)

func (r *purchaseRepository) Create(req dto.PremiumPurchaseRequest) (dto.PremiumPurchaseResponse, error) {
	now := time.Now()
	create := entity.PurchasePremium{
		ID:            req.ID,
		UserID:        req.UserID,
		DetailID:      req.DetailID,
		ProductID:     req.ProductID,
		TransactionID: req.TransactionID,
		PurchaseToken: req.PurchaseToken,
		ReceiptData:   req.ReceiptData,
		Signature:     req.Signature,
		CreatedBy:     req.CreatedBy,
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	if err := r.DB.Create(&create).Error; err != nil {
		return dto.PremiumPurchaseResponse{}, err
	}

	return dto.PremiumPurchaseResponse{
		ID:            create.ID,
		DetailID:      create.DetailID,
		UserID:        create.UserID,
		ProductID:     create.ProductID,
		TransactionID: create.TransactionID,
		PurchaseToken: create.PurchaseToken,
		ReceiptData:   create.ReceiptData,
		Signature:     create.Signature,
		CreatedBy:     create.CreatedBy,
	}, nil
}

// Method: create purchase dan update user limit berdasarkan produk
func (r *purchaseRepository) CreateWithPremium(req dto.PremiumPurchaseRequest) (dto.PremiumPurchaseResponse, error) {
	const UNLIMITED = -1

	// Start transaction
	tx := r.DB.Begin()
	if tx.Error != nil {
		return dto.PremiumPurchaseResponse{}, tx.Error
	}

	// Defer rollback jika terjadi error
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 1. Create purchase record
	now := time.Now()
	create := entity.PurchasePremium{
		ID:            req.ID,
		UserID:        req.UserID,
		DetailID:      req.DetailID,
		CreatedBy:     req.CreatedBy,
		ProductID:     req.ProductID,
		TransactionID: req.TransactionID,
		PurchaseToken: req.PurchaseToken,
		ReceiptData:   req.ReceiptData,
		Signature:     req.Signature,
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	if err := tx.Create(&create).Error; err != nil {
		tx.Rollback()
		return dto.PremiumPurchaseResponse{}, err
	}

	// 2. Update user limit berdasarkan product ID
	var userLimit entity.UserLimit
	result := tx.Where("user_id = ?", req.UserID).First(&userLimit)

	// ==================== PRODUCT P1 ====================
	if req.ProductID == "p1" {
		if err := tx.Model(&entity.UserDetail{}).
			Where("user_id = ?", req.UserID).
			Update("is_premium", true).Error; err != nil {
			tx.Rollback()
			return dto.PremiumPurchaseResponse{}, err
		}
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				// Create new record
				userLimit = entity.UserLimit{
					UserID:          req.UserID,
					RemainingSwipe:  UNLIMITED,
					RemainingRewind: UNLIMITED,
					UpdatedAt:       now,
				}
				if err := tx.Create(&userLimit).Error; err != nil {
					tx.Rollback()
					return dto.PremiumPurchaseResponse{}, err
				}
			} else {
				tx.Rollback()
				return dto.PremiumPurchaseResponse{}, result.Error
			}
		} else {
			// Update existing record
			if err := tx.Model(&entity.UserLimit{}).
				Where("user_id = ?", req.UserID).
				Updates(map[string]interface{}{
					"remaining_swipe":  UNLIMITED,
					"remaining_rewind": UNLIMITED,
					"updated_at":       now,
				}).Error; err != nil {
				tx.Rollback()
				return dto.PremiumPurchaseResponse{}, err
			}
		}
	}

	// ==================== PRODUCT P2 ====================
	// +5 swipe, +1 super like, +1 boost
	if req.ProductID == "p2" {
		if err := tx.Model(&entity.UserDetail{}).
			Where("user_id = ?", req.UserID).
			Update("is_premium", true).Error; err != nil {
			tx.Rollback()
			return dto.PremiumPurchaseResponse{}, err
		}
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				// Create new record
				userLimit = entity.UserLimit{
					UserID:             req.UserID,
					RemainingSwipe:     UNLIMITED,
					RemainingBoost:     1,
					RemainingSuperLike: 5,
					RemainingRewind:    UNLIMITED,
					UpdatedAt:          now,
				}
				if err := tx.Create(&userLimit).Error; err != nil {
					tx.Rollback()
					return dto.PremiumPurchaseResponse{}, err
				}
			} else {
				tx.Rollback()
				return dto.PremiumPurchaseResponse{}, result.Error
			}
		} else {
			// Update existing record
			if err := tx.Model(&entity.UserLimit{}).
				Where("user_id = ?", req.UserID).
				Updates(map[string]interface{}{
					"remaining_swipe":      UNLIMITED,
					"remaining_boost":      gorm.Expr("remaining_boost + ?", 1),
					"remaining_super_like": gorm.Expr("remaining_super_like + ?", 5),
					"remaining_rewind":     UNLIMITED,
					"updated_at":           now,
				}).Error; err != nil {
				tx.Rollback()
				return dto.PremiumPurchaseResponse{}, err
			}
		}
	}

	// ==================== PRODUCT P3 ====================
	// +10 swipe, +2 super like, +2 boost
	if req.ProductID == "p3" {
		if err := tx.Model(&entity.UserDetail{}).
			Where("user_id = ?", req.UserID).
			Update("is_premium", true).Error; err != nil {
			tx.Rollback()
			return dto.PremiumPurchaseResponse{}, err
		}
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				// Create new record
				userLimit = entity.UserLimit{
					UserID:             req.UserID,
					RemainingSwipe:     UNLIMITED,
					RemainingBoost:     2,
					RemainingSuperLike: 5,
					UpdatedAt:          now,
				}
				if err := tx.Create(&userLimit).Error; err != nil {
					tx.Rollback()
					return dto.PremiumPurchaseResponse{}, err
				}
			} else {
				tx.Rollback()
				return dto.PremiumPurchaseResponse{}, result.Error
			}
		} else {
			// Update existing record
			if err := tx.Model(&entity.UserLimit{}).
				Where("user_id = ?", req.UserID).
				Updates(map[string]interface{}{
					"remaining_swipe":      UNLIMITED,
					"remaining_boost":      gorm.Expr("remaining_boost + ?", 2),
					"remaining_super_like": gorm.Expr("remaining_super_like + ?", 5),
					"updated_at":           now,
				}).Error; err != nil {
				tx.Rollback()
				return dto.PremiumPurchaseResponse{}, err
			}
		}
	}

	// ==================== PRODUCT SL1 ====================
	// +6 super like
	if req.ProductID == "star_like_1" {
		if err := tx.Model(&entity.UserDetail{}).
			Where("user_id = ?", req.UserID).
			Update("is_star_like", true).Error; err != nil {
			tx.Rollback()
			return dto.PremiumPurchaseResponse{}, err
		}
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				// Create new record
				userLimit = entity.UserLimit{
					UserID:             req.UserID,
					RemainingSuperLike: 4,
					UpdatedAt:          now,
				}
				if err := tx.Create(&userLimit).Error; err != nil {
					tx.Rollback()
					return dto.PremiumPurchaseResponse{}, err
				}
			} else {
				tx.Rollback()
				return dto.PremiumPurchaseResponse{}, result.Error
			}
		} else {
			// Update existing record
			if err := tx.Model(&entity.UserLimit{}).
				Where("user_id = ?", req.UserID).
				Updates(map[string]interface{}{
					"remaining_super_like": gorm.Expr("remaining_super_like + ?", 6),
					"updated_at":           now,
				}).Error; err != nil {
				tx.Rollback()
				return dto.PremiumPurchaseResponse{}, err
			}
		}
	}

	// ==================== PRODUCT SL2 ====================
	// +13 super like
	if req.ProductID == "star_like_2" {
		if err := tx.Model(&entity.UserDetail{}).
			Where("user_id = ?", req.UserID).
			Update("is_star_like", true).Error; err != nil {
			tx.Rollback()
			return dto.PremiumPurchaseResponse{}, err
		}
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				// Create new record
				userLimit = entity.UserLimit{
					UserID:             req.UserID,
					RemainingSuperLike: 4,
					UpdatedAt:          now,
				}
				if err := tx.Create(&userLimit).Error; err != nil {
					tx.Rollback()
					return dto.PremiumPurchaseResponse{}, err
				}
			} else {
				tx.Rollback()
				return dto.PremiumPurchaseResponse{}, result.Error
			}
		} else {
			// Update existing record
			if err := tx.Model(&entity.UserLimit{}).
				Where("user_id = ?", req.UserID).
				Updates(map[string]interface{}{
					"remaining_super_like": gorm.Expr("remaining_super_like + ?", 13),
					"updated_at":           now,
				}).Error; err != nil {
				tx.Rollback()
				return dto.PremiumPurchaseResponse{}, err
			}
		}
	}
	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return dto.PremiumPurchaseResponse{}, err
	}

	// Return response
	return dto.PremiumPurchaseResponse{
		ID:            create.ID,
		DetailID:      create.DetailID,
		UserID:        create.UserID,
		ProductID:     create.ProductID,
		TransactionID: create.TransactionID,
		PurchaseToken: create.PurchaseToken,
		ReceiptData:   create.ReceiptData,
		Signature:     create.Signature,
		CreatedBy:     create.CreatedBy,
	}, nil
}
