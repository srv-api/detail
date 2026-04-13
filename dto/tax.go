package dto

import (
	"time"
)

type TaxRequest struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	MerchantID    string    `json:"merchant_id"`
	Tax           string    `json:"tax"`
	TaxPercentage uint      `json:"tax_percentage"`
	Status        int       `json:"status"`
	Description   string    `json:"description"`
	CreatedBy     string    `json:"created_by"`
	UpdatedBy     string    `json:"updated_by"`
	DeletedBy     string    `json:"deleted_by"`
	CreatedAt     time.Time `json:"created_at"`
}

type TaxResponse struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	MerchantID    string    `json:"merchant_id"`
	Tax           string    `json:"tax"`
	TaxPercentage uint      `json:"tax_percentage"`
	Status        string    `json:"status"`
	Description   string    `json:"description"`
	CreatedBy     string    `json:"created_by"`
	UpdatedBy     string    `json:"updated_by"`
	DeletedBy     string    `json:"deleted_by"`
	CreatedAt     time.Time `json:"created_at"`
}
type TaxUpdateRequest struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	MerchantID    string    `json:"merchant_id"`
	Tax           string    `json:"tax"`
	TaxPercentage uint      `json:"tax_percentage"`
	Status        int       `json:"status"`
	Description   string    `json:"description"`
	CreatedBy     string    `json:"created_by"`
	UpdatedBy     string    `json:"updated_by"`
	DeletedBy     string    `json:"deleted_by"`
	CreatedAt     time.Time `json:"created_at"`
}

type TaxUpdateResponse struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	MerchantID    string    `json:"merchant_id"`
	Tax           string    `json:"tax"`
	TaxPercentage uint      `json:"tax_percentage"`
	Status        int       `json:"status"`
	Description   string    `json:"description"`
	CreatedBy     string    `json:"created_by"`
	UpdatedBy     string    `json:"updated_by"`
	DeletedBy     string    `json:"deleted_by"`
	CreatedAt     time.Time `json:"created_at"`
}
