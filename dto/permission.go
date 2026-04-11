package dto

import (
	"time"
)

type PermissionRequest struct {
	ID         uint      `json:"id"`
	UserID     string    `json:"user_id"`
	MerchantID string    `json:"merchant_id"`
	Label      string    `json:"label"`
	Icon       string    `json:"icon"`
	To         string    `json:"to"`
	CreatedBy  string    `json:"created_by"`
	UpdatedBy  string    `json:"updated_by"`
	DeletedBy  string    `json:"deleted_by"`
	CreatedAt  time.Time `json:"created_at"`
}

type PermissionResponse struct {
	Label string `json:"label"`
	Icon  string `json:"icon"`
	To    string `json:"to"`
}

type GetPermissionByIdRequest struct {
	ID         uint      `json:"id"`
	UserID     string    `json:"user_id"`
	MerchantID string    `json:"merchant_id"`
	Label      string    `json:"label"`
	Icon       string    `json:"icon"`
	To         string    `json:"to"`
	CreatedBy  string    `json:"created_by"`
	UpdatedBy  string    `json:"updated_by"`
	DeletedBy  string    `json:"deleted_by"`
	CreatedAt  time.Time `json:"created_at"`
}

type DeletePermissionRequest struct {
	ID        uint   `param:"id" validate:"required"`
	DeletedBy string `json:"deleted_by"`
}

type DeletePermissionResponse struct {
	ID        uint   `param:"id" validate:"required"`
	DeletedBy string `json:"deleted_by"`
}

type GetPermissionResponse struct {
	Items []PermissionResponse `json:"items"`
}

type PermissionItem struct {
	ID    uint   `json:"id"`
	Label string `json:"label"`
	Icon  string `json:"icon"`
	To    string `json:"to"`
}

type PermissionUpdateRequest struct {
	ID         string `json:"id"`
	Label      string `json:"label"`
	Icon       string `json:"icon"`
	To         string `json:"to"`
	UserID     string `json:"user_id"`
	MerchantID string `json:"merchant_id"`
	UpdatedBy  string `json:"updated_by"`
}

type PermissionUpdateResponse struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id"`
	MerchantID string    `json:"merchant_id"`
	Label      string    `json:"label"`
	Icon       string    `json:"icon"`
	To         string    `json:"to"`
	CreatedBy  string    `json:"created_by"`
	UpdatedBy  string    `json:"updated_by"`
	DeletedBy  string    `json:"deleted_by"`
	CreatedAt  time.Time `json:"created_at"`
}
