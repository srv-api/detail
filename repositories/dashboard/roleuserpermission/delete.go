package roleuserpermission

import (
	dto "github.com/srv-api/merchant/dto"
	"github.com/srv-api/merchant/entity"
)

func (b *RoleUserPermissionRepository) Delete(req dto.DeleteRoleUserPermissionRequest) (dto.DeleteResponse, error) {
	tr := dto.GetRoleUserPermissionByIdRequest{
		ID: req.ID,
	}

	_, err := b.GetById(tr)
	if err != nil {
		return dto.DeleteResponse{}, err
	}

	// Use GORM BeforeDelete hook to set DeletedBy
	if err := b.DB.Where("id = ?", req.ID).Delete(&entity.RoleUserPermission{}).Error; err != nil {
		return dto.DeleteResponse{}, err
	}

	response := dto.DeleteResponse{
		DeletedBy: req.DeletedBy,
	}

	return response, nil
}
