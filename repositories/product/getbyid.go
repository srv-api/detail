package product

import (
	"fmt"
	"strconv"

	dto "github.com/srv-api/detail/dto"
	"github.com/srv-api/product/entity"
)

func (b *productRepository) GetById(req dto.GetByIdRequest) (*dto.ProductResponse, error) {
	tr := entity.Product{
		ID: req.ID,
	}

	if err := b.DB.Where("id = ?", tr.ID).Take(&tr).Error; err != nil {
		return nil, err
	}

	statusMap := map[int]string{
		1: "active",
		2: "inactive",
	}

	createdStatus, err := strconv.Atoi(fmt.Sprintf("%v", tr.Status))
	if err != nil {
		return &dto.ProductResponse{}, fmt.Errorf("invalid status value: %v", tr.Status)
	}

	statusString, ok := statusMap[createdStatus]
	if !ok {
		return &dto.ProductResponse{}, fmt.Errorf("invalid status value in database")
	}

	response := &dto.ProductResponse{
		ID:          tr.ID,
		UserID:      tr.UserID,
		DetailID:    tr.MerchantID,
		ProductName: tr.ProductName,
		Price:       tr.Price,
		Status:      statusString,
		Description: tr.Description,
		CreatedAt:   dto.Timestamp(tr.CreatedAt),
		CreatedBy:   tr.CreatedBy,
	}

	return response, nil
}
