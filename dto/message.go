package dto

type MessageRequest struct {
	ID         string `json:"id"`
	UserID     string `json:"user_id"`
	MerchantID string `json:"merchant_id"`
	Message    string `json:"message"`
	CreatedBy  string `json:"created_by"`
	Status     int    `json:"status"`
}

type MessageResponse struct {
	ID         string `json:"id"`
	UserID     string `json:"user_id"`
	MerchantID string `json:"merchant_id"`
	Message    string `json:"message"`
	CreatedBy  string `json:"created_by"`
	Status     int    `json:"status"`
}
