package dto

import "time"

type RoleRequest struct {
	ID         string `json:"id"`
	Role       string `json:"role`
	UserID     string `json:"user_id"`
	MerchantID string `json:"merchant_id"`
	CreatedBy  string `json:"created_by"`
}

type RoleResponse struct {
	ID   string `json:"id"`
	Role string `json:"role`
}

type RoleUpdateRequest struct {
	ID         string `json:"id"`
	Role       string `json:"role"`
	UserID     string `json:"user_id"`
	MerchantID string `json:"merchant_id"`
	CreatedBy  string `json:"created_by"`
	UpdatedBy  string `json:"updated_by"`
}

type RoleUpdateResponse struct {
	ID         string `json:"id"`
	Role       string `json:"role"`
	UserID     string `json:"user_id"`
	MerchantID string `json:"merchant_id"`
	CreatedBy  string `json:"created_by"`
	UpdatedBy  string `json:"updated_by"`
}

type DeleteRoleRequest struct {
	ID        string `param:"id" validate:"required"`
	DeletedBy string `json:"deleted_by"`
}

type DeleteRoleResponse struct {
	ID        string `param:"id" validate:"required"`
	DeletedBy string `json:"deleted_by"`
}

type GetRoleRequest struct {
	UserID     string `json:"user_id"`
	MerchantID string `json:"merchant_id"`
}

type GetRoleResponse struct {
	Role []RoleResponse `json:"roles"`
}

type GetRoleByIdRequest struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id"`
	MerchantID string    `json:"merchant_id"`
	Role       string    `json:"role"`
	CreatedBy  string    `json:"created_by"`
	UpdatedBy  string    `json:"updated_by"`
	DeletedBy  string    `json:"deleted_by"`
	CreatedAt  time.Time `json:"created_at"`
}
