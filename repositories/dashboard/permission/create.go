package permission

import (
	"crypto/rand"
	"fmt"

	dto "github.com/srv-api/merchant/dto"
	"github.com/srv-api/merchant/entity"
)

func (r *PermissionRepository) Create(req dto.PermissionRequest) (dto.PermissionResponse, error) {

	// Create the new Permission entry
	create := entity.Permission{
		Label:      req.Label,
		Icon:       req.Icon,
		UserID:     req.UserID,
		MerchantID: req.MerchantID,
		CreatedBy:  req.CreatedBy,
		To:         req.To,
	}

	// Save the new Permission to the database
	if err := r.DB.Save(&create).Error; err != nil {
		return dto.PermissionResponse{}, err
	}

	// Build the response for the created Permission
	response := dto.PermissionResponse{
		Label: create.Label,
		Icon:  create.Icon,
		To:    create.To,
	}

	return response, nil
}

// Function to generate the Permission ID
func generatePermissionID(prefix string, autoIncrement int) (string, error) {
	// Format auto-increment value as a 5-digit string
	autoIncStr := fmt.Sprintf("%05d", autoIncrement)

	// Generate a secure random part of the Permission ID
	securePart, err := generateSecurePart()
	if err != nil {
		return "", err
	}

	// Combine the prefix, secure random part, and auto-increment value to form the final Permission ID
	return fmt.Sprintf("%s%s%s", prefix, securePart, autoIncStr), nil
}

// Function to generate a secure random part for the Permission ID
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
