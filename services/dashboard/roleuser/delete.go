package roleuser

import (
	dto "github.com/srv-api/merchant/dto"
)

func (b *roleuserService) Delete(req dto.DeleteRoleUserRequest) (dto.DeleteRoleUserResponse, error) {
	transactionBody := dto.DeleteRoleUserRequest{
		ID:        req.ID,
		DeletedBy: req.DeletedBy,
	}

	_, err := b.Repo.Delete(req)
	if err != nil {
		return dto.DeleteRoleUserResponse{}, err
	}

	response := dto.DeleteRoleUserResponse{
		ID:        transactionBody.ID,
		DeletedBy: transactionBody.DeletedBy,
	}

	return response, nil
}
