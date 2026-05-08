package entity

import "time"

type UserLimit struct {
	UserID             string    `gorm:"type:varchar(50);primaryKey"`
	RemainingSwipe     int       `gorm:"default:50"`
	RemainingBoost     int       `gorm:"default:0"`
	RemainingSuperLike int       `gorm:"default:1"`
	RemainingRewind    int       `gorm:"default:0"`
	UpdatedAt          time.Time `gorm:"autoUpdateTime"`
}
