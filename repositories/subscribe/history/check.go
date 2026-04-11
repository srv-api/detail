package history

import (
	"github.com/srv-api/merchant/entity"
)

func (r *historyRepository) FindByOrderID(orderID string) (*entity.Subscribe, error) {
	var sub entity.Subscribe
	if err := r.DB.Where("order_id = ?", orderID).First(&sub).Error; err != nil {
		return nil, err
	}
	return &sub, nil
}

func (r *historyRepository) UpdateStatus(orderID, status string) error {
	return r.DB.Model(&entity.Subscribe{}).
		Where("order_id = ?", orderID).
		Update("status", status).Error
}
