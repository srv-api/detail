package deleteaccount

import (
	dto "github.com/srv-api/merchant/dto"
	"github.com/srv-api/merchant/entity"
)

func (b *deleteaccountRepository) Delete(req dto.DeleteDeleteAccountRequest) (dto.DeleteDeleteAccountResponse, error) {
	tr := dto.GetDeleteAccountByIdRequest{
		ID: req.ID,
	}

	_, err := b.GetById(tr)
	if err != nil {
		return dto.DeleteDeleteAccountResponse{}, err
	}

	// Use GORM BeforeDelete hook to set DeletedBy
	if err := b.DB.Where("id = ?", req.ID).Delete(&entity.DeleteAccount{}).Error; err != nil {
		return dto.DeleteDeleteAccountResponse{}, err
	}

	response := dto.DeleteDeleteAccountResponse{
		DeletedBy: req.DeletedBy,
	}

	return response, nil
}
