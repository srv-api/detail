package dto

import "time"

type RoleUserPermissionRequest struct {
	ID           uint   `json:"id"`
	RoleUserID   uint   `json:"role_id"`
	PermissionID uint   `json:"permission_id"`
	UserID       string `json:"user_id"`
	MerchantID   string `json:"merchant_id"`
	CreatedBy    string `json:"created_by"`
}

type RoleUserPermissionResponse struct {
	ID           uint   `json:"id"`
	RoleUserID   uint   `json:"role_id"`
	PermissionID uint   `json:"permission_id"`
	UserID       string `json:"user_id"`
	MerchantID   string `json:"merchant_id"`
	CreatedBy    string `json:"created_by"`
}

type RoleUserPermissionUpdateRequest struct {
	ID           uint   `json:"id"`
	RoleUserID   uint   `json:"role_id"`
	PermissionID uint   `json:"permission_id"`
	UserID       string `json:"user_id"`
	MerchantID   string `json:"merchant_id"`
	CreatedBy    string `json:"created_by"`
	UpdatedBy    string `json:"updated_by"`
}

type RoleUserPermissionUpdateResponse struct {
	ID           uint   `json:"id"`
	RoleUserID   uint   `json:"role_id"`
	PermissionID uint   `json:"permission_id"`
	UserID       string `json:"user_id"`
	MerchantID   string `json:"merchant_id"`
	CreatedBy    string `json:"created_by"`
	UpdatedBy    string `json:"updated_by"`
}

type DeleteRoleUserPermissionRequest struct {
	ID        uint   `param:"id" validate:"required"`
	DeletedBy string `json:"deleted_by"`
}

type DeleteRoleUserPermissionResponse struct {
	ID        uint   `param:"id" validate:"required"`
	DeletedBy string `json:"deleted_by"`
}

type GetRoleUserPermissionResponse struct {
	Items []RoleUserPermissionResponse `json:"items"`
}

type GetRoleUserPermissionByIdRequest struct {
	ID           uint      `json:"id"`
	RoleUserID   uint      `json:"role_id"`
	PermissionID uint      `json:"permission_id"`
	UserID       string    `json:"user_id"`
	MerchantID   string    `json:"merchant_id"`
	CreatedBy    string    `json:"created_by"`
	UpdatedBy    string    `json:"updated_by"`
	DeletedBy    string    `json:"deleted_by"`
	CreatedAt    time.Time `json:"created_at"`
}
