package subscribe

import (
	"time"

	auth "github.com/srv-api/auth/entity"
	dto "github.com/srv-api/merchant/dto"
	"github.com/srv-api/merchant/entity"
)

func (r *subscribeRepository) Create(req dto.SubscribeRequest) (dto.SubscribeResponse, error) {
	// 1. Membuat entitas package berdasarkan request

	create := entity.Subscribe{
		ID:          req.ID,
		UserID:      req.UserID,
		CreatedBy:   req.CreatedBy,
		Status:      req.Status,
		GrossAmount: req.GrossAmount,
		OrderID:     req.OrderID, // Order ID di sini digunakan untuk transaksi Midtrans
	}
	// 2. Simpan package ke database
	if err := r.DB.Save(&create).Error; err != nil {
		return dto.SubscribeResponse{}, err
	}

	// 3. Return response dengan data yang telah disimpan
	response := dto.SubscribeResponse{
		ID:          create.ID,
		UserID:      create.UserID,
		CreatedBy:   create.CreatedBy,
		Status:      create.Status,
		GrossAmount: create.GrossAmount,
		OrderID:     create.OrderID,
	}

	return response, nil
}

func (r *subscribeRepository) UpdateStatus(orderID string, status string) error {
	// Update status berdasarkan OrderID
	if err := r.DB.Model(&entity.Subscribe{}).
		Where("order_id = ?", orderID).
		Update("status", status).Error; err != nil {
		return err
	}
	return nil
}

func (r *subscribeRepository) UpdateUserVerified(orderID string) error {
	// Cari UserID berdasarkan OrderID
	var pkg entity.Subscribe
	if err := r.DB.Where("order_id = ?", orderID).First(&pkg).Error; err != nil {
		return err
	}

	// Update UserVerified
	if err := r.DB.Model(&auth.UserVerified{}).
		Where("user_id = ?", pkg.UserID).
		Updates(map[string]interface{}{
			"status_account":  true,
			"account_expired": time.Now().Add(30 * 24 * time.Hour), // fix 30 hari
		}).Error; err != nil {
		return err
	}

	return nil
}
