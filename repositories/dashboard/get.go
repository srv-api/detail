package dashboard

import (
	"time"

	"github.com/srv-api/merchant/dto"
	"github.com/srv-api/merchant/entity"
	entityPos "github.com/srv-api/pos/entity"
	entityProduct "github.com/srv-api/product/entity"
)

func (r *dashboardRepository) Get(req dto.GetDashboardRequest) (dto.GetDashboardResponse, error) {
	var response dto.GetDashboardResponse

	var isSubscribed bool
	var count int64
	err := r.DB.Model(&entity.Subscribe{}).
		Where("merchant_id = ? AND status = ? AND expiry_time > ?", req.MerchantID, "settlement", time.Now()).
		Count(&count).Error
	if err != nil {
		return response, err
	}
	isSubscribed = count > 0
	response.IsSubscribed = isSubscribed

	var merchant entity.UserDetail
	if err := r.DB.Where("id = ?", req.MerchantID).First(&merchant).Error; err != nil {
		return response, err
	}

	merchantName := merchant.ID
	if len(merchantName) > 25 {
		merchantName = merchantName[:25] + "..."
	}
	response.MerchantName = merchantName
	// Hitung total produk aktif
	var totalWebOrder int64
	if err := r.DB.Model(&entity.Order{}).
		Where("merchant_id = ? ", req.MerchantID).
		Count(&totalWebOrder).Error; err != nil {
		return dto.GetDashboardResponse{}, err
	}
	response.TotalWebOrder = totalWebOrder
	// Hitung total produk aktif
	if err := r.DB.Model(&entityProduct.Product{}).
		Where("merchant_id = ? AND status = ?", req.MerchantID, 1).
		Count(&req.TotalProductsActive).Error; err != nil {
		return dto.GetDashboardResponse{}, err
	}
	response.TotalProductsActive = req.TotalProductsActive

	// Hitung total produk tidak aktif
	if err := r.DB.Model(&entityProduct.Product{}).
		Where("merchant_id = ? AND status = ?", req.MerchantID, 2).
		Count(&req.TotalProductsInactive).Error; err != nil {
		return dto.GetDashboardResponse{}, err
	}
	response.TotalProductsInactive = req.TotalProductsInactive

	var totalSales int64
	if err := r.DB.Model(&entityPos.Pos{}).
		Where("merchant_id = ?", req.MerchantID).
		Count(&totalSales).Error; err != nil {
		return dto.GetDashboardResponse{}, err
	}
	response.TotalSales = totalSales

	if err := r.DB.Raw(`
		SELECT COALESCE(SUM((json_data->>'price')::int * (json_data->>'quantity')::int), 0) 
		FROM pos,
		LATERAL json_array_elements(pos.product) AS json_data
		WHERE pos.merchant_id = ? AND pos.deleted_at IS NULL 
		AND pos.status_payment IN ('Paid')
	`, req.MerchantID).Scan(&req.TotalPrice).Error; err != nil {
		return response, err
	}
	response.TotalPrice = req.TotalPrice

	if err := r.DB.Raw(`
		SELECT COALESCE(SUM((json_data->>'price')::int * (json_data->>'quantity')::int), 0) 
		FROM pos,
		LATERAL json_array_elements(pos.product) AS json_data
		WHERE pos.merchant_id = ? AND pos.deleted_at IS NULL 
		AND pos.status_payment IN ('Unpaid')
	`, req.MerchantID).Scan(&req.TotalWaiting).Error; err != nil {
		return response, err
	}
	response.TotalWaiting = req.TotalWaiting

	var waitingPayment int64
	if err := r.DB.Model(&entityPos.Pos{}).
		Where("merchant_id = ? AND status_payment = ?", req.MerchantID, "Unpaid").
		Count(&waitingPayment).Error; err != nil {
		return dto.GetDashboardResponse{}, err
	}
	response.WaitingPayment = waitingPayment

	if err := r.DB.Model(&entityPos.Pos{}).
		Where("merchant_id = ? AND status_payment = ?", req.MerchantID, "Cancel").
		Count(&req.PaymentCancel).Error; err != nil {
		return dto.GetDashboardResponse{}, err
	}
	response.PaymentCancel = req.PaymentCancel

	var productPercentages []dto.ProductPercentage
	if err := r.DB.Raw(`
    SELECT 
        json_data->>'product_id' AS product_id,
        json_data->>'product_name' AS product_name,
        COUNT(*) AS total_transactions,
        SUM(CASE WHEN pos.status_payment = 'Lunas' THEN 1 ELSE 0 END) AS total_lunas,
        (SUM(CASE WHEN pos.status_payment = 'Lunas' THEN 1 ELSE 0 END) * 100.0 / COUNT(*)) AS percentage
    FROM pos,
    LATERAL json_array_elements(pos.product) AS json_data
    WHERE pos.merchant_id = ? 
    AND pos.deleted_at IS NULL
    GROUP BY json_data->>'product_id', json_data->>'product_name'
`, req.MerchantID).Scan(&productPercentages).Error; err != nil {
		return response, err
	}

	response.ProductPercentages = productPercentages

	return response, nil
}
