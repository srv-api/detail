package deleteaccount

import (
	dto "github.com/srv-api/merchant/dto"
	"github.com/srv-api/merchant/entity"
)

func (b *deleteaccountRepository) GetById(req dto.GetDeleteAccountByIdRequest) (*dto.DeleteAccountResponse, error) {
	tr := entity.DeleteAccount{
		ID: req.ID,
	}

	if err := b.DB.Where("id = ?", tr.ID).Take(&tr).Error; err != nil {
		return nil, err
	}

	response := &dto.DeleteAccountResponse{
		Email:  tr.Email,
		Reason: tr.Reason,
	}

	return response, nil
}
