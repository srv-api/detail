package roleuser

import (
	dto "github.com/srv-api/merchant/dto"
	"github.com/srv-api/merchant/entity"
)

func (b *RoleUserRepository) Delete(req dto.DeleteRoleUserRequest) (dto.DeleteResponse, error) {
	tr := dto.GetRoleUserByIdRequest{
		ID: req.ID,
	}

	_, err := b.GetById(tr)
	if err != nil {
		return dto.DeleteResponse{}, err
	}

	// Use GORM BeforeDelete hook to set DeletedBy
	if err := b.DB.Where("id = ?", req.ID).Delete(&entity.RoleUser{}).Error; err != nil {
		return dto.DeleteResponse{}, err
	}

	response := dto.DeleteResponse{
		DeletedBy: req.DeletedBy,
	}

	return response, nil
}
