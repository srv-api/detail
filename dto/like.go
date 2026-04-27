package dto

import "time"

type LikeRequest struct {
	UserID       string `json:"user_id"`
	TargetUserID string `json:"target_user_id" validate:"required"`
	IsSuperLike  bool   `json:"is_super_like"`
}

type LikeResponse struct {
	IsMatch bool   `json:"is_match"`
	Message string `json:"message"`
}

type LikeMeResponse struct {
	UserID      string    `json:"user_id"`
	FullName    string    `json:"full_name"`
	PhotoURL    string    `json:"photo_url"`
	IsSuperLike bool      `json:"is_super_like"`
	CreatedAt   time.Time `json:"created_at"`
}
