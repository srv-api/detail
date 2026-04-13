package user

import (
	"crypto/rand"
	"fmt"
	"strconv"

	"github.com/srv-api/auth/entity"
	dto "github.com/srv-api/merchant/dto"
)

func (r *userRepository) Create(req dto.UserRequest) (dto.UserResponse, error) {
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
		return dto.UserResponse{}, err
	}

	// Generate User ID with prefix and auto increment value
	prefix := "p="
	secureID, err := generateUserID(prefix, autoIncrement)
	if err != nil {
		return dto.UserResponse{}, err
	}

	// Create the new user entry
	create := entity.AccessDoor{
		ID:       secureID,
		FullName: req.FullName,
	}

	// Save the new user to the database
	if err := r.DB.Save(&create).Error; err != nil {
		return dto.UserResponse{}, err
	}

	// Map the status from integer to string
	statusMap := map[int]string{
		1: "active",
		2: "inactive",
	}

	createdStatus, err := strconv.Atoi(fmt.Sprintf("%v", create.Verified.StatusAccount))
	if err != nil {
		return dto.UserResponse{}, fmt.Errorf("invalid status value: %v", create.Verified.StatusAccount)
	}

	statusString, ok := statusMap[createdStatus]
	if !ok {
		return dto.UserResponse{}, fmt.Errorf("invalid status value in database")
	}

	// Build the response for the created user
	response := dto.UserResponse{
		ID:       create.ID,
		FullName: create.FullName,
		Status:   statusString,
	}

	return response, nil
}

// Function to generate the user ID
func generateUserID(prefix string, autoIncrement int) (string, error) {
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
