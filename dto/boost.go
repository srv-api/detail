package dto

import "time"

type BoostRequest struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	DetailID  string    `json:"detail_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	CreatedAt time.Time `json:"created_at"`
}

type BoostResponse struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	DetailID  string    `json:"detail_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	CreatedAt time.Time `json:"created_at"`
}
