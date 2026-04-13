package role

import (
	dto "github.com/srv-api/merchant/dto"
)

func (b *roleService) Delete(req dto.DeleteRoleRequest) (dto.DeleteRoleResponse, error) {
	transactionBody := dto.DeleteRoleRequest{
		ID:        req.ID,
		DeletedBy: req.DeletedBy,
	}

	_, err := b.Repo.Delete(req)
	if err != nil {
		return dto.DeleteRoleResponse{}, err
	}

	response := dto.DeleteRoleResponse{
		ID:        transactionBody.ID,
		DeletedBy: transactionBody.DeletedBy,
	}

	return response, nil
}
