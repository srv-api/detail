package user

import "github.com/srv-api/merchant/dto"

func (b *userService) Update(req dto.UserUpdateRequest) (dto.UserUpdateResponse, error) {
	request := dto.UserUpdateRequest{
		FullName:    req.FullName,
		Status:      req.Status,
		UpdatedBy:   req.UpdatedBy,
		Description: req.Description,
	}

	user, err := b.Repo.Update(req)
	if err != nil {
		return user, err
	}

	response := dto.UserUpdateResponse{
		FullName:    request.FullName,
		Status:      request.Status,
		UpdatedBy:   request.UpdatedBy,
		Description: request.Description,
	}

	return response, nil
}
