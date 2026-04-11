package pin

import (
	dto "github.com/srv-api/merchant/dto"
	"github.com/srv-api/merchant/entity"
)

func (b *pinRepository) Delete(req dto.DeletePinRequest) (dto.DeletePinResponse, error) {
	tr := dto.GetByIdPinRequest{
		ID: req.ID,
	}

	_, err := b.GetById(tr)
	if err != nil {
		return dto.DeletePinResponse{}, err
	}

	// Use GORM BeforeDelete hook to set DeletedBy
	if err := b.DB.Where("id = ?", req.ID).Delete(&entity.Pin{}).Error; err != nil {
		return dto.DeletePinResponse{}, err
	}

	response := dto.DeletePinResponse{
		DeletedBy: req.DeletedBy,
	}

	return response, nil
}
