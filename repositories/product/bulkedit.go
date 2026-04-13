package product

import (
	"time"

	dto "github.com/srv-api/merchant/dto"
	"github.com/srv-api/product/entity"
)

func (b *productRepository) BulkEdit(req dto.BulkEditRequest) (int, error) {
	tx := b.DB.Begin()
	if tx.Error != nil {
		return 0, tx.Error
	}

	var totalUpdated int

	for _, item := range req.Items {
		if item.ID == "" {
			continue
		}

		updateData := map[string]interface{}{}

		if item.SKU != 0 {
			updateData["sku"] = item.SKU
		}
		if item.Barcode != "" {
			updateData["barcode"] = item.Barcode
		}
		if item.MerkID != "" {
			updateData["merk_id"] = item.MerkID
		}
		if item.CategoryID != "" {
			updateData["category_id"] = item.CategoryID
		}
		if item.ProductName != "" {
			updateData["product_name"] = item.ProductName
		}
		if item.Description != "" {
			updateData["description"] = item.Description
		}
		if item.Stock != 0 {
			updateData["stock"] = item.Stock
		}
		if item.MinimalStock != 0 {
			updateData["minimal_stock"] = item.MinimalStock
		}
		if item.Price != 0 {
			updateData["price"] = item.Price
		}
		if item.Status != 0 {
			updateData["status"] = item.Status
		}

		// Selalu update info pengguna & waktu
		updateData["updated_by"] = item.UpdatedBy
		updateData["updated_at"] = time.Now()

		// Jika tidak ada field yang diubah, lewati
		if len(updateData) == 0 {
			continue
		}

		// Update satu produk
		result := tx.Model(&entity.Product{}).Where("id = ?", item.ID).Updates(updateData)
		if result.Error != nil {
			tx.Rollback()
			return 0, result.Error
		}

		totalUpdated += int(result.RowsAffected)
	}

	if err := tx.Commit().Error; err != nil {
		return 0, err
	}

	return totalUpdated, nil
}
