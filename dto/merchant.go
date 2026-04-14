package dto

import (
	"time"
)

type GetUserDetailByIdRequest struct {
	ID string `query:"id" validate:"required, id"`
}

type UpdateUserDetailRequest struct {
	ID           string  `json:"id"`
	UserID       string  `json:"user_id" validate:"required"`
	Latitude     float64 `json:"latitude" validate:"required"`
	Longitude    float64 `json:"longitude" validate:"required"`
	Radius       int     `json:"radius"`
	MinAge       int     `json:"min_age"`
	MaxAge       int     `json:"max_age"`
	GenderTarget string  `json:"gender_target"`
	UpdatedBy    string  `json:"updated_by"`
}

type UpdateUserDetailResponse struct {
	ID           string  `json:"id"`
	UserID       string  `json:"user_id" validate:"required"`
	Latitude     float64 `json:"latitude" validate:"required"`
	Longitude    float64 `json:"longitude" validate:"required"`
	Radius       int     `json:"radius"`
	MinAge       int     `json:"min_age"`
	MaxAge       int     `json:"max_age"`
	GenderTarget string  `json:"gender_target"`
	UpdatedBy    string  `json:"updated_by"`
}

type UserDetailRequest struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	Latitude     float64   `json:"latitude"`
	Longitude    float64   `json:"longitude"`
	Radius       int       `json:"radius"`
	MinAge       int       `json:"min_age"`
	MaxAge       int       `json:"max_age"`
	GenderTarget string    `json:"gender_target"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UserDetailResponse struct {
	ID             string    `json:"id"`
	UserID         string    `json:"user_id"`
	Latitude       float64   `json:"latitude"`
	Longitude      float64   `json:"longitude"`
	Radius         int       `json:"radius"`
	MinAge         int       `json:"min_age"`
	MaxAge         int       `json:"max_age"`
	GenderTarget   string    `json:"gender_target"`
	ProfilePicture string    `json:"profile_picture"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type ExploreResponse struct {
	Users []ExploreUserResponse `json:"users"`
}

type ExploreUserResponse struct {
	UserID         string  `json:"user_id"`
	FullName       string  `json:"full_name"`
	Gender         string  `json:"gender"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	Distance       float64 `json:"distance"`
	Bio            string  `json:"bio"`
	Radius         int     `json:"radius"`
	MinAge         int     `json:"min_age"`
	MaxAge         int     `json:"max_age"`
	Age            int     `json:"age"`
	GenderTarget   string  `json:"gender_target"`
	ProfilePicture string  `json:"profile_picture"`
}

type UpdateLocationRequest struct {
	Latitude  float64 `json:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" validate:"required"`
}

type UpdatePreferenceRequest struct {
	Radius       int    `json:"radius"`
	MinAge       int    `json:"min_age"`
	MaxAge       int    `json:"max_age"`
	GenderTarget string `json:"gender_target"`
}
