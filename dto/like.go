package dto

type LikeRequest struct {
	TargetUserID string `json:"target_user_id" validate:"required"`
}

type LikeResponse struct {
	IsMatch bool   `json:"is_match"`
	Message string `json:"message"`
}
