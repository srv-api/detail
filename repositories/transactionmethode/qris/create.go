package qris

import (
	dto "github.com/srv-api/merchant/dto"
	"github.com/srv-api/merchant/entity"
	util "github.com/srv-api/util/s"
)

func (r *qrisRepository) Create(req dto.CoQrisRequest) (dto.CoQrisResponse, error) {

	// Create the new qris entry
	create := entity.Qris{
		ID:         util.GenerateRandomString(),
		QrisName:   req.QrisName,
		Link:       req.Link,
		Status:     req.Status,
		UserID:     req.UserID,
		MerchantID: req.MerchantID,
		CreatedBy:  req.CreatedBy,
	}

	// Save the new qris to the database
	if err := r.DB.Save(&create).Error; err != nil {
		return dto.CoQrisResponse{}, err
	}

	// Map the status from integer to string

	// Build the response for the created qris
	response := dto.CoQrisResponse{
		ID:         create.ID,
		QrisName:   create.QrisName,
		Link:       create.Link,
		UserID:     create.UserID,
		MerchantID: create.MerchantID,
		Status:     create.Status,
		CreatedBy:  create.CreatedBy,
	}

	return response, nil
}

// Function to generate the qris ID
