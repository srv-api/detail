package dto

import (
	"mime/multipart"
	"time"
)

type ProductRequest struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	DetailID    string `json:"detail_id"`
	ProductName string `json:"product_name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Status      int    `json:"status"`
	CreatedBy   string `json:"created_by"`
}

type BulkEditAvail struct {
	ID           string `json:"id"`
	SKU          uint64 `json:"sku,omitempty"`
	Barcode      string `json:"barcode,omitempty"`
	MerkID       string `json:"merk_id,omitempty"`
	CategoryID   string `json:"category_id,omitempty"`
	ProductName  string `json:"product_name,omitempty"`
	Description  string `json:"description,omitempty"`
	Stock        int    `json:"stock,omitempty"`
	MinimalStock int    `json:"minimal_stock,omitempty"`
	Price        int    `json:"price,omitempty"`
	Status       int    `json:"status,omitempty"`
	UpdatedBy    string `json:"updated_by"`
	UpdatedAt    string `json:"updated_at"`
}

type ProductResponse struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	DetailID    string    `json:"detail_id"`
	ProductName string    `json:"product_name"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	Status      string    `json:"status"`
	CreatedBy   string    `json:"created_by"`
	CreatedAt   Timestamp `json:"created_at"`
}

type ProductGetByIdRequest struct {
	ID string `param:"id" validate:"required"`
}

type ProductDeleteRequest struct {
	ID        string `param:"id" validate:"required"`
	DeletedBy string `json:"deleted_by"`
}

type ProductDeleteResponse struct {
	ID        string `json:"id"`
	DeletedBy string `json:"deleted_by"`
}

type ProductBulkDeleteRequest struct {
	ID        []string `json:"id"`
	DeletedBy string   `json:"deleted_by"`
}

type ProductBulkDeleteResponse struct {
	ID        []string `json:"id"`
	DeletedBy string   `json:"deleted_by"`
	Count     int      `json:"count"`
}

type ProductUpdateRequest struct {
	ID          string `json:"id"`
	ProductName string `json:"product_name"`
	Price       int    `json:"price"`
	Status      int    `json:"status"`
	UpdatedBy   string `json:"updated_by"`
	UserID      string `json:"user_id"`
	DetailID    string `json:"detail_id"`
	Description string `json:"description"`
}

type ProductUploadRequest struct {
	ID          string `json:"id"`
	File        *multipart.FileHeader
	CreatedBy   string `json:"created_by"`
	UpdatedBy   string `json:"updated_by"`
	ProductID   string `json:"product_id"`
	UserID      string `json:"user_id"`
	DetailID    string `json:"detail_id"`
	Destination string `json:"destination"`
}

type ProductUploadResponse struct {
	FilePath string `json:"file_path"`
}

type GetProductUploadRequest struct {
	FileName string `param:"file_name" validate:"required"`
}

type GetProductUploadResponse struct {
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
}

type ProductUpdateResponse struct {
	ID          string `json:"id"`
	ProductName string `json:"product_name"`
	Price       int    `json:"price"`
	Status      int    `json:"status"`
	UpdatedBy   string `json:"updated_by"`
	UserID      string `json:"user_id"`
	DetailID    string `json:"detail_id"`
	Description string `json:"description"`
}

// Formatter untuk timestamp dengan nama bulan
type Timestamp time.Time

// Format waktu: 19 December 2024, 09:28:17
const timeFormat = "02 January 2006, 15:04:05"

// MarshalJSON untuk memformat waktu
func (t Timestamp) MarshalJSON() ([]byte, error) {
	// Konversi waktu ke zona waktu lokal
	localTime := time.Time(t).Local()
	formattedTime := "\"" + localTime.Format(timeFormat) + "\""
	return []byte(formattedTime), nil
}

// UnmarshalJSON untuk parsing waktu dari JSON (opsional)
func (t *Timestamp) UnmarshalJSON(data []byte) error {
	parsedTime, err := time.Parse("\""+timeFormat+"\"", string(data))
	if err != nil {
		return err
	}
	*t = Timestamp(parsedTime)
	return nil
}

// ToTime untuk mengonversi kembali ke time.Time
func (t Timestamp) ToTime() time.Time {
	return time.Time(t)
}

type ImportResultDTO struct {
	Total    int      `json:"total"`
	Success  int      `json:"success"`
	Failed   int      `json:"failed"`
	Failures []string `json:"failures"`
}

type ExportFilter struct {
	From           string `json:"from"`
	To             string `json:"to"`
	UserID         string `json:"user_id"`
	Barcode        string `json:"barcode"`
	DetailIDstring `json:"detail_id"`
}

type BulkEditRequest struct {
	ID        []string        `json:"id"`
	Items     []BulkEditAvail `json:"items"`
	UpdatedBy string          `json:"updated_by"`
}

type BulkEditResponse struct {
	ID        []string `json:"id"`
	UpdatedBy string   `json:"updated_by"`
	Count     int      `json:"count"`
}
