package dto

import (
	"time"
)

type PaymentMethodRequest struct {
	ID            uint      `json:"id"`
	UserID        string    `json:"user_id"`
	MerchantID    string    `json:"merchant_id"`
	PaymentMethod string    `json:"payment_method"`
	PaymentLabel  string    `json:"payment_label"`
	Category      string    `json:"category"`
	Status        int       `json:"status"`
	FileName      string    `json:"file_name"`
	FilePath      string    `json:"file_path"`
	CreatedBy     string    `json:"created_by"`
	UpdatedBy     string    `json:"updated_by"`
	DeletedBy     string    `json:"deleted_by"`
	CreatedAt     time.Time `json:"created_at"`
}

type PaymentMethodResponse struct {
	ID            uint             `json:"id"`
	UserID        string           `json:"user_id"`
	MerchantID    string           `json:"merchant_id"`
	PaymentMethod string           `json:"payment_method"`
	Category      string           `json:"category"`
	PaymentLabel  string           `json:"payment_label"`
	Image         *UploadedPayment `json:"image"` // tambahkan ini
	Status        int              `json:"status"`
	CreatedBy     string           `json:"created_by"`
	UpdatedBy     string           `json:"updated_by"`
	DeletedBy     string           `json:"deleted_by"`
	CreatedAt     time.Time        `json:"created_at"`
}

type UploadedPayment struct {
	ID        uint   `json:"id"`
	PaymentID uint   `json:"payment_id"`
	FileName  string `json:"file_name"`
	FilePath  string `json:"file_path"`
}

type PaymentMethodUpdateRequest struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	MerchantID    string    `json:"merchant_id"`
	PaymentMethod string    `json:"payment_method"`
	PaymentLabel  string    `json:"payment_label"`
	Category      string    `json:"category"`
	Status        int       `json:"status"`
	CreatedBy     string    `json:"created_by"`
	UpdatedBy     string    `json:"updated_by"`
	DeletedBy     string    `json:"deleted_by"`
	CreatedAt     time.Time `json:"created_at"`
}

type PaymentMethodUpdateResponse struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	MerchantID    string    `json:"merchant_id"`
	PaymentMethod string    `json:"payment_method"`
	PaymentLabel  string    `json:"payment_label"`
	Category      string    `json:"category"`
	Status        int       `json:"status"`
	CreatedBy     string    `json:"created_by"`
	UpdatedBy     string    `json:"updated_by"`
	DeletedBy     string    `json:"deleted_by"`
	CreatedAt     time.Time `json:"created_at"`
}

type GetByIdPaymentRequest struct {
	ID uint `param:"id" validate:"required"`
}

type DeletePaymentRequest struct {
	ID        uint   `param:"id" validate:"required"`
	DeletedBy string `json:"deleted_by"`
}

type DeletePaymentResponse struct {
	ID        uint   `param:"id" validate:"required"`
	DeletedBy string `json:"deleted_by"`
}

type GetPaymentploadRequest struct {
	FileName string `param:"file_name" validate:"required"`
}

type GetPaymentUploadResponse struct {
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
}
