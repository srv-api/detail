package permission

import (
	dto "github.com/srv-api/merchant/dto"
)

func (b *permissionService) Delete(req dto.DeletePermissionRequest) (dto.DeletePermissionResponse, error) {
	transactionBody := dto.DeletePermissionRequest{
		ID:        req.ID,
		DeletedBy: req.DeletedBy,
	}

	_, err := b.Repo.Delete(req)
	if err != nil {
		return dto.DeletePermissionResponse{}, err
	}

	response := dto.DeletePermissionResponse{
		ID:        transactionBody.ID,
		DeletedBy: transactionBody.DeletedBy,
	}

	return response, nil
}
