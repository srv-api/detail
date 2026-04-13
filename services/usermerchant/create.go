package user

import (
	res "github.com/srv-api/util/s/response"

	dto "github.com/srv-api/merchant/dto"
	util "github.com/srv-api/util/s"
)

func (s *userService) Create(req dto.UserMerchantRequest) (dto.UserMerchantResponse, error) {
	// Validate email
	if !util.IsValidEmail(req.Email) {
		return dto.UserMerchantResponse{}, res.ErrorBuilder(&res.ErrorConstant.RegisterMail, nil)
	}

	req.Whatsapp = util.FormatWhatsappNumber(req.Whatsapp)

	// Encrypt the email
	encryptedEmail, err := util.Encrypt(req.Email)
	if err != nil {
		return dto.UserMerchantResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	// Encrypt the email
	encryptedWhatsapp, err := util.Encrypt(req.Whatsapp)
	if err != nil {
		return dto.UserMerchantResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	// Proceed with the signup process
	encryp := util.EncryptPasswordUserMerchant(&req)
	if encryp != nil {
		return dto.UserMerchantResponse{}, encryp
	}

	secureID, err := util.GenerateSecureID()
	if err != nil {
		return dto.UserMerchantResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	create := dto.UserMerchantRequest{
		ID:           secureID,
		AccessRoleID: req.AccessRoleID,
		FullName:     req.FullName,
		Whatsapp:     encryptedWhatsapp,
		Email:        encryptedEmail,
		Password:     req.Password,
		Description:  req.Description,
		UserID:       req.UserID,
		MerchantID:   req.MerchantID,
		CreatedBy:    req.CreatedBy,
	}

	created, err := s.Repo.Create(create)
	if err != nil {
		return dto.UserMerchantResponse{}, err
	}

	response := dto.UserMerchantResponse{
		AccessRoleID: created.AccessRoleID,
		FullName:     created.FullName,
		Whatsapp:     created.Whatsapp,
		Email:        created.Email,
		Password:     created.Password,
		Description:  created.Description,
		UserID:       created.UserID,
		MerchantID:   created.MerchantID,
		CreatedBy:    created.CreatedBy,
	}

	return response, nil
}
