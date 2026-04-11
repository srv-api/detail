package product

import (
	dto "github.com/srv-api/merchant/dto"
)

func (b *productService) BulkEdit(req dto.BulkEditRequest) (dto.BulkEditResponse, error) {
	count, err := b.Repo.BulkEdit(req)
	if err != nil {
		return dto.BulkEditResponse{}, err
	}

	return dto.BulkEditResponse{
		Count:     count,
		UpdatedBy: req.UpdatedBy,
	}, nil
}
