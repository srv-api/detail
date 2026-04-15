package entity

import "time"

type UserLimit struct {
	UserID             string    `gorm:"type:varchar(50);primaryKey"`
	RemainingSwipe     int       `gorm:"default:50"`
	RemainingSuperLike int       `gorm:"default:1"`
	UpdatedAt          time.Time `gorm:"autoUpdateTime"`
}
