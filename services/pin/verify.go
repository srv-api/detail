package pin

import (
	"errors"

	dto "github.com/srv-api/merchant/dto"
	"golang.org/x/crypto/bcrypt"
)

func (s *pinService) VerifyPIN(req dto.VerifyPinRequest) (*dto.VerifyPinResponse, error) {
	user, err := s.Repo.Verify(req)
	if err != nil || user == nil {
		return nil, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Pin), []byte(req.PIN))
	if err != nil {
		return &dto.VerifyPinResponse{
			IsValid: false,
			Message: "PIN is incorrect",
		}, nil
	}

	return &dto.VerifyPinResponse{
		IsValid: true,
		Message: "PIN is correct",
	}, nil
}
