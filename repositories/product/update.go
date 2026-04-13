package product

import (
	"github.com/srv-api/merchant/dto"
	"github.com/srv-api/product/entity"
)

func (b *productRepository) Update(req dto.ProductUpdateRequest) (dto.ProductUpdateResponse, error) {
	// Menyiapkan struktur update untuk produk
	updateProduct := entity.Product{
		ProductName: req.ProductName,
		Price:       req.Price,
		Status:      req.Status,
		UpdatedBy:   req.UpdatedBy,
		UserID:      req.UserID,
		Description: req.Description,
		MerchantID:  req.MerchantID,
	}

	// Cek apakah produk ada terlebih dahulu
	var existingProduct entity.Product
	err := b.DB.Where("id = ?", req.ID).First(&existingProduct).Error
	if err != nil {
		return dto.ProductUpdateResponse{}, err
	}

	// Update produk dengan nilai yang baru
	err = b.DB.Model(&existingProduct).Updates(updateProduct).Error
	if err != nil {
		return dto.ProductUpdateResponse{}, err
	}

	// Menyiapkan response setelah pembaruan berhasil
	response := dto.ProductUpdateResponse{
		ProductName: updateProduct.ProductName,
		Price:       updateProduct.Price,
		Status:      updateProduct.Status,
		UpdatedBy:   updateProduct.UpdatedBy,
		UserID:      updateProduct.UserID,
		MerchantID:  updateProduct.MerchantID,
		Description: updateProduct.Description,
	}

	return response, nil
}
