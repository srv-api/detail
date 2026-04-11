package pin

import (
	dto "github.com/srv-api/merchant/dto"
)

func (b *pinService) BulkDelete(req dto.BulkDeletePinRequest) (dto.BulkDeletePinResponse, error) {
	transactionBody := dto.BulkDeletePinRequest{
		ID:        req.ID,
		DeletedBy: req.DeletedBy,
	}

	count, err := b.Repo.BulkDelete(req)
	if err != nil {
		return dto.BulkDeletePinResponse{}, err
	}

	response := dto.BulkDeletePinResponse{
		ID:        transactionBody.ID,
		DeletedBy: transactionBody.DeletedBy,
		Count:     count, // Menyimpan jumlah yang dihapus
	}

	return response, nil
}
