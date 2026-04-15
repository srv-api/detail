package dto

type MatchResponse struct {
	UserID         string `json:"user_id"`
	FullName       string `json:"full_name"`
	ProfilePicture string `json:"profile_picture"`
}
