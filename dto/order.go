package dto

type FCMRequest struct {
	Title       string        `json:"title"`
	Body        string        `json:"body"`
	Name        string        `json:"name"`
	ID          uint          `json:"id"`
	UserID      string        `json:"user_id"`
	MerchantID  string        `json:"merchant_id"`
	Product     []ProductItem `json:"product"`
	ProductJSON string        `json:"-" gorm:"column:product"`
	OrderName   string        `json:"order_name"`
	CreatedBy   string        `json:"created_by"`
	Description string        `json:"description"`
	Status      int           `json:"status"`
}

type FCMResponse struct {
	Name      string `json:"name"`
	OrderName string `json:"order_name"`
}

type TokenRequest struct {
	UserID string `json:"user_id"`
	Token  string `json:"token"`
}

type TokenResponse struct {
	Status string `json:"status"`
}

type OrderRequest struct {
	ID          uint          `json:"id"`
	UserID      string        `json:"user_id"`
	MerchantID  string        `json:"merchant_id"`
	Product     []ProductItem `json:"product"`
	ProductJSON string        `json:"-" gorm:"column:product"`
	OrderName   string        `json:"order_name"`
	CreatedBy   string        `json:"created_by"`
	Description string        `json:"description"`
	Status      int           `json:"status"`
}

type ProductItem struct {
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

type OrderResponse struct {
	ID          uint          `json:"id"`
	UserID      string        `json:"user_id"`
	MerchantID  string        `json:"merchant_id"`
	OrderName   string        `json:"order_name"`
	CreatedBy   string        `json:"created_by"`
	Description string        `json:"description"`
	Status      int           `json:"status"`
	Product     []ProductItem `json:"product"`
}

type OrderUpdateRequest struct {
	ID          uint   `json:"id"`
	UserID      string `json:"user_id"`
	MerchantID  string `json:"merchant_id"`
	OrderName   string `json:"order_name"`
	UpdatedBy   string `json:"updated_by"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

type OrderUpdateResponse struct {
	ID          uint   `json:"id"`
	UserID      string `json:"user_id"`
	MerchantID  string `json:"merchant_id"`
	OrderName   string `json:"order_name"`
	UpdatedBy   string `json:"updated_by"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

type GetByIdOrderRequest struct {
	ID uint `param:"id" validate:"required"`
}

type DeleteOrderRequest struct {
	ID        uint   `param:"id" validate:"required"`
	DeletedBy string `json:"deleted_by"`
}

type DeleteOrderResponse struct {
	ID        uint   `param:"id" validate:"required"`
	DeletedBy string `json:"deleted_by"`
}
