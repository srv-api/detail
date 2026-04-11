package deleteaccount

import "github.com/srv-api/merchant/dto"

func (b *deleteaccountService) Update(req dto.DeleteAccountUpdateRequest) (dto.DeleteAccountUpdateResponse, error) {
	request := dto.DeleteAccountUpdateRequest{
		Email:      req.Email,
		Reason:     req.Reason,
		UpdatedBy:  req.UpdatedBy,
		UserID:     req.UserID,
		MerchantID: req.MerchantID,
	}

	deleteaccount, err := b.Repo.Update(req)
	if err != nil {
		return deleteaccount, err
	}

	response := dto.DeleteAccountUpdateResponse{
		Email:      request.Email,
		Reason:     request.Reason,
		UpdatedBy:  request.UpdatedBy,
		UserID:     request.UserID,
		MerchantID: request.MerchantID,
	}

	return response, nil
}
