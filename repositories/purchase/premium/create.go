package purchase

import (
	dto "github.com/srv-api/detail/dto"
	"github.com/srv-api/detail/entity"
)

func (r *purchaseRepository) Create(req dto.PremiumPurchaseRequest) (dto.PremiumPurchaseResponse, error) {

	create := entity.PurchasePremium{
		ID:        req.ID,
		UserID:    req.UserID,
		UpdatedAt: req.UpdatedAt,
	}

	if err := r.DB.Save(&create).Error; err != nil {
		return dto.PremiumPurchaseResponse{}, err
	}

	response := dto.PremiumPurchaseResponse{
		ID:        req.ID,
		DetailID:  req.DetailID,
		UserID:    req.UserID,
		CreatedBy: req.CreatedBy,
	}

	return response, nil

}
