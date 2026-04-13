package product

import (
	"errors"
	"fmt"

	"github.com/srv-api/detail/dto"
	"github.com/srv-api/detail/entity"
	"gorm.io/gorm"
)

func (b *productService) Update(req dto.ProductUpdateRequest) (dto.ProductUpdateResponse, error) {
	// Validasi MerchantDetail
	var merchantDetail entity.UserDetail
	err := b.Repo.CheckMerchantDetail(req.DetailID, &merchantDetail)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.ProductUpdateResponse{}, fmt.Errorf("merchant detail not found for detail_id: %s", req.DetailID)
		}
		return dto.ProductUpdateResponse{}, err
	}

	request := dto.ProductUpdateRequest{
		ProductName: req.ProductName,
		Price:       req.Price,
		Status:      req.Status,
		UpdatedBy:   req.UpdatedBy,
		UserID:      req.UserID,
		Description: req.Description,
		DetailID:    req.DetailID,
	}

	product, err := b.Repo.Update(req)
	if err != nil {
		return product, err
	}

	response := dto.ProductUpdateResponse{
		ProductName: request.ProductName,
		Price:       request.Price,
		Status:      request.Status,
		UpdatedBy:   request.UpdatedBy,
		UserID:      request.UserID,
		Description: request.Description,
		DetailID:    request.DetailID,
	}

	return response, nil
}
