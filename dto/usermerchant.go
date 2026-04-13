package dto

import (
	"time"
)

type UserDetailPaginationResponse struct {
	Limit        int                     `json:"limit"`
	Page         int                     `json:"page"`
	Sort         string                  `json:"sort"`
	TotalRows    int                     `json:"total_rows"`
	TotalPages   int                     `json:"total_page"`
	FirstPage    string                  `json:"first_page"`
	PreviousPage string                  `json:"previous_page"`
	NextPage     string                  `json:"next_page"`
	LastPage     string                  `json:"last_page"`
	FromRow      int                     `json:"from_row"`
	ToRow        int                     `json:"to_row"`
	Data         []GetUserDetailResponse `json:"data"`
	Searchs      []Search                `json:"searchs"`
}

type UserFullRequest struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	DetailID      string    `json:"detail_id"`
	FullName      string    `json:"full_name"`
	Whatsapp      string    `json:"whatsapp"`
	Email         string    `json:"email"`
	Password      string    `json:"password"`
	AccessRoleID  string    `json:"access_role_id"`
	LoginAttempts int       `json:"login_attempts"`
	Suspended     bool      `json:"suspended"`
	LastAttempt   time.Time `json:"last_attempt"`
	Description   string    `json:"description"`
	CreatedBy     string    `json:"created_by"`
	UpdatedBy     string    `json:"updated_by"`
	DeletedBy     string    `json:"deleted_by"`
	CreatedAt     time.Time `json:"created_at"`
}

type UserFullResponse struct {
	ID            string             `json:"id"`
	UserID        string             `json:"user_id"`
	DetailID      string             `json:"detail_id"`
	FullName      string             `json:"full_name"`
	Whatsapp      string             `json:"whatsapp"`
	Email         string             `json:"email"`
	Password      string             `json:"password"`
	AccessRoleID  string             `json:"access_role_id"`
	RoleName      string             `json:"role_name"`
	LoginAttempts int                `json:"login_attempts"`
	Suspended     bool               `json:"suspended"`
	LastAttempt   time.Time          `json:"last_attempt"`
	Description   string             `json:"description"`
	CreatedBy     string             `json:"created_by"`
	UpdatedBy     string             `json:"updated_by"`
	DeletedBy     string             `json:"deleted_by"`
	CreatedAt     time.Time          `json:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at"`
	DeletedAt     *time.Time         `json:"deleted_at,omitempty"`
	Verified      UserDetailVerified `json:"verified"`
	UserDetail    UserDetailResponse `json:"user_detail"`
}

type UserDetailByIdResponse struct {
	ID            string                 `json:"id"`
	UserID        string                 `json:"user_id"`
	DetailID      string                 `json:"detail_id"`
	FullName      string                 `json:"full_name"`
	Whatsapp      string                 `json:"whatsapp"`
	Email         string                 `json:"email"`
	Password      string                 `json:"password"`
	AccessRoleID  string                 `json:"access_role_id"`
	RoleName      string                 `json:"role_name"`
	LoginAttempts int                    `json:"login_attempts"`
	Suspended     bool                   `json:"suspended"`
	LastAttempt   time.Time              `json:"last_attempt"`
	Description   string                 `json:"description"`
	CreatedBy     string                 `json:"created_by"`
	UpdatedBy     string                 `json:"updated_by"`
	DeletedBy     string                 `json:"deleted_by"`
	CreatedAt     time.Time              `json:"created_at"`
	UpdatedAt     time.Time              `json:"updated_at"`
	DeletedAt     *time.Time             `json:"deleted_at,omitempty"`
	Verified      UserDetailVerifiedByID `json:"verified"`
	UserDetail    UserDetailResponse     `json:"user_detail"`
}

type UserDetailVerified struct {
	ID             string    `json:"id"`
	UserID         string    `json:"user_id"`
	Token          string    `json:"token"`
	Verified       string    `json:"verified"`
	StatusAccount  string    `json:"status_account"`
	AccountExpired time.Time `json:"account_expired"`
	Otp            string    `json:"otp"`
	ExpiredAt      time.Time `json:"expired_at"`
}

type UserDetailVerifiedByID struct {
	ID             string    `json:"id"`
	UserID         string    `json:"user_id"`
	Token          string    `json:"token"`
	Verified       bool      `json:"verified"`
	StatusAccount  bool      `json:"status_account"`
	AccountExpired time.Time `json:"account_expired"`
	Otp            string    `json:"otp"`
	ExpiredAt      time.Time `json:"expired_at"`
}

type UserDetailUpdateRequest struct {
	ID           string                 `json:"id"`
	UserID       string                 `json:"user_id"`
	DetailID     string                 `json:"detail_id"`
	FullName     string                 `json:"full_name"`
	Email        string                 `json:"email"`
	Whatsapp     string                 `json:"whatsapp"`
	Password     string                 `json:"password"`
	AccessRoleID string                 `json:"access_role_id"`
	RoleName     string                 `json:"role_name"`
	Verified     UserDetailVerifiedByID `json:"verified"`
	CreatedBy    string                 `json:"created_by"`
	UpdatedBy    string                 `json:"updated_by"`
	CreatedAt    time.Time              `json:"created_at"`
}

type UserDetailUpdateResponse struct {
	ID           string                 `json:"id"`
	UserID       string                 `json:"user_id"`
	DetailID     string                 `json:"detail_id"`
	FullName     string                 `json:"full_name"`
	Email        string                 `json:"email"`
	Whatsapp     string                 `json:"whatsapp"`
	Password     string                 `json:"password"`
	AccessRoleID string                 `json:"access_role_id"`
	RoleName     string                 `json:"role_name"`
	Verified     UserDetailVerifiedByID `json:"verified"`
	CreatedBy    string                 `json:"created_by"`
	UpdatedBy    string                 `json:"updated_by"`
	CreatedAt    time.Time              `json:"created_at"`
}

type GetUserDetailResponse struct {
	ID            string             `json:"id"`
	UserID        string             `json:"user_id"`
	DetailID      string             `json:"detail_id"`
	FullName      string             `json:"full_name"`
	Whatsapp      string             `json:"whatsapp"`
	Email         string             `json:"email"`
	Password      string             `json:"password"`
	AccessRoleID  string             `json:"access_role_id"`
	RoleName      string             `json:"role_name"`
	Permission    string             `json:"permission"`
	LoginAttempts int                `json:"login_attempts"`
	Suspended     bool               `json:"suspended"`
	LastAttempt   time.Time          `json:"last_attempt"`
	Description   string             `json:"description"`
	Verified      UserDetailVerified `json:"verified"`
	UserDetail    UserDetailResponse `json:"user_detail"`
	CreatedBy     string             `json:"created_by"`
	UpdatedBy     string             `json:"updated_by"`
	DeletedBy     string             `json:"deleted_by"`
	CreatedAt     time.Time          `json:"created_at"`
}
