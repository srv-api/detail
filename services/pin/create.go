package pin

import (
	dto "github.com/srv-api/merchant/dto"
	util "github.com/srv-api/util/s"
	"golang.org/x/crypto/bcrypt"
)

func (s *pinService) Create(req dto.PinRequest) (dto.PinResponse, error) {
	// Hash the PIN
	hashedPin, err := bcrypt.GenerateFromPassword([]byte(req.Pin), bcrypt.DefaultCost)
	if err != nil {
		return dto.PinResponse{}, err
	}

	// Proses pembuatan data Pin
	create := dto.PinRequest{
		ID:         util.GenerateRandomString(),
		UserID:     req.UserID,
		MerchantID: req.MerchantID,
		Pin:        string(hashedPin),
		CreatedBy:  req.CreatedBy,
	}

	created, err := s.Repo.Create(create)
	if err != nil {
		return dto.PinResponse{}, err
	}

	response := dto.PinResponse{
		ID:         created.ID,
		MerchantID: created.MerchantID,
		Pin:        "",
		UserID:     created.UserID,
		CreatedBy:  created.CreatedBy,
	}

	return response, nil
}
