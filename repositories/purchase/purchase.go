package purchase

import "github.com/srv-api/detail/entity"

func (r *purchaseRepository) Create(purchase *entity.PurchaseHistory) error {
	return r.DB.Create(purchase).Error
}

func (r *purchaseRepository) FindByToken(token string) (*entity.PurchaseHistory, error) {
	var purchase entity.PurchaseHistory
	err := r.DB.Where("token = ?", token).First(&purchase).Error
	if err != nil {
		return nil, err
	}
	return &purchase, nil
}

func (r *purchaseRepository) FindByUserID(userID string) ([]entity.PurchaseHistory, error) {
	var purchases []entity.PurchaseHistory
	err := r.DB.Where("user_id = ?", userID).Order("created_at desc").Find(&purchases).Error
	return purchases, err
}

func (r *purchaseRepository) UpdateStatus(id string, status string) error {
	return r.DB.Model(&entity.PurchaseHistory{}).
		Where("id = ?", id).
		Update("status", status).Error
}

func (r *purchaseRepository) ExistsByToken(token string) (bool, error) {
	var count int64
	err := r.DB.Model(&entity.PurchaseHistory{}).
		Where("token = ?", token).
		Count(&count).Error
	return count > 0, err
}
