package user

import (
	"crypto/rand"
	"fmt"

	"github.com/srv-api/auth/entity"
	dto "github.com/srv-api/merchant/dto"
)

func (r *userRepository) Create(req dto.UserMerchantRequest) (dto.UserMerchantResponse, error) {

	// Create the new user entry
	create := entity.AccessDoor{
		ID:           req.ID,
		AccessRoleID: req.AccessRoleID,
		FullName:     req.FullName,
		Whatsapp:     req.Whatsapp,
		Email:        req.Email,
		Password:     req.Password,
		MerchantID:   req.MerchantID,
	}

	// Save the new user to the database
	if err := r.DB.Save(&create).Error; err != nil {
		return dto.UserMerchantResponse{}, err
	}

	// Build the response for the created user
	response := dto.UserMerchantResponse{
		ID:           create.ID,
		AccessRoleID: create.AccessRoleID,
		FullName:     create.FullName,
		Whatsapp:     create.Whatsapp,
		Email:        create.Email,
		Password:     create.Password,
		MerchantID:   create.MerchantID,
	}

	return response, nil
}

// Function to generate the user ID
func generateUserMerchantID(prefix string, autoIncrement int) (string, error) {
	// Format auto-increment value as a 5-digit string
	autoIncStr := fmt.Sprintf("%05d", autoIncrement)

	// Generate a secure random part of the user ID
	securePart, err := generateSecurePart()
	if err != nil {
		return "", err
	}

	// Combine the prefix, secure random part, and auto-increment value to form the final user ID
	return fmt.Sprintf("%s%s%s", prefix, securePart, autoIncStr), nil
}

// Function to generate a secure random part for the user ID
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
