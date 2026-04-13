package dto

type SubscribeRequest struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Phone       string `json:"phone"`
	UserID      string `json:"user_id"`
	CreatedBy   string `json:"created_by"`
	PaymentType string `json:"payment_type"`
	OrderID     string `json:"order_id"`
	GrossAmount int64  `json:"gross_amount"`
	Status      string `json:"status"`
	RedirectURL string `json:"redirect_url"`
}

type SubscribeResponse struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	CreatedBy   string `json:"created_by"`
	OrderID     string `json:"order_id"`
	GrossAmount int64  `json:"gross_amount"`
	Status      string `json:"status"`
	Token       string `json:"token"`
	RedirectURL string `json:"redirect_url"`
}
