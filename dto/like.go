package dto

type LikeRequest struct {
	UserID       string `json:"user_id"`
	TargetUserID string `json:"target_user_id" validate:"required"`
	IsSuperLike  bool   `json:"is_super_like"`
}

type LikeResponse struct {
	IsMatch bool   `json:"is_match"`
	Message string `json:"message"`
}
