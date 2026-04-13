package dto

import (
	"time"
)

type DiscountRequest struct {
	ID                 string    `json:"id"`
	UserID             string    `json:"user_id"`
	MerchantID         string    `json:"merchant_id"`
	DiscountName       string    `json:"discount_name"`
	DiscountPercentage uint      `json:"discount_percentage"`
	Status             int       `json:"status"`
	Description        string    `json:"description"`
	CreatedBy          string    `json:"created_by"`
	UpdatedBy          string    `json:"updated_by"`
	DeletedBy          string    `json:"deleted_by"`
	CreatedAt          time.Time `json:"created_at"`
}

type DiscountResponse struct {
	ID                 string    `json:"id"`
	UserID             string    `json:"user_id"`
	MerchantID         string    `json:"merchant_id"`
	DiscountName       string    `json:"discount_name"`
	DiscountPercentage uint      `json:"discount_percentage"`
	Status             string    `json:"status"`
	Description        string    `json:"description"`
	CreatedBy          string    `json:"created_by"`
	UpdatedBy          string    `json:"updated_by"`
	DeletedBy          string    `json:"deleted_by"`
	CreatedAt          time.Time `json:"created_at"`
}
type DiscountUpdateRequest struct {
	ID                 string    `json:"id"`
	UserID             string    `json:"user_id"`
	MerchantID         string    `json:"merchant_id"`
	DiscountName       string    `json:"discount_name"`
	DiscountPercentage uint      `json:"discount_percentage"`
	Status             int       `json:"status"`
	Description        string    `json:"description"`
	CreatedBy          string    `json:"created_by"`
	UpdatedBy          string    `json:"updated_by"`
	DeletedBy          string    `json:"deleted_by"`
	CreatedAt          time.Time `json:"created_at"`
}

type DiscountUpdateResponse struct {
	ID                 string    `json:"id"`
	UserID             string    `json:"user_id"`
	MerchantID         string    `json:"merchant_id"`
	DiscountName       string    `json:"discount_name"`
	DiscountPercentage uint      `json:"discount_percentage"`
	Status             int       `json:"status"`
	Description        string    `json:"description"`
	CreatedBy          string    `json:"created_by"`
	UpdatedBy          string    `json:"updated_by"`
	DeletedBy          string    `json:"deleted_by"`
	CreatedAt          time.Time `json:"created_at"`
}
