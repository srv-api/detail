package dto

import (
	"time"
)

type UserPaginationResponse struct {
	Limit        int            `json:"limit"`
	Page         int            `json:"page"`
	Sort         string         `json:"sort"`
	TotalRows    int            `json:"total_rows"`
	TotalPages   int            `json:"total_page"`
	FirstPage    string         `json:"first_page"`
	PreviousPage string         `json:"previous_page"`
	NextPage     string         `json:"next_page"`
	LastPage     string         `json:"last_page"`
	FromRow      int            `json:"from_row"`
	ToRow        int            `json:"to_row"`
	Data         []UserResponse `json:"data"`
	Searchs      []Search       `json:"searchs"`
}

type UserRequest struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	MerchantID  string    `json:"merchant_id"`
	FullName    string    `json:"full_name"`
	User        string    `json:"user"`
	Status      int       `json:"status"`
	Description string    `json:"description"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `json:"updated_by"`
	DeletedBy   string    `json:"deleted_by"`
	CreatedAt   time.Time `json:"created_at"`
}

type UserResponse struct {
	ID            string             `json:"id"`
	UserID        string             `json:"user_id"`
	MerchantID    string             `json:"merchant_id"`
	FullName      string             `json:"full_name"`
	Whatsapp      string             `json:"whatsapp"`
	Email         string             `json:"email"`
	Password      string             `json:"password"`
	AccessRoleID  string             `json:"access_role_id"`
	LoginAttempts int                `json:"login_attempts"`
	Suspended     bool               `json:"suspended"`
	LastAttempt   time.Time          `json:"last_attempt"`
	Status        string             `json:"status"`
	Description   string             `json:"description"`
	Verified      UserVerified       `json:"verified"`
	UserDetail    UserDetailResponse `json:"user_detail"`
	CreatedBy     string             `json:"created_by"`
	UpdatedBy     string             `json:"updated_by"`
	DeletedBy     string             `json:"deleted_by"`
	CreatedAt     time.Time          `json:"created_at"`
}

type UserVerified struct {
	ID             string    `json:"id"`
	UserID         string    `json:"user_id"`
	Token          string    `json:"token"`
	Verified       bool      `json:"verified"`
	StatusAccount  bool      `json:"status_account"`
	AccountExpired time.Time `json:"account_expired"`
	Otp            string    `json:"otp"`
	ExpiredAt      time.Time `json:"expired_at"`
}

type UserUpdateRequest struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	MerchantID  string    `json:"merchant_id"`
	FullName    string    `json:"full_name"`
	Email       string    `json:"email"`
	Status      bool      `json:"status"`
	Description string    `json:"description"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `json:"updated_by"`
	DeletedBy   string    `json:"deleted_by"`
	CreatedAt   time.Time `json:"created_at"`
}

type UserUpdateResponse struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	MerchantID  string    `json:"merchant_id"`
	FullName    string    `json:"full_name"`
	Status      bool      `json:"status"`
	Description string    `json:"description"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `json:"updated_by"`
	DeletedBy   string    `json:"deleted_by"`
	CreatedAt   time.Time `json:"created_at"`
}
