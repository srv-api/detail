package user

import (
	res "github.com/srv-api/util/s/response"

	dto "github.com/srv-api/detail/dto"
	util "github.com/srv-api/util/s"
)

func (s *userService) Create(req dto.UserFullRequest) (dto.UserFullResponse, error) {
	// Validate email
	if !util.IsValidEmail(req.Email) {
		return dto.UserFullResponse{}, res.ErrorBuilder(&res.ErrorConstant.RegisterMail, nil)
	}

	req.Whatsapp = util.FormatWhatsappNumber(req.Whatsapp)

	// Encrypt the email
	encryptedEmail, err := util.Encrypt(req.Email)
	if err != nil {
		return dto.UserFullResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	// Encrypt the email
	encryptedWhatsapp, err := util.Encrypt(req.Whatsapp)
	if err != nil {
		return dto.UserFullResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	// Proceed with the signup process
	encryp := util.EncryptPasswordUserDetail(&req)
	if encryp != nil {
		return dto.UserFullResponse{}, encryp
	}

	secureID, err := util.GenerateSecureID()
	if err != nil {
		return dto.UserFullResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	create := dto.UserFullRequest{
		ID:           secureID,
		AccessRoleID: req.AccessRoleID,
		FullName:     req.FullName,
		Whatsapp:     encryptedWhatsapp,
		Email:        encryptedEmail,
		Password:     req.Password,
		Description:  req.Description,
		UserID:       req.UserID,
		DetailID:     req.DetailID,
		CreatedBy:    req.CreatedBy,
	}

	created, err := s.Repo.Create(create)
	if err != nil {
		return dto.UserDetailResponse{}, err
	}

	response := dto.UserDetailResponse{
		AccessRoleID: created.AccessRoleID,
		FullName:     created.FullName,
		Whatsapp:     created.Whatsapp,
		Email:        created.Email,
		Password:     created.Password,
		Description:  created.Description,
		UserID:       created.UserID,
		DetailID:     created.DetailID,
		CreatedBy:    created.CreatedBy,
	}

	return response, nil
}
