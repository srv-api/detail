package dto

import (
	"time"

	"gorm.io/gorm"
)

type DeleteAccountRequest struct {
	ID         uint           `json:"id"`
	UserID     string         `json:"user_id"`
	MerchantID string         `json:"merchant_id"`
	Reason     string         `json:"reason"`
	Email      string         `json:"email"`
	CreatedBy  string         `json:"created_by"`
	UpdatedBy  string         `json:"updated_by"`
	DeletedBy  string         `json:"deleted_by"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}

type DeleteAccountResponse struct {
	ID         uint           `json:"id"`
	UserID     string         `json:"user_id"`
	MerchantID string         `json:"merchant_id"`
	Reason     string         `json:"reason"`
	Email      string         `json:"email"`
	CreatedBy  string         `json:"created_by"`
	UpdatedBy  string         `json:"updated_by"`
	DeletedBy  string         `json:"deleted_by"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}

type DeleteAccountUpdateRequest struct {
	ID         uint           `json:"id"`
	UserID     string         `json:"user_id"`
	MerchantID string         `json:"merchant_id"`
	Reason     string         `json:"reason"`
	Email      string         `json:"email"`
	CreatedBy  string         `json:"created_by"`
	UpdatedBy  string         `json:"updated_by"`
	DeletedBy  string         `json:"deleted_by"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}

type DeleteAccountUpdateResponse struct {
	ID         uint           `json:"id"`
	UserID     string         `json:"user_id"`
	MerchantID string         `json:"merchant_id"`
	Reason     string         `json:"reason"`
	Email      string         `json:"email"`
	CreatedBy  string         `json:"created_by"`
	UpdatedBy  string         `json:"updated_by"`
	DeletedBy  string         `json:"deleted_by"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}

type DeleteDeleteAccountRequest struct {
	ID        uint   `param:"id" validate:"required"`
	DeletedBy string `json:"deleted_by"`
}

type DeleteDeleteAccountResponse struct {
	ID        uint   `param:"id" validate:"required"`
	DeletedBy string `json:"deleted_by"`
}

type GetDeleteAccountByIdRequest struct {
	ID uint `param:"id" validate:"required"`
}
