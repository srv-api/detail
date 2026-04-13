package dto

import (
	"time"

	"gorm.io/gorm"
)

type PrinterRequest struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	MerchantID  string    `json:"merchant_id"`
	PrinterName string    `json:"printer_name"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `json:"updated_by"`
	DeletedBy   string    `json:"deleted_by"`
	CreatedAt   time.Time `json:"created_at"`
}

type DeletePrinterRequest struct {
	ID        string `param:"id" validate:"required"`
	DeletedBy string `json:"deleted_by"`
}

type DeletePrinterResponse struct {
	ID        string `param:"id" validate:"required"`
	DeletedBy string `json:"deleted_by"`
}

type PrinterResponse struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	MerchantID  string    `json:"merchant_id"`
	PrinterName string    `json:"printer_name"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `json:"updated_by"`
	DeletedBy   string    `json:"deleted_by"`
	CreatedAt   time.Time `json:"created_at"`
}

type GetPrinterByIdRequest struct {
	ID string `query:"id" validate:"required, id"`
}

type GetPrinterRequest struct {
	ID          string         `json:"id"`
	UserID      string         `json:"user_id"`
	MerchantID  string         `json:"merchant_id"`
	PrinterName string         `json:"printer_name"`
	UpdatedBy   string         `json:"update_by"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

type GetPrinterResponse struct {
	ID          string         `json:"id"`
	UserID      string         `json:"user_id"`
	MerchantID  string         `json:"merchant_id"`
	PrinterName string         `json:"printer_name"`
	UpdatedBy   string         `json:"update_by"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

type UpdatePrinterRequest struct {
	ID          string    `query:"id" validate:"required, id"`
	UserID      string    `json:"user_id"`
	PrinterName string    `json:"printer_name"`
	UpdatedBy   string    `json:"update_by"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UpdatePrinterResponse struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	PrinterName string    `json:"printer_name"`
	UpdatedBy   string    `json:"update_by"`
	UpdatedAt   time.Time `json:"updated_at"`
}
