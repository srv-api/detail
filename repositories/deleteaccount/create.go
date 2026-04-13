package deleteaccount

import (
	"crypto/rand"
	"fmt"

	dto "github.com/srv-api/merchant/dto"
	"github.com/srv-api/merchant/entity"
)

func (r *deleteaccountRepository) Create(req dto.DeleteAccountRequest) (dto.DeleteAccountResponse, error) {
	// Insert or update the auto_increment value based on merchant_id
	var autoIncrement int
	err := r.DB.Raw(`
		INSERT INTO merchant_auto_increments (merchant_id, next_increment)
		VALUES (?, 1)
		ON CONFLICT (merchant_id) DO UPDATE
		SET next_increment = merchant_auto_increments.next_increment + 1
		RETURNING next_increment - 1;
	`, req.MerchantID).Scan(&autoIncrement).Error

	if err != nil {
		return dto.DeleteAccountResponse{}, err
	}

	// Create the new deleteaccount entry
	create := entity.DeleteAccount{
		ID:         req.ID,
		Email:      req.Email,
		Reason:     req.Reason,
		UserID:     req.UserID,
		MerchantID: req.MerchantID,
		CreatedBy:  req.CreatedBy,
	}

	// Save the new deleteaccount to the database
	if err := r.DB.Save(&create).Error; err != nil {
		return dto.DeleteAccountResponse{}, err
	}

	// Build the response for the created deleteaccount
	response := dto.DeleteAccountResponse{
		ID:         create.ID,
		UserID:     create.UserID,
		MerchantID: create.MerchantID,
		Email:      create.Email,
		Reason:     create.Reason,
		CreatedBy:  create.CreatedBy,
	}

	return response, nil
}

// Function to generate the deleteaccount ID
func generateDeleteAccountID(prefix string, autoIncrement int) (string, error) {
	// Format auto-increment value as a 5-digit string
	autoIncStr := fmt.Sprintf("%05d", autoIncrement)

	// Generate a secure random part of the deleteaccount ID
	securePart, err := generateSecurePart()
	if err != nil {
		return "", err
	}

	// Combine the prefix, secure random part, and auto-increment value to form the final deleteaccount ID
	return fmt.Sprintf("%s%s%s", prefix, securePart, autoIncStr), nil
}

// Function to generate a secure random part for the deleteaccount ID
func generateSecurePart() (string, error) {
	const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz"

	// Generate a random string of length 12
	securePart := make([]byte, 12)
	_, err := rand.Read(securePart)
	if err != nil {
		return "", err
	}

	// Map each byte to a character from the chars string
	for i := range securePart {
		securePart[i] = chars[securePart[i]%byte(len(chars))]
	}

	return string(securePart), nil
}
