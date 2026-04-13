package paymentmethod

import (
	"crypto/rand"
	"fmt"
	"time"

	dto "github.com/srv-api/merchant/dto"
	"github.com/srv-api/merchant/entity"
	"golang.org/x/crypto/blake2b"
)

func (s *paymentmethodService) Create(req dto.PaymentMethodRequest) (dto.PaymentMethodResponse, error) {
	created, err := s.Repo.Create(req)
	if err != nil {
		return dto.PaymentMethodResponse{}, err
	}

	if req.FilePath != "" && req.FileName != "" {
		image := entity.UploadedPayment{
			UserID:     req.UserID,
			MerchantID: req.MerchantID,
			PaymentID:  created.ID,
			FileName:   req.FileName,
			FilePath:   req.FilePath,
			CreatedBy:  req.CreatedBy,
		}
		if err := s.Repo.SaveImage(image); err != nil {
			return dto.PaymentMethodResponse{}, err
		}
	}

	resp := dto.PaymentMethodResponse{
		ID:            created.ID,
		UserID:        created.UserID,
		MerchantID:    created.MerchantID,
		PaymentMethod: created.PaymentMethod,
		Status:        created.Status,
		CreatedBy:     created.CreatedBy,
		Category:      created.Category,
	}
	return resp, nil
}

func GenerateSecureID() (string, error) {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_-="

	// Generate a salt
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	// Combine salt and current timestamp for uniqueness
	timestamp := time.Now().UnixNano()
	saltedID := fmt.Sprintf("%x%d", salt, timestamp)

	// Hash the combination using Blake2
	hash, err := blake2b.New512(nil)
	if err != nil {
		return "", err
	}
	hash.Write([]byte(saltedID))
	hashBytes := hash.Sum(nil)

	// Convert hash bytes into a valid string
	var secureID []byte
	for i := 0; i < 12; i++ {
		secureID = append(secureID, chars[hashBytes[i]%byte(len(chars))])
	}

	return string(secureID), nil
}
