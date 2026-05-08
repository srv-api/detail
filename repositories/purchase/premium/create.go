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

// Method baru: create purchase dan update user limit dalam 1 transaction
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

	// 2. Update user limit to unlimited
	var userLimit entity.UserLimit
	result := tx.Where("user_id = ?", req.UserID).First(&userLimit)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// Create new record if not exists
			userLimit = entity.UserLimit{
				UserID:             req.UserID,
				RemainingSwipe:     UNLIMITED,
				RemainingBoost:     UNLIMITED,
				RemainingSuperLike: UNLIMITED,
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
				"remaining_boost":      UNLIMITED,
				"remaining_super_like": UNLIMITED,
				"updated_at":           now,
			}).Error; err != nil {
			tx.Rollback()
			return dto.PremiumPurchaseResponse{}, err
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
