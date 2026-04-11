package product

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	dto "github.com/srv-api/merchant/dto"
	"github.com/srv-api/product/entity"
)

func (r *productRepository) SaveFile(req dto.ProductUploadRequest) (dto.ProductUploadResponse, error) {
	// Save the file physically
	src, err := req.File.Open()
	if err != nil {
		return dto.ProductUploadResponse{}, fmt.Errorf("failed to open source file: %w", err)
	}
	defer src.Close()

	// Ensure the destination directory exists
	dir := filepath.Dir(req.Destination)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return dto.ProductUploadResponse{}, fmt.Errorf("failed to create directory: %w", err)
	}

	// Create the destination file
	dst, err := os.Create(req.Destination)
	if err != nil {
		return dto.ProductUploadResponse{}, fmt.Errorf("failed to create destination file: %w", err)
	}
	defer dst.Close()

	// Copy file content
	if _, err := io.Copy(dst, src); err != nil {
		return dto.ProductUploadResponse{}, fmt.Errorf("failed to copy file content: %w", err)
	}

	// Prepare metadata for database
	fileRecord := entity.UploadedFile{
		FileName:   filepath.Base(req.Destination),
		FilePath:   req.Destination,
		MerchantID: req.MerchantID,
		UserID:     req.UserID,
		ProductID:  req.ID,
		CreatedBy:  req.CreatedBy,
	}

	// Save metadata to database
	if err := r.DB.Create(&fileRecord).Error; err != nil {
		return dto.ProductUploadResponse{}, fmt.Errorf("failed to save file metadata to database: %w", err)
	}

	// Return response
	return dto.ProductUploadResponse{
		FilePath: req.Destination,
	}, nil
}
