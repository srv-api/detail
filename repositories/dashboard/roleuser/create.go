package roleuser

import (
	"encoding/json"

	dto "github.com/srv-api/detail/dto"
	"github.com/srv-api/detail/entity"
)

func (r *RoleUserRepository) Create(req dto.RoleUserRequest) (dto.RoleUserResponse, error) {

	// Convert array → JSON
	jsonData, err := json.Marshal(req.PermissionID)
	if err != nil {
		return dto.RoleUserResponse{}, err
	}

	create := entity.RoleUser{
		DetailID:     req.DetailID,
		RoleID:       req.RoleID,
		UserID:       req.UserID,
		PermissionID: jsonData, // sudah []byte
	}

	if err := r.DB.Save(&create).Error; err != nil {
		return dto.RoleUserResponse{}, err
	}

	response := dto.RoleUserResponse{
		DetailID:     req.DetailID,
		RoleID:       create.RoleID,
		UserID:       req.UserID,
		PermissionID: req.PermissionID, // array int
	}

	return response, nil
}
