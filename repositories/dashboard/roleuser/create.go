package roleuser

import (
	"encoding/json"

	dto "github.com/srv-api/merchant/dto"
	"github.com/srv-api/merchant/entity"
)

func (r *RoleUserRepository) Create(req dto.RoleUserRequest) (dto.RoleUserResponse, error) {

	// Convert array → JSON
	jsonData, err := json.Marshal(req.PermissionID)
	if err != nil {
		return dto.RoleUserResponse{}, err
	}

	create := entity.RoleUser{
		MerchantID:   req.MerchantID,
		RoleID:       req.RoleID,
		UserID:       req.UserID,
		PermissionID: jsonData, // sudah []byte
	}

	if err := r.DB.Save(&create).Error; err != nil {
		return dto.RoleUserResponse{}, err
	}

	response := dto.RoleUserResponse{
		MerchantID:   req.MerchantID,
		RoleID:       create.RoleID,
		UserID:       req.UserID,
		PermissionID: req.PermissionID, // array int
	}

	return response, nil
}
