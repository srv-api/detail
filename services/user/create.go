package user

import (
	"fmt"

	dto "github.com/srv-api/merchant/dto"
)

func (s *userService) Create(req dto.UserRequest) (dto.UserResponse, error) {
	if req.Status != 1 && req.Status != 2 {
		return dto.UserResponse{}, fmt.Errorf("invalid status: must be 1 (active) or 2 (inactive)")
	}

	create := dto.UserRequest{
		User:        req.User,
		Description: req.Description,
		Status:      req.Status,
		UserID:      req.UserID,
		MerchantID:  req.MerchantID,
		CreatedBy:   req.CreatedBy,
	}

	created, err := s.Repo.Create(create)
	if err != nil {
		return dto.UserResponse{}, err
	}

	statusMap := map[int]string{
		1: "active",
		2: "inactive",
	}

	// Dapatkan string status berdasarkan nilai integer
	statusString, ok := statusMap[create.Status]
	if !ok {
		return dto.UserResponse{}, fmt.Errorf("invalid status value in database")
	}

	response := dto.UserResponse{
		ID:          created.ID,
		UserID:      created.UserID,
		FullName:    created.FullName,
		Description: created.Description,
		Status:      statusString,
		MerchantID:  created.MerchantID,
		CreatedBy:   created.CreatedBy,
	}

	return response, nil
}
