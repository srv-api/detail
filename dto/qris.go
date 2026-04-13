package dto

import (
	"mime/multipart"
	"time"
)

type CoQrisRequest struct {
	ID         string `json:"id"`
	UserID     string `json:"user_id"`
	MerchantID string `json:"merchant_id"`
	QrisName   string `json:"qris_name"`
	Link       string `json:"link"`
	File       *multipart.FileHeader
	Status     string    `json:"status"`
	CreatedBy  string    `json:"created_by"`
	UpdatedBy  string    `json:"updated_by"`
	DeletedBy  string    `json:"deleted_by"`
	CreatedAt  time.Time `json:"created_at"`
}

type CoQrisResponse struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id"`
	MerchantID string    `json:"merchant_id"`
	QrisName   string    `json:"qris_name"`
	Link       string    `json:"link"`
	FilePath   string    `json:"file_path"`
	Status     string    `json:"status"`
	CreatedBy  string    `json:"created_by"`
	UpdatedBy  string    `json:"updated_by"`
	DeletedBy  string    `json:"deleted_by"`
	CreatedAt  time.Time `json:"created_at"`
}
