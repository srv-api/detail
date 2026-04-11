package dto

type CategoryRequest struct {
	ID           string `json:"id"`
	UserID       string `json:"user_id"`
	MerchantID   string `json:"merchant_id"`
	CategoryName string `json:"category_name"`
	CreatedBy    string `json:"created_by"`
	Status       int    `json:"status"`
	Description  string `json:"description"`
}

type CategoryResponse struct {
	ID           string `json:"id"`
	UserID       string `json:"user_id"`
	MerchantID   string `json:"merchant_id"`
	Description  string `json:"description"`
	CategoryName string `json:"category_name"`
	CreatedBy    string `json:"created_by"`
	CreatedAt    string `json:"createdAt_at"`
	Status       int    `json:"status"`
}

type CategoryUpdateRequest struct {
	ID           string `json:"id"`
	UserID       string `json:"user_id"`
	MerchantID   string `json:"merchant_id"`
	CategoryName string `json:"category_name"`
	UpdatedBy    string `json:"updated_by"`
	Description  string `json:"description"`
	Status       int    `json:"status"`
}

type CategoryUpdateResponse struct {
	ID           string `json:"id"`
	UserID       string `json:"user_id"`
	MerchantID   string `json:"merchant_id"`
	CategoryName string `json:"category_name"`
	UpdatedBy    string `json:"updated_by"`
	Description  string `json:"description"`
	Status       int    `json:"status"`
	UpdatedAt    string `json:"updated_at"`
}
