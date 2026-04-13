package product

import (
	"crypto/rand"
	"fmt"
	"strconv"

	dto "github.com/srv-api/detail/dto"
	"github.com/srv-api/detail/entity"
	product "github.com/srv-api/product/entity"
	res "github.com/srv-api/util/s/response"
)

func (r *productRepository) Create(req dto.ProductRequest) (dto.ProductResponse, error) {
	// Insert or update the auto_increment value based on detail_id
	var autoIncrement int
	err := r.DB.Raw(`
		INSERT INTO merchant_auto_increments (detail_id, next_increment)
		VALUES (?, 1)
		ON CONFLICT (detail_id) DO UPDATE
		SET next_increment = merchant_auto_increments.next_increment + 1
		RETURNING next_increment - 1;
	`, req.DetailID).Scan(&autoIncrement).Error

	if err != nil {
		return dto.ProductResponse{}, err
	}

	// Generate Product ID with prefix and auto increment value
	prefix := "p="
	secureID, err := generateProductID(prefix, autoIncrement)
	if err != nil {
		return dto.ProductResponse{}, err
	}

	// Create the new product entry
	create := product.Product{
		ID:          secureID,
		ProductName: req.ProductName,
		Price:       req.Price,
		Status:      req.Status,
		UserID:      req.UserID,
		MerchantID:  req.DetailID,
		CreatedBy:   req.CreatedBy,
		Description: req.Description,
	}

	// Save the new product to the database
	if err := r.DB.Save(&create).Error; err != nil {
		return dto.ProductResponse{}, err
	}

	// Map the status from integer to string
	statusMap := map[int]string{
		1: "active",
		2: "inactive",
	}

	createdStatus, err := strconv.Atoi(fmt.Sprintf("%v", create.Status))
	if err != nil {
		return dto.ProductResponse{}, fmt.Errorf("invalid status value: %v", create.Status)
	}

	statusString, ok := statusMap[createdStatus]
	if !ok {
		return dto.ProductResponse{}, fmt.Errorf("invalid status value in database")
	}

	// Build the response for the created product
	response := dto.ProductResponse{
		ID:          create.ID,
		UserID:      create.UserID,
		DetailID:    create.MerchantID,
		ProductName: create.ProductName,
		Description: create.Description,
		Price:       create.Price,
		Status:      statusString,
		CreatedBy:   create.CreatedBy,
	}

	return response, nil
}

// Function to generate the product ID
func generateProductID(prefix string, autoIncrement int) (string, error) {
	// Format auto-increment value as a 5-digit string
	autoIncStr := fmt.Sprintf("%05d", autoIncrement)

	// Generate a secure random part of the product ID
	securePart, err := generateSecurePart()
	if err != nil {
		return "", err
	}

	// Combine the prefix, secure random part, and auto-increment value to form the final product ID
	return fmt.Sprintf("%s%s%s", prefix, securePart, autoIncStr), nil
}

// Function to generate a secure random part for the product ID
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

func (r *productRepository) CheckMerchantDetail(DetailID string, merchantDetail *entity.UserDetail) error {
	err := r.DB.Where("id = ?", DetailID).First(merchantDetail).Error
	if err != nil {
		return err
	}

	// Periksa apakah semua kolom penting sudah terisi
	if merchantDetail.Latitude == 0 || merchantDetail.Latitude == 0 {
		return res.ErrorBuilder(&res.ErrorConstant.DetailNoProvide, nil)
	}

	return nil
}
